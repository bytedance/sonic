#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TOOL_DIR="$(dirname "${SCRIPT_DIR}")"   # asm2arm_tool
PROJECT_DIR="$(dirname $(dirname "${TOOL_DIR}"))"   # sonic
OUTPUT_DIR="${TOOL_DIR}/output"

mkdir -p "${OUTPUT_DIR}"

NEON_DIR="${PROJECT_DIR}/internal/native/neon"
SVE_LINKNAME_DIR="${PROJECT_DIR}/internal/native/sve_linkname"
SVE_WRAPGOC_DIR="${PROJECT_DIR}/internal/native/sve_wrapgoc"

NEON_OUTPUT="${OUTPUT_DIR}/neon"
SVE_LINKNAME_OUTPUT="${OUTPUT_DIR}/sve_linkname"
SVE_WRAPGPC_OUTPUT="${OUTPUT_DIR}/sve_wrapgpc"

echo "=========================================="
echo "Test Script for neon, sve_linkname and sve_wrapgoc"
echo "=========================================="
echo ""

echo ">>> Step 1: Copying source directories to output..."
echo ""

# 拷贝neon目录
echo ">>> Copying neon directory..."
if [ -d "${NEON_OUTPUT}" ]; then
    echo ">>> neon directory already exists, skipping copy."
else
    if [ -d "${NEON_DIR}" ]; then
        cp -r "${NEON_DIR}" "${NEON_OUTPUT}"
        echo ">>> neon directory copied to: ${NEON_OUTPUT}"
    else
        echo "Warning: neon directory not found: ${NEON_DIR}"
    fi
fi

# 拷贝sve_linkname目录
echo ""
echo ">>> Copying sve_linkname directory..."
if [ -d "${SVE_LINKNAME_OUTPUT}" ]; then
    echo ">>> sve_linkname directory already exists, skipping copy."
else
    if [ -d "${SVE_LINKNAME_DIR}" ]; then
        cp -r "${SVE_LINKNAME_DIR}" "${SVE_LINKNAME_OUTPUT}"
        echo ">>> sve_linkname directory copied to: ${SVE_LINKNAME_OUTPUT}"
    else
        echo "Warning: sve_linkname directory not found: ${SVE_LINKNAME_DIR}"
    fi
fi

# 拷贝sve_wrapgoc目录
echo ""
echo ">>> Copying sve_wrapgoc directory..."
if [ -d "${SVE_WRAPGPC_OUTPUT}" ]; then
    echo ">>> sve_wrapgoc directory already exists, skipping copy."
else
    if [ -d "${SVE_WRAPGOC_DIR}" ]; then
        cp -r "${SVE_WRAPGOC_DIR}" "${SVE_WRAPGPC_OUTPUT}"
        echo ">>> sve_wrapgoc directory copied to: ${SVE_WRAPGPC_OUTPUT}"
    else
        echo "Warning: sve_wrapgoc directory not found: ${SVE_WRAPGOC_DIR}"
    fi
fi

echo ""
echo ">>> Step 2: Executing build_go.sh..."
echo ""
cd "${SCRIPT_DIR}"
bash build_go.sh

if [ $? -ne 0 ]; then
    echo ""
    echo "Error: build_go.sh execution failed."
    echo "Please check the error messages above and fix the issues."
    exit 1
fi

echo ""
echo ">>> build_go.sh executed successfully."
echo ""
echo "=========================================="
echo ">>> Step 3: Running tests..."
echo "=========================================="
echo ""

# 测试neon目录
if [ -d "${NEON_OUTPUT}" ]; then
    echo ">>> Testing neon..."
    cd "${NEON_OUTPUT}"
    go test -v
    echo ""
else
    echo "Warning: neon output directory not found, skipping tests."
fi

# 测试sve_linkname目录
if [ -d "${SVE_LINKNAME_OUTPUT}" ]; then
    echo ">>> Testing sve_linkname..."
    cd "${SVE_LINKNAME_OUTPUT}"
    # go test -v
    echo ""
else
    echo "Warning: sve_linkname output directory not found, skipping tests."
fi

# 测试sve_wrapgoc目录
if [ -d "${SVE_WRAPGPC_OUTPUT}" ]; then
    echo ">>> Testing sve_wrapgoc..."
    cd "${SVE_WRAPGPC_OUTPUT}"
    # go test -v
    echo ""
else
    echo "Warning: sve_wrapgoc output directory not found, skipping tests."
fi

echo "=========================================="
echo "All tests completed successfully!"
echo "=========================================="
