#!/usr/bin/env bash
#
# scripts/build-x86_output.sh
# -----------------------------------------------------------------------------
# 1. *.tmpl  → internal/native/<arch>/<name>.go
# 2. *.c     → clang → .s → asm2asm.py → internal/native/<arch>/<name>.go 
# 3. output：
#    - files output from tmpl:
#    - files output from asm2asm execution:
#    - suffixes of asm2asm output files：<suffix list>
# -----------------------------------------------------------------------------

set -Eeuo pipefail

# ---------- Directories -------------------------------------------------------------
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
cd "$ROOT_DIR"

SRC_DIR="native"
TMP_DIR="output"
OUT_DIR="internal/native"
TMPL_DIR="internal/native"
TOOL_DIR="tools/asm2asm"

# ---------- flags --------------------------------------------------------
CC="${1:-clang-13}"

CPU_ARCHES=(sse avx2)
CFLAGS=(
  "-msse -mpclmul -mno-sse4 -mno-avx -mno-avx2"
  "-mno-sse4 -mavx -mpclmul -mavx2 -DUSE_AVX2=1"
)

# ---------- read_into_array ----------------------------
read_into_array() {            # read_into_array <arr> "command …"
  local _arr="$1"; shift
  local _tmp=() _line
  while IFS= read -r _line; do
    [[ -n "$_line" ]] && _tmp+=("$_line")
  done < <(eval "$*")
  eval "$_arr=(\"\${_tmp[@]}\")"
}

# -----------------------------------------------------------------------------
for ((idx=0; idx<${#CPU_ARCHES[@]}; idx++)); do
  arch="${CPU_ARCHES[$idx]}"
  out_dir="$OUT_DIR/$arch"
  tmp_dir="$TMP_DIR/$arch"

  echo "==> [$arch] clear directory: $out_dir"
  rm -rf "$out_dir" "$tmp_dir"
  mkdir -p "$out_dir" "$tmp_dir"

  # ---------- 1. tmpl -------------------------------------------------------
  tmpl_files=()
  shopt -s nullglob
  for tmpl in "$TMPL_DIR"/*.tmpl; do
    base=$(basename "$tmpl" .tmpl)
    sed "s/{{PACKAGE}}/${arch}/g" "$tmpl" > "$out_dir/${base}.go"
    tmpl_files+=("$out_dir/${base}.go")
  done
  shopt -u nullglob

  # ---------- 2. C → S → asm2asm ------------------------------------------
  base_names=()                
  shopt -s nullglob
  for cfile in "$SRC_DIR"/*.c; do
    base=$(basename "$cfile" .c)
    base_names+=("$base")
    asm="$tmp_dir/${base}.s"
    go_out="$out_dir/${base}.go"

    $CC ${CFLAGS[$idx]} -target x86_64-apple-macos11 \
      -mno-red-zone -fno-asynchronous-unwind-tables \
      -fno-builtin -fno-exceptions -fno-rtti -fno-stack-protector \
      -nostdlib -O3 -S -o "$asm" "$cfile"

    python3 "$TOOL_DIR/asm2asm.py" -r "$go_out" "$asm"
  done
  shopt -u nullglob

  # ---------- 3. tmpl output files --------------------------------------------------
  echo
  echo "files output from tmpl:"
  if ((${#tmpl_files[@]})); then
    printf '  %s\n' "${tmpl_files[@]}"
  else
    echo "  (none)"
  fi

  # ---------- 4. asm2asm output files ----------------------------------------------
  read_into_array all_go "find \"$out_dir\" -type f -name '*.go' | sort"
  asm2asm_files=()
  for f in "${all_go[@]}"; do
    skip=false
    for t in "${tmpl_files[@]}"; do
      [[ $f == "$t" ]] && { skip=true; break; }
    done
    $skip || asm2asm_files+=("$f")
  done

  echo
  echo "files output from asm2asm execution:"
  if ((${#asm2asm_files[@]})); then
    printf '  %s\n' "${asm2asm_files[@]}"
  else
    echo "  (none)"
  fi

  # ---------- 5. get suffix -------------------------------
  suffix_set=()
  for file in "${asm2asm_files[@]}"; do
    bn=$(basename "$file" .go)                
    longest_base=""                           
    for base in "${base_names[@]}"; do
      [[ "$bn" == "$base"* ]] || continue
      if (( ${#base} > ${#longest_base} )); then
        longest_base="$base"
      fi
    done

    
    [[ "$bn" == "$longest_base" ]] && continue

    suffix_part="${bn#$longest_base}"         
    if [[ "$suffix_part" == _* && -n "$suffix_part" ]]; then
      suffix="${suffix_part}.go"
      [[ " ${suffix_set[*]} " == *" $suffix "* ]] || suffix_set+=("$suffix")
    fi
  done

  echo
  if ((${#suffix_set[@]})); then
    echo "asm2asm suffix list: ${suffix_set[*]}"
  fi
  echo
done
