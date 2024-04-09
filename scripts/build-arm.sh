#!/bin/bash

# Define the directories
SRC_DIR="native"
TMP_DIR="output/arm"
OUT_DIR="internal/native/neon"
TOOL_DIR="tools"
CC=clang
if [ "$1" != "" ]; then
    CC=$1
fi
echo $CC

# Create the output directory if it doesn't exist
mkdir -p "$TMP_DIR"
mkdir -p "$OUT_DIR"

# Loop through all .c files in the source directory
for src_file in "$SRC_DIR"/*.c; do
    # Extract the filename without the extension
    base_name=$(basename "$src_file" .c)
    
    # Define the output file path
    asm_file="$TMP_DIR/${base_name}.s"

    # Compile the source file into an assembly file
    $CC -Wno-error -Wno-nullability-completeness -mllvm=--go-frame -mllvm=--enable-shrink-wrap=0 -target aarch64-apple-macos11 -march=armv8-a+simd -Itools/simde/simde -mno-red-zone -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions -fno-rtti -fno-stack-protector -nostdlib -O3 -mno-red-zone -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions -fno-rtti -fno-stack-protector -nostdlib -S -o "$asm_file" "$src_file" 

    # Execute asm2asm tool
    python3 ${TOOL_DIR}/asm2arm/arm.py ${OUT_DIR}/${base_name}_arm64.go $asm_file

done
