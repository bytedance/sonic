#!/bin/bash

# tool/scripts/build.sh
# 在 tool/build/ 下进行 out-of-source 构建

set -e

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TOOL_DIR="$(dirname "$SCRIPT_DIR")"   # asm2arm_tool
BUILD_DIR="$TOOL_DIR/build"
LLVM_DIR="$TOOL_DIR/llvm-project"
LLVM_BUILD_DIR="$BUILD_DIR/llvm-build"
LLVM_INSTALL_DIR="$BUILD_DIR/llvm-install"

# 检查是否需要拉取 LLVM
echo ">>> Tool directory: $TOOL_DIR"
echo ">>> Build directory: $BUILD_DIR"
echo ">>> LLVM directory: $LLVM_DIR"

# 创建 build 目录
mkdir -p "$BUILD_DIR"

# 拉取 LLVM 代码
if [ ! -d "$LLVM_DIR" ]; then
    echo ">>> Cloning LLVM project..."
    git clone -b feat-goframe-sve --single-branch --depth=1 https://gitcode.com/whoiskk/llvm-project.git "$LLVM_DIR"
else
    echo ">>> LLVM directory already exists, skipping clone."
fi

# 构建 LLVM
if [ ! -d "$LLVM_INSTALL_DIR" ]; then
    echo ">>> Building LLVM..."
    mkdir -p "$LLVM_BUILD_DIR"
    cd "$LLVM_BUILD_DIR"
    
    # 配置 LLVM，只启用 clang 和 lld
    if command -v ninja &> /dev/null; then
        echo ">>> Configuring LLVM with CMake + Ninja..."
        cmake -G Ninja "$LLVM_DIR/llvm" \
            -DCMAKE_BUILD_TYPE=Release \
            -DCMAKE_INSTALL_PREFIX="$LLVM_INSTALL_DIR" \
            -DLLVM_ENABLE_PROJECTS="clang;lld" \
            -DLLVM_ABI_BREAKING_CHECKS=FORCE_OFF
        
        echo ">>> Building LLVM with Ninja..."
        ninja
        echo ">>> Installing LLVM with Ninja..."
        ninja install
    else
        echo ">>> Configuring LLVM with CMake + Makefiles..."
        cmake "$LLVM_DIR/llvm" \
            -DCMAKE_BUILD_TYPE=Release \
            -DCMAKE_INSTALL_PREFIX="$LLVM_INSTALL_DIR" \
            -DLLVM_ENABLE_PROJECTS="clang;lld" \
            -DLLVM_ABI_BREAKING_CHECKS=FORCE_OFF
        
        echo ">>> Building LLVM with Make..."
        make -j$(nproc)
        echo ">>> Installing LLVM with Make..."
        make install
    fi
else
    echo ">>> LLVM already built and installed, skipping."
fi

cd "$BUILD_DIR"

# 编译工具，依赖本地构建的 LLVM
if command -v ninja &> /dev/null; then
    echo ">>> Configuring project with CMake + Ninja..."
    cmake -G Ninja ../src \
        -DLLVM_DIR="$LLVM_INSTALL_DIR/lib/cmake/llvm"
    echo ">>> Building with Ninja..."
    ninja
else
    echo ">>> Configuring project with CMake + Makefiles..."
    cmake ../src \
        -DLLVM_DIR="$LLVM_INSTALL_DIR/lib/cmake/llvm"
    echo ">>> Building with Make..."
    make -j$(nproc)
fi

echo ">>> LLVM installed in: $LLVM_INSTALL_DIR/"
echo ">>> Build succeeded! Binary is in: $BUILD_DIR/"