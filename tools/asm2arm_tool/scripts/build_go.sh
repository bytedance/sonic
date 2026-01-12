#!/bin/bash

set -e

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TOOL_DIR="$(dirname "${SCRIPT_DIR}")"   # asm2arm_tool
PROJECT_DIR="$(dirname $(dirname "${TOOL_DIR}"))"   # goserviceopt

BUILD_DIR="${TOOL_DIR}/build"
TOOL_PATH="${BUILD_DIR}/asm2arm_tool"

SRC_DIR="${PROJECT_DIR}/native"
TMPL_DIR="${PROJECT_DIR}/internal/native"
OUTPUT_DIR="${TOOL_DIR}/output"
CC=clang

if [ "$1" != "" ]; then
    CC=$1
fi
mkdir -p "${OUTPUT_DIR}"

echo ">>> Using $CC compiler"

for src_file in "${SRC_DIR}"/*.c; do
    base_name="$(basename "${src_file}" .c)"
    asm_file="${OUTPUT_DIR}/${base_name}.s"
    cerr_log="${OUTPUT_DIR}/${base_name}.log"

    echo ">>> Compile ${src_file}... --> ${asm_file}"
    ${CC} -ffixed-x28 -Wno-error -Wno-nullability-completeness -Wno-incompatible-pointer-types \
    -mno-red-zone -fno-rtti -fno-stack-protector -nostdlib -O3 -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions \
    -march=armv8-a+sve+aes -I/usr/include/simde -D__SVE__ -S -o "${asm_file}" "${src_file}"

    echo ">>> Execute Raw JIT mode..."
    ${TOOL_PATH} --source=${asm_file} --output=${OUTPUT_DIR} --link-ld=${SCRIPT_DIR}/link.ld --TmplDir=${TMPL_DIR} \
    --package=sve_wrapgoc --features=+sve,+aes 2>${cerr_log}
done