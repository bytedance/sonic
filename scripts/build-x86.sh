#!/bin/bash

# Define the directories
SRC_DIR="native/arm"
TMP_DIR="output/$2"
OUT_DIR="internal/native/$2"
TOOL_DIR="tools/asm/x86"
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
    $CC -msse -mno-sse4 -mno-avx -mno-avx2 -mpclmul -mno-red-zone -target x86_64-apple-macos11 -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions -fno-rtti -fno-stack-protector -nostdlib -O3 -Wall -Werror -S -o "$asm_file" "$src_file" 

    # Execute asm2asm tool
    python3 ${TOOL_DIR}/asm2asm.py ${OUT_DIR}/${base_name}_amd64.go $asm_file

done
