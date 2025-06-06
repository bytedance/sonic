#!/usr/bin/env python3

# Copyright 2022 ByteDance Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import csv
import io
import tempfile
import os
import re
# import matplotlib.pyplot as plt
import numpy as np
import subprocess
import argparse

    
def run(cmd):
    print(cmd)
    if os.system(cmd):
        print ("Failed to run cmd: %s"%(cmd))
        exit(1)

def run_s(cmd):
    print (cmd)
    try:
        res = os.popen(cmd)
    except subprocess.CalledProcessError as e:
        if e.returncode:
            print (e.output)
            exit(1)
    return res.read()

def run_r(cmd):
    print (cmd)
    try:
        cmds = cmd.split(' ')
        data = subprocess.check_output(cmds, stderr=subprocess.STDOUT)
    except subprocess.CalledProcessError as e:
        if e.returncode:
            print (e.output)
            exit(1)
    return data.decode("utf-8") 

def compare(args):
    # detect current branch.
    # result = run_r("git branch")
    current_branch = run_s("git status | head -n1 | sed 's/On branch //'")
    # for br in result.split('\n'):
    #     if br.startswith("* "):
    #         current_branch = br.lstrip('* ')
    #         break

    if not current_branch:
        print ("Failed to detect current branch")
        return None
    
    # get the current diff
    (fd, diff) = tempfile.mkstemp()
    run("git diff > %s"%diff)

    # early return if current is main branch.
    print ("Current branch: %s"%(current_branch))
    if current_branch == "main":
        print ("Cannot compare at the main branch.Please build a new branch")
        return None

    # benchmark current branch    
    target = run_bench(args, "target")
    
   
    # trying to switch to the latest main branch
    run("git checkout -- .")
    if current_branch != "main":
        run("git checkout main")
    run("git pull --allow-unrelated-histories origin main")

    # benchmark main branch
    main = run_bench(args, "main")
    run("git checkout -- .")

    # restore branch
    if current_branch != "main":
        run("git checkout %s"%(current_branch))
        
    run("patch -p1 < %s" % (diff))
    
    # diff the result
    bench_diff(main, target, args.threshold)
    return target

def run_bench(args, name):
    (fd, fname) = tempfile.mkstemp(".%s.txt"%name)
    run("%s 2>&1 | tee %s" %(args.cmd, fname))
    return fname

def bench_name(library, use_reflect, pattern, count=1):
    enable = "0"
    if use_reflect and library == "Sonic":
        enable = "1"

    """运行 Go 基准测试并返回结果"""
    env = {
        "SONIC_USE_OPTDEC": enable,
        "SONIC_USE_FASTMAP": enable,
        "SONIC_NO_ASYNC_GC": enable,
    }
    
    # 构建命令
    cmd = ["go", "test", "-benchmem", "-run=none", f"-bench={pattern}", f"-count={count}", "./generic_test"]
    print(f"运行命令: {' '.join(cmd)}")
    
    # 执行命令并捕获输出
    return ' '.join(cmd)

def run_benchmark(library, pattern, count=1):
    cmd = bench_name(library, True, pattern, count)
    return run_s(cmd)


def bench_diff(main, target, threshold=0.05):
    run("go get golang.org/x/perf/cmd/benchstat && go install golang.org/x/perf/cmd/benchstat")
    csv_content = run_s( "benchstat -format=csv %s %s"%(main, target))
    print("benchstat: %s"%csv_content)
    
    # filter out invalid lines
    valid_headers = {',sec/op,CI,sec/op,CI,vs base,P', ',B/s,CI,B/s,CI,vs base,P'}
    valid_blocks = []
    current_block = []
    parsing = False

    # Filter out valid CSV blocks
    for line in csv_content.splitlines():
        if line in valid_headers:
            if current_block:
                valid_blocks.append('\n'.join(current_block))
            current_block = [line]
            parsing = True
        elif parsing:
            if line.strip() == '':
                parsing = False
                if current_block:
                    valid_blocks.append('\n'.join(current_block))
                current_block = []
            else:
                current_block.append(line)

    if current_block:
        valid_blocks.append('\n'.join(current_block))
        
    # Parse each valid CSV block
    significant_decline = []
    for block in valid_blocks:
        csv_reader = csv.DictReader(io.StringIO(block))
        for row in csv_reader:
            benchmark = row['']
            vs_base = row['vs base']
            if benchmark == 'geomean':
                continue
            
            # Skip rows without a valid "vs base" value
            if not vs_base or vs_base == '~':
                continue

            # Convert "vs base" to a float percentage
            try:
                vs_base_percentage = float(vs_base.strip('%')) / 100
            except ValueError:
                continue

            # Check if the decline is significant
            if vs_base_percentage < 0 and -vs_base_percentage > threshold:
                significant_decline.append({
                    'benchmark': benchmark,
                    'vs_base': vs_base_percentage
                })

    if significant_decline:
        print("found significant decline! %s %f"%(significant_decline[0]['benchmark'], significant_decline[0]['vs_base']))
        exit(2)
        
    return

TIME="time(ns/op)"
TP="throughput(MB/s)"
BYTES="bytes(B/op)"
ALLOCS="allocs(allocs/op)"

def parse_benchmark_output(mode, output, library):
    """解析基准测试输出"""
    pattern = re.compile(
        r"Benchmark" + r"(\w+)" + r"/(\w+)_" + re.escape(library) + 
        r".*?\s+\d+\s+(\d+)\s+ns/op\s+([\d.]+)\s+MB/s\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op"
    )
    
    results = dict()
    for line in output.splitlines():
        match = pattern.search(line)
        if match:
            mode2 = match.group(1)
            test_case = match.group(2)
            ns_per_op = int(match.group(3))
            throughput = float(match.group(4))
            bytes_per_op = int(match.group(5))
            allocs_per_op = int(match.group(6))
            
            name = f"{test_case}_{mode2}"
            # Ensure the test_case key exists in the results dictionary
            if test_case not in results:
                results[name] = []

            results[name].append({
                TIME: ns_per_op,
                TP: throughput,
                BYTES: bytes_per_op,
                ALLOCS: allocs_per_op
            })
    print(results)
    return results

def calculate_averages(results):
    """计算每个测试用例的平均值"""
    averages = {}
    for test_case, runs in results.items():
        if not runs:
            continue
            
        avg_throughput = sum(r[TP] for r in runs) / len(runs)
        avg_allocs = sum(r[ALLOCS] for r in runs) / len(runs)
        avg_bytes = sum(r[BYTES] for r in runs) / len(runs)
        avg_time = sum(r[TIME] for r in runs) / len(runs)
        
        averages[test_case] = {
            TP: avg_throughput,
            ALLOCS: avg_allocs,
            BYTES: avg_bytes,
            TIME: avg_time
        }
    
    return averages

def normalize_to_std(data, std_data):
    """相对于标准库进行归一化处理"""
    normalized = {}
    for test_case, lib_data in data.items():
        if test_case not in std_data:
            continue
            
        normalized[test_case] = {
            TP:  lib_data[TP] / std_data[test_case][TP],
            ALLOCS: lib_data[ALLOCS] / std_data[test_case][ALLOCS],
            BYTES:  lib_data[BYTES] / std_data[test_case][BYTES],
            TIME:  lib_data[TIME] / std_data[test_case][TIME],
        }
    return normalized

def plot_comparison(title, metrics, comparison_data, path):
    # """绘制性能对比图"""
    # test_cases = sorted(comparison_data["Std"].keys())
    
    # # 准备性能数据
    # throughput_data = {}
    # allocs_data = {}
    # bytes_data = {}
    # time_data = {}
    
    # for lib in comparison_data:
    #     throughput_data[lib] = [comparison_data[lib][tc][metrics] for tc in test_cases]

    
    # # 创建图表
    # fig, (ax1) = plt.subplots(1, 1, figsize=(14, 16))
    
    # # 性能对比图
    # x = np.arange(len(test_cases))
    # width = 0.2
    # multiplier = 0
    
    # print("plot data ing")
    # for lib, values in throughput_data.items():
    #     offset = width * multiplier
    #     ax1.bar(x + offset, values, width, label=lib)
    #     for i, v in enumerate(values):
    #         ax1.text(i + offset, v + 0.1, f"{v:.1f}x", ha='center', fontsize=8)
    #     multiplier += 1
    
    # ax1.set_ylabel(metrics)
    # ax1.set_title(title)
    # ax1.set_xticks(x + width * (multiplier - 1) / 2)
    # ax1.set_xticklabels(test_cases, rotation=15, ha='right')
    # # ax1.axhline(y=1, color='gray', linestyle='--')
    # ax1.legend()
    
    # # # 内存分配对比图
    # # multiplier = 0
    # # for lib, values in allocs_data.items():
    # #     offset = width * multiplier
    # #     ax2.bar(x + offset, values, width, label=lib)
    # #     for i, v in enumerate(values):
    # #         ax2.text(i + offset, v + 0.01, f"{v:.3f}x", ha='center', fontsize=8)
    # #     multiplier += 1
    
    # # ax2.set_ylabel('内存分配比例 (相对于标准库)')
    # # ax2.set_title('JSON 库内存分配对比')
    # # ax2.set_xticks(x + width * (multiplier - 1) / 2)
    # # ax2.set_xticklabels(test_cases, rotation=15, ha='right')
    # # ax2.axhline(y=1, color='gray', linestyle='--')
    
    # # # 内存使用对比图
    # # multiplier = 0
    # # for lib, values in bytes_data.items():
    # #     offset = width * multiplier
    # #     ax3.bar(x + offset, values, width, label=lib)
    # #     for i, v in enumerate(values):
    # #         ax3.text(i + offset, v + 0.01, f"{v:.3f}x", ha='center', fontsize=8)
    # #     multiplier += 1
    
    # # ax3.set_ylabel('内存使用比例 (相对于标准库)')
    # # ax3.set_title('JSON 库内存使用对比')
    # # ax3.set_xticks(x + width * (multiplier - 1) / 2)
    # # ax3.set_xticklabels(test_cases, rotation=15, ha='right')
    # # ax3.axhline(y=1, color='gray', linestyle='--')
    
    # plt.tight_layout()
    # plt.savefig(path)
    # plt.show()

def export_comparison_to_csv(csv_filename_base, metrics_key, comparison_data, path):
    """将对比数据导出为 CSV 文件"""
    if not comparison_data:
        print("没有对比数据可导出。")
        return

    # 尝试从 "Std" 获取测试用例，如果不存在则从第一个可用的库获取
    first_lib_key = next(iter(comparison_data)) if comparison_data else None
    if not first_lib_key:
        print("对比数据为空，无法确定测试用例。")
        return
        
    test_cases = sorted(comparison_data.get("Std", comparison_data[first_lib_key]).keys())
    
    libraries = sorted(comparison_data.keys()) # 确保库的顺序一致

    # 构建 CSV 文件名
    csv_file_path = path

    try:
        with open(csv_file_path, 'w', newline='') as csvfile:
            csv_writer = csv.writer(csvfile)
            
            # 写入表头
            header = ["Test Case"] + [f"{lib} ({metrics_key})" for lib in libraries]
            csv_writer.writerow(header)
            
            # 写入数据行
            for test_case in test_cases:
                row = [test_case]
                for lib in libraries:
                    # 从 comparison_data 中获取对应的值
                    # comparison_data[lib][test_case][metrics_key]
                    value = comparison_data.get(lib, {}).get(test_case, {}).get(metrics_key, 'N/A')
                    row.append(value)
                csv_writer.writerow(row)
        print(f"CSV 数据已保存到: {csv_file_path}")
    except IOError as e:
        print(f"写入 CSV 文件失败: {csv_file_path}, 错误: {e}")

def main():
    # 收集所有库的数据
    all_data = {}
    comparison_data = {}

    argparser = argparse.ArgumentParser(description='Tools to test the performance. Example: ./bench.py "go test -bench=. ./..."')
    argparser.add_argument('cmd', type=str, help='Golang benchmark command')
    argparser.add_argument('-d', type=str, dest='diff', help='diff bench')
    argparser.add_argument('-t', type=float, dest='threshold', default=0.1, help='diff bench threshold')
    argparser.add_argument('-c', '--compare', dest='compare', action='store_true',
        help='Compare the current branch with the main branch')

    argparser.add_argument('--libraries', nargs='+', default=["Sonic", "Std"],
                        help='要测试的 JSON 库列表 (如 Sonic, Std, Jsoniter, GoJson)')
    argparser.add_argument('--count', type=int, default=1,
                        help='每个基准测试运行的次数')
    argparser.add_argument('--plot-only', action='store_true',
                        help='仅绘制图表，不运行测试 (需提供 --results-dir)')
    argparser.add_argument('--results-dir', default="./bench_results",
                        help='测试结果保存目录')
    argparser.add_argument('--mode', type=str, default="Unmarshal",
                        help='测试场景')
    args = argparser.parse_args()


    for lib in args.libraries:
        pattern = f"\"Benchmark{args.mode}.*/.*_{lib}$\""
        output_file = f"{args.results_dir}/{lib}_results.txt"
        run_s(f"mkdir -p {args.results_dir}")
        
        if not args.plot_only:
            print(f"测试 {lib} 库...")
            output = run_benchmark(lib, pattern, args.count)
            print("output is ", output)
            if not output:
                print(f"{lib} 测试失败，跳过")
                continue
                
            # 保存结果
            with open(output_file, "w") as f:
                f.write(output)
            print(f"{lib} 结果保存到 {output_file}")
        else:
            # 从文件加载结果
            try:
                with open(output_file, "r") as f:
                    output = f.read()
            except FileNotFoundError:
                print(f"错误: 找不到结果文件 {output_file}")
                continue
        
        # 解析结果
        results = parse_benchmark_output(args.mode, output, lib)
        averages = calculate_averages(results)
        all_data[lib] = averages
    # 检查是否有标准库数据
    if "Std" not in all_data:
        print("错误: 未找到标准库测试结果")
        return
    
    # 归一化处理
    std_data = all_data["Std"]
    comparison_data["Std"] = {tc: {TP: 1, ALLOCS: 1, BYTES: 1, TIME: 1} 
                             for tc in std_data}
    
    for lib, data in all_data.items():
        if lib == "Std":
            continue
        normalized = normalize_to_std(data, std_data)
        comparison_data[lib] = normalized
    
    for lib, data in comparison_data.items():
        if lib == "Std":
            continue
        all_data[f"{lib}_Normalized"] = data

    # 绘制对比图
    export_comparison_to_csv(f"{args.mode} time(ns/op)", TP, all_data, f"{args.results_dir}/{args.mode}_throughput.csv")
    export_comparison_to_csv(f"{args.mode} time(ns/op)", TP, comparison_data, f"{args.results_dir}/{args.mode}_throughput_nomarl.csv")
    plot_comparison(f"{args.mode} Performance (Low is Better)", TP, comparison_data, f"{args.results_dir}/{args.mode}_throughput.png")
    # if args.compare:
    #     compare(args)
    # elif args.diff:
    #     target, base = args.diff.split(',')
    #     bench_diff(target, base, args.threshold)
    # else:
    #     run_bench(args, "target")

if __name__ == "__main__":
    main()
