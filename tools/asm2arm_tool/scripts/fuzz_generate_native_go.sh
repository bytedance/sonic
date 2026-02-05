#!/bin/bash
set -e

# 设置脚本目录和工具路径
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TOOL_DIR="$(dirname "${SCRIPT_DIR}")"   # asm2arm_tool
PROJECT_DIR="$(dirname $(dirname "${TOOL_DIR}"))"   # sonic
BUILD_DIR="${TOOL_DIR}/build"
TOOL_PATH="${BUILD_DIR}/asm2arm_tool"
LLVM_INSTALL_DIR="${BUILD_DIR}/llvm-install"
CLANG_PATH="${LLVM_INSTALL_DIR}/bin/clang"

# 输出目录
FUZZ_OUTPUT_DIR="${TOOL_DIR}/output/fuzz_generate_native_go"
TEST_RESULTS_DIR="${FUZZ_OUTPUT_DIR}/results"

# 确保目录存在
mkdir -p "${FUZZ_OUTPUT_DIR}"
mkdir -p "${TEST_RESULTS_DIR}"

# 测试选项组合

# 优化级别
optimization_levels=(
    "-O0"
    "-O1"
    "-O2"
    "-O3"
)

# 备份原始 generate_native_go.sh 文件
backup_file="${SCRIPT_DIR}/generate_native_go.sh.bak"
cp "${SCRIPT_DIR}/generate_native_go.sh" "${backup_file}"

# 修改 generate_native_go.sh 文件的函数
modify_generate_script() {
    local opt_level=$1
    local test_case=$2
    
    echo "Modifying generate_native_go.sh for test case: ${test_case}"
    
    # 读取原始文件内容
    local original_content=$(cat "${backup_file}")
    
    # 修改优化级别
    local modified_content=${original_content}
    
    # 修改 neon 编译选项
    modified_content=$(echo "${modified_content}" | sed "s/-O3/${opt_level}/g")
    
    # 写入修改后的内容
    echo "${modified_content}" > "${SCRIPT_DIR}/generate_native_go.sh"
    
    echo "Modified generate_native_go.sh with:"
    echo "  Optimization: ${opt_level}"
}

# 去除架构特性的函数
remove_arch_features() {
    local opt_level=$1
    local test_case=$2
    
    echo "Modifying generate_native_go.sh for test case: ${test_case} (no arch features)"
    
    # 读取原始文件内容
    local original_content=$(cat "${backup_file}")
    
    # 修改优化级别
    local modified_content=${original_content}
    
    # 修改优化级别
    modified_content=$(echo "${modified_content}" | sed "s/-O3/${opt_level}/g")
    
    # 去除 neon 的 simd 架构特性
    modified_content=$(echo "${modified_content}" | sed "s/-march=armv8-a+simd/-march=armv8-a/g")
    
    # 去除 sve 的架构特性 (sve_linkname 和 sve_wrapgoc)
    modified_content=$(echo "${modified_content}" | sed "s/-march=armv8-a+sve+aes/-march=armv8-a/g")
    
    # 去除 __SVE__ 宏定义
    modified_content=$(echo "${modified_content}" | sed "s/-D__SVE__//g")
    
    # 去除 features 参数
    modified_content=$(echo "${modified_content}" | sed "s/--features=+sve,+aes//g")
    
    # 写入修改后的内容
    echo "${modified_content}" > "${SCRIPT_DIR}/generate_native_go.sh"
    
    echo "Modified generate_native_go.sh with:"
    echo "  Optimization: ${opt_level}"
    echo "  Arch features: removed (no simd/sve)"
}

# 运行测试脚本的函数
run_native_recover_test() {
    local test_case=$1
    local log_file="${TEST_RESULTS_DIR}/${test_case}_native_recover.log"
    
    echo "Running test_native_recover.sh for test case: ${test_case}"
    
    # 执行测试脚本
    cd "${SCRIPT_DIR}"
    bash test_native_recover.sh 2>"${log_file}"
    
    local exit_code=$?
    
    if [ ${exit_code} -eq 0 ]; then
        echo "test_native_recover.sh execution succeeded for test case: ${test_case}"
    else
        echo "test_native_recover.sh execution failed for test case: ${test_case}"
        echo "Check log file: ${log_file}"
    fi
    
    return ${exit_code}
}

# 运行 encoder api 测试脚本的函数
run_encoder_api_test() {
    local test_case=$1
    local log_file="${TEST_RESULTS_DIR}/${test_case}_encoder_api.log"
    
    echo "Running test_encoder_api.sh for test case: ${test_case}"
    
    # 执行测试脚本
    cd "${SCRIPT_DIR}"
    bash test_encoder_api.sh 2>"${log_file}"
    
    local exit_code=$?
    
    if [ ${exit_code} -eq 0 ]; then
        echo "test_encoder_api.sh execution succeeded for test case: ${test_case}"
    else
        echo "test_encoder_api.sh execution failed for test case: ${test_case}"
        echo "Check log file: ${log_file}"
    fi
    
    return ${exit_code}
}

# 主测试函数
run_fuzz_tests() {
    local test_counter=0
    
    echo "=========================================="
    echo "Running fuzz tests for generate_native_go.sh with different optimization levels"
    echo "=========================================="
    echo ""
    
    # 遍历所有测试组合
    for opt_level in "${optimization_levels[@]}"; do
        test_counter=$((test_counter + 1))
        test_case="test_${test_counter}_opt_${opt_level//-/_}"
        local status_file="${TEST_RESULTS_DIR}/${test_case}.status"
        
        echo ""
        echo "Test case ${test_counter}: ${test_case}"
        echo "--------------------------------------------------"
        
        # 修改 generate_native_go.sh 文件
        modify_generate_script "${opt_level}" "${test_case}"
        
        # 运行 native recover 测试
        run_native_recover_test "${test_case}"
        local native_exit_code=$?
        
        # 运行 encoder api 测试
        run_encoder_api_test "${test_case}"
        local encoder_exit_code=$?
        
        # 记录测试状态
        if [ ${native_exit_code} -eq 0 ] && [ ${encoder_exit_code} -eq 0 ]; then
            echo "0" > "${status_file}"
        else
            echo "1" > "${status_file}"
        fi
        
        echo "--------------------------------------------------"
    done
    
    # 遍历所有测试组合（去除架构特性）
    echo ""
    echo "=========================================="
    echo "Running fuzz tests WITHOUT arch features"
    echo "=========================================="
    echo ""
    
    for opt_level in "${optimization_levels[@]}"; do
        test_counter=$((test_counter + 1))
        test_case="test_${test_counter}_opt_${opt_level//-/_}_noarch"
        local status_file="${TEST_RESULTS_DIR}/${test_case}.status"
        
        echo ""
        echo "Test case ${test_counter}: ${test_case}"
        echo "--------------------------------------------------"
        
        # 修改 generate_native_go.sh 文件（去除架构特性）
        remove_arch_features "${opt_level}" "${test_case}"
        
        # 运行 native recover 测试
        run_native_recover_test "${test_case}"
        local native_exit_code=$?
        
        # 运行 encoder api 测试
        run_encoder_api_test "${test_case}"
        local encoder_exit_code=$?
        
        # 记录测试状态
        if [ ${native_exit_code} -eq 0 ] && [ ${encoder_exit_code} -eq 0 ]; then
            echo "0" > "${status_file}"
        else
            echo "1" > "${status_file}"
        fi
        
        echo "--------------------------------------------------"
    done
    
    # 恢复原始 generate_native_go.sh 文件
    echo "Restoring original generate_native_go.sh file..."
    cp "${backup_file}" "${SCRIPT_DIR}/generate_native_go.sh"
    rm "${backup_file}"
    
    echo ""
    echo "=========================================="
    echo "Fuzz tests completed!"
    echo "=========================================="
    echo "Test results are in: ${TEST_RESULTS_DIR}"
}

# 运行测试
run_fuzz_tests

# 汇总测试结果
echo ""
echo "=========================================="
echo "Test Summary"
echo "=========================================="

# 统计成功和失败的测试
success_count=0
failure_count=0

for status_file in "${TEST_RESULTS_DIR}"/*.status; do
    if [ -f "${status_file}" ]; then
        exit_code=$(cat "${status_file}")
        if [ "${exit_code}" -eq 0 ]; then
            success_count=$((success_count + 1))
        else
            failure_count=$((failure_count + 1))
        fi
    fi
done

total_count=$((success_count + failure_count))

echo "Total tests: ${total_count}"
echo "Success: ${success_count}"
echo "Failure: ${failure_count}"

echo ""
if [ "${failure_count}" -eq 0 ]; then
    echo "All tests passed!"
else
    echo "Some tests failed. Check the log files in ${TEST_RESULTS_DIR} for details."
fi

echo ""
echo "=========================================="
echo "Fuzz test script completed!"
echo "=========================================="
