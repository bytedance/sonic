#!/bin/bash

# tool/scripts/build.sh
# 在 tool/build/ 下进行 out-of-source 构建

set -e

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TOOL_DIR="$(dirname "$SCRIPT_DIR")"   # asm2arm_tool
BUILD_DIR="$TOOL_DIR/build"

echo ">>> Tool directory: $TOOL_DIR"
echo ">>> Build directory: $BUILD_DIR"

# 创建 build 目录
mkdir -p "$BUILD_DIR"

# 进入 build 目录
cd "$BUILD_DIR"

# 编译：优先使用 ninja，否则用 make
if command -v ninja &> /dev/null; then
    echo ">>> Configuring project with CMake + Ninja..."
    cmake -G Ninja ..
    echo ">>> Building with Ninja..."
    ninja
else
    echo ">>> Configuring project with CMake + Makefiles..."
    cmake ..
    echo ">>> Building with Make..."
    make -j$(nproc)
fi

echo ">>> Build succeeded! Binary is in: $BUILD_DIR/"