#!/bin/bash
set -xe

# -----------------------------------------------------------------------------
# Directory configuration (default)
# -----------------------------------------------------------------------------
SRC_DIR="native"
TMP_DIR="output"
OUT_DIR_DEFAULT="internal/native"   # Final output
TOOL_DIR="tools/asm2asm"
TMPL_DIR="internal/native"
CACHE_DIR="${TMP_DIR}/build_cache"

# -----------------------------------------------------------------------------
# Default parameters
# -----------------------------------------------------------------------------
CC="clang-13"
USE_CACHE=1
FORCE=0
TEST_MODE=0

# -----------------------------------------------------------------------------
# Parse command-line arguments
# -----------------------------------------------------------------------------
HAS_CC=0
while [[ $# -gt 0 ]]; do
    case "$1" in
        --no-cache)
            USE_CACHE=0
            shift
            ;;
        --force)
            USE_CACHE=0
            FORCE=1
            shift
            ;;
        --test)
            TEST_MODE=1
            shift
            ;;
        *)
            if [[ ${HAS_CC} -eq 0 ]]; then
                CC="$1"
                HAS_CC=1
                shift
            else
                echo "Error: Only one compiler argument is allowed; the rest must be --no-cache / --force / --test only. " >&2
                exit 1
            fi
            ;;
    esac
done

# Determine the final output directory based on --test.
if [[ ${TEST_MODE} -eq 1 ]]; then
    OUT_DIR="${TMP_DIR}/test"
else
    OUT_DIR="${OUT_DIR_DEFAULT}"
fi

#If --force is specified, clear the cache.
if [[ ${FORCE} -eq 1 ]]; then
    rm -rf "${CACHE_DIR}"
fi
mkdir -p "${CACHE_DIR}"

echo "CC=${CC}  USE_CACHE=${USE_CACHE}  FORCE=${FORCE}  TEST_MODE=${TEST_MODE}"
echo "OUT_DIR=${OUT_DIR}"

# -----------------------------------------------------------------------------
# Helper function
# -----------------------------------------------------------------------------
hash_data() { /usr/bin/shasum -a 1 | awk '{print $1}'; }

# -----------------------------------------------------------------------------
# Compilation process 
# -----------------------------------------------------------------------------
CPU_ARCS=("sse" "avx2")
CHECK_ARCS=("xmm" "vpcmpeqb")
CLAGS=(
    "-msse -mpclmul -mno-sse4 -mno-avx -mno-avx2"
    "-mno-sse4 -mavx -mpclmul -mavx2 -DUSE_AVX2=1"
)

i=0
for arc in "${CPU_ARCS[@]}"; do
    out_dir="${OUT_DIR}/${arc}"
    tmp_dir="${TMP_DIR}/${arc}"

    rm -rf "${out_dir}" "${tmp_dir}"
    mkdir -p "${out_dir}" "${tmp_dir}"

    # -------------------------------------------------------------------------
    # 1. tmpl → .go
    # -------------------------------------------------------------------------
    for tmpl in "${TMPL_DIR}"/*.tmpl; do
        tmpl_name=$(basename "${tmpl}" .tmpl)
        go_file="${out_dir}/${tmpl_name}.go"
        tmpl_hash=$( { cat "${tmpl}"; echo "${arc}"; } | hash_data )
        cache_go="${CACHE_DIR}/${tmpl_hash}.go"

        if [[ ${USE_CACHE} -eq 1 && -f "${cache_go}" ]]; then
            cp "${cache_go}" "${go_file}"
        else
            sed -e "s/{{PACKAGE}}/${arc}/g" "${tmpl}" > "${go_file}"
            cp "${go_file}" "${cache_go}"
        fi
    done

    # -------------------------------------------------------------------------
    # 2. .c → .s   3. asm2asm
    # -------------------------------------------------------------------------
    for src_file in "${SRC_DIR}"/*.c; do
        base_name=$(basename "${src_file}" .c)
        asm_file="${tmp_dir}/${base_name}.s"

        # ------------ 2) Compile C → .s ------------
        compile_hash=$( { cat "${src_file}"; echo "${CC} ${CLAGS[$i]}"; } | hash_data )
        cache_s="${CACHE_DIR}/${compile_hash}.s"

        if [[ ${USE_CACHE} -eq 1 && -f "${cache_s}" ]]; then
            cp "${cache_s}" "${asm_file}"
        else
            $CC ${CLAGS[$i]} -target x86_64-apple-macos11 -mno-red-zone \
                -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions \
                -fno-rtti -fno-stack-protector -nostdlib \
                -O3 -S -o "${asm_file}" "${src_file}"
            cp "${asm_file}" "${cache_s}"
        fi

        # ------------ 3) asm2asm ------------
        subr_file="${out_dir}/${base_name}_subr.go"
        text_file="${out_dir}/${base_name}_text_amd64.go"
        asm2asm_hash=$( { cat "${asm_file}"; cat "${out_dir}/${base_name}.go"; } | hash_data )
        cache_subr="${CACHE_DIR}/${asm2asm_hash}_subr.go"
        cache_text="${CACHE_DIR}/${asm2asm_hash}_text.go"

        if [[ ${USE_CACHE} -eq 1 && -f "${cache_subr}" && -f "${cache_text}" ]]; then
            cp "${cache_subr}"  "${subr_file}"
            cp "${cache_text}"  "${text_file}"
        else
            python3 "${TOOL_DIR}/asm2asm.py" -r "${out_dir}/${base_name}.go" "${asm_file}"
            cp "${subr_file}" "${cache_subr}"
            cp "${text_file}" "${cache_text}"
        fi
    done

    # -------------------------------------------------------------------------
    # Output command verification
    # -------------------------------------------------------------------------
    if ! grep -rq "${CHECK_ARCS[$i]}" "${out_dir}"; then
        echo "compiled instructions incorrect, please check again"
        exit 1
    fi

    ((i=i+1))
done
