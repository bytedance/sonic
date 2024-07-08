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

import tempfile
import os
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
    # detech current branch.
    # result = run_r("git branch")
    current_branch = run_s("git status | head -n1 | sed 's/On branch //'")
    # for br in result.split('\n'):
    #     if br.startswith("* "):
    #         current_branch = br.lstrip('* ')
    #         break

    if not current_branch:
        print ("Failed to detech current branch")
        return None
    
    # get the current diff
    (fd, diff) = tempfile.mkstemp()
    run("git diff > %s"%diff)

    # early return if currrent is main branch.
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

    # diff the result
    # benchstat = "go get golang.org/x/perf/cmd/benchstat && go install golang.org/x/perf/cmd/benchstat"
    run( "benchstat %s %s"%(main, target))
    run("git checkout -- .")

    # restore branch
    if current_branch != "main":
        run("git checkout %s"%(current_branch))
        
    run("patch -p1 < %s" % (diff))
    return target

def run_bench(args, name):
    (fd, fname) = tempfile.mkstemp(".%s.txt"%name)
    run("%s 2>&1 | tee %s" %(args.cmd, fname))
    return fname


def main():
    argparser = argparse.ArgumentParser(description='Tools to test the performance. Example: ./bench.py "go test -bench=. ./..."')
    argparser.add_argument('cmd', type=str, help='Golang benchmark command')
    argparser.add_argument('-c', '--compare', dest='compare', action='store_true',
        help='Compare the current branch with the main branch')
    args = argparser.parse_args()

    if args.compare:
        compare(args)
    else:
        run_bench(args, "target")

if __name__ == "__main__":
    main()
