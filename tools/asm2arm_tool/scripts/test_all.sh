#!/bin/bash

# 项目根目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TOOL_DIR="$(dirname "${SCRIPT_DIR}")"   # asm2arm_tool
PROJECT_DIR="$(dirname $(dirname "${TOOL_DIR}"))"   # sonic
OUTPUT_DIR="${TOOL_DIR}/output"
NATIVE_DIR="$PROJECT_DIR/internal/native"

# 拷贝文件
function copy_files() {
    echo "Copying files from $OUTPUT_DIR to $NATIVE_DIR..."
    
    # 拷贝 neon 文件
    if [ -d "$OUTPUT_DIR/neon" ]; then
        cp -r "$OUTPUT_DIR/neon"/* "$NATIVE_DIR/neon/" 2>/dev/null || true
        echo "Copied neon files"
    fi
    
    # 拷贝 sve_wrapgoc 文件
    if [ -d "$OUTPUT_DIR/sve_wrapgoc" ]; then
        cp -r "$OUTPUT_DIR/sve_wrapgoc"/* "$NATIVE_DIR/sve_wrapgoc/" 2>/dev/null || true
        echo "Copied sve_wrapgoc files"
    fi
    
    # 拷贝 sve_linkname 文件
    if [ -d "$OUTPUT_DIR/sve_linkname" ]; then
        cp -r "$OUTPUT_DIR/sve_linkname"/* "$NATIVE_DIR/sve_linkname/" 2>/dev/null || true
        echo "Copied sve_linkname files"
    fi
}

# 执行测试
function run_tests() {
    local wrapgoc=$1
    local linkname=$2
    local test_dir="$OUTPUT_DIR/test_results/wrapgoc_${wrapgoc}_linkname_${linkname}"
    
    echo "\nRunning tests with SONIC_USE_SVE_WRAPGOC=$wrapgoc, SONIC_USE_SVE_LINKNAME=$linkname..."
    mkdir -p "$test_dir"
    
    # 设置环境变量
    export SONIC_USE_SVE_WRAPGOC=$wrapgoc
    export SONIC_USE_SVE_LINKNAME=$linkname
    
    # 执行编码器测试
    echo "Running encoder tests..."
    go test -v $PROJECT_DIR/internal/encoder/... > "$test_dir/encoder_test.log"
    echo "Encoder tests completed. Log saved to $test_dir/encoder_test.log"
    
    # 执行 UTF-8 测试
    echo "Running UTF-8 tests..."
    go test -v $PROJECT_DIR/utf8/... > "$test_dir/utf8_test.log"
    echo "UTF-8 tests completed. Log saved to $test_dir/utf8_test.log"
}

# 主函数
function main() {
    # 拷贝文件
    copy_files
    
    # 创建测试结果目录
    mkdir -p "$OUTPUT_DIR/test_results"
    
    # 执行所有组合的测试
    run_tests 0 0
    run_tests 1 0
    run_tests 0 1
    run_tests 1 1
    
    echo "\nAll tests completed!"
    echo "Test results saved in $OUTPUT_DIR/test_results/"
}

# 执行主函数
main
