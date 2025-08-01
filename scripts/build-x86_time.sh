#!/bin/bash

set -e

# Define the directories
SRC_DIR="native"
TMP_DIR="output"
OUT_DIR="internal/native"
TOOL_DIR="tools/asm2asm"
TMPL_DIR="internal/native"
EXTRA_CLAGS=$2
CC=clang-13
if [ "$1" != "" ]; then
    CC=$1
fi
# echo $CC

CPU_ARCS=("sse" "avx2")

# used for check the generated assembly
# `vpcmpeqb` is the necessary instruction in native codes when using avx2
CHECK_ARCS=("xmm" "vpcmpeqb")

CLAGS=(
    "-msse -mpclmul -mno-sse4 -mno-avx -mno-avx2" 
    "-mno-sse4 -mavx -mpclmul -mavx2 -DUSE_AVX2=1" 
)

i=0

total_compile_time=0
total_asm2asm_time=0

for arc in "${CPU_ARCS[@]}"; do
    # Create the output directory if it doesn't exist
    out_dir="$OUT_DIR/$arc"
    tmp_dir="$TMP_DIR/$arc"

    # remove old files and create new
    rm -rf $out_dir
    rm -rf $tmp_dir
    mkdir -p $out_dir
    mkdir -p $tmp_dir
        
    # all tmplates
    for tmpl in "$TMPL_DIR"/*.tmpl; do
        tmpl_name=$(basename "$tmpl" .tmpl)
        sed -e 's/{{PACKAGE}}/'${arc}'/g' $tmpl > "$out_dir/${tmpl_name}.go"
    done

    # all .c files in the source directory
    for src_file in "$SRC_DIR"/*.c; do
        base_name=$(basename "$src_file" .c)
        asm_file="$tmp_dir/${base_name}.s"
    
        # Compile the source file into an assembly file
        # -Wall -Werror 
        compile_start=$(date +%s)
        $CC ${CLAGS[$i]} -target x86_64-apple-macos11 -mno-red-zone -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions -fno-rtti -fno-stack-protector -nostdlib  ${EXTRA_CLAGS} -O3 -S -o $asm_file $src_file
        compile_end=$(date +%s)
        compile_time=$((compile_end - compile_start))
        total_compile_time=$((total_compile_time + compile_time))

        # Execute asm2asm tool
        asm_start=$(date +%s)
        python3 $TOOL_DIR/asm2asm.py -r $out_dir/${base_name}.go $asm_file
        asm_end=$(date +%s)
        asm_time=$((asm_end - asm_start))
        total_asm2asm_time=$((total_asm2asm_time + asm_time))
    done

    # should check the output assembly files
    if ! grep -rq ${CHECK_ARCS[$i]} $out_dir; then
        echo "compiled instructions is incorrect, please check again"
        exit 1
    fi

    ((i=i+1))
done

echo "==================================================="
echo "Total C compile time         :        ${total_compile_time} s"
echo "Total asm2asm execution time :        ${total_asm2asm_time} s"
echo "==================================================="

