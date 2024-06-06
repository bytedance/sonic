#!/bin/bash

set -xe

# Define the directories
SRC_DIR="native"
TMP_DIR="output"
OUT_DIR="internal/native"
TOOL_DIR="tools/asm2asm"
TMPL_DIR="internal/native"
CC=clang
if [ "$1" != "" ]; then
    CC=$1
fi

CPU_ARCS=("sse" "avx2")

# used for check the generated assembly
# `vpcmpeqb` is the necessary instruction in native codes when using avx2
CHECK_ARCS=("xmm" "vpcmpeqb")

CLAGS=(
    "-msse -mpclmul -mno-sse4 -mno-avx -mno-avx2" 
    "-mno-sse4 -mavx -mpclmul -mavx2 -DUSE_AVX2=1" 
)

i=0

for arc in "${CPU_ARCS[@]}"; do
    # Create the output directory if it doesn't exist
    out_dir="$OUT_DIR/$arc"
    tmp_dir="$TMP_DIR/$arc"

    # remove old files and create new
    rm -vrf $out_dir
    rm -vrf $tmp_dir
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
        $CC ${CLAGS[$i]} -target x86_64-apple-macos11 -mno-red-zone -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions -fno-rtti -fno-stack-protector -nostdlib -O3 -Wall -Werror -S -o $asm_file $src_file

        # Execute asm2asm tool
        python3 $TOOL_DIR/asm2asm.py -r $out_dir/${base_name}.go $asm_file
    done

    # should check the output assembly files
    if ! grep -rq ${CHECK_ARCS[$i]} $out_dir; then
        echo "compiled instructions is incorret, pleas check again"
        exit 1
    fi

    ((i=i+1))

done
