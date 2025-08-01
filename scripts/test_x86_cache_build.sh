#!/usr/bin/env bash
#
# Verify in four stages whether the caching behavior of build-x86_cache.sh 
# is consistent with the old script (silent mode).
#
# Usage: 
#   ./scripts/test-build-cache.sh             # default clang
#   ./scripts/test-build-cache.sh /path/clang # other clang
set -euo pipefail

###############################################################################
# Paths & constants
###############################################################################
OLD_SCRIPT="scripts/build-x86.sh"
CACHE_SCRIPT="scripts/build-x86_cache.sh"
ROOT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

OUTPUT_DIR="${ROOT_DIR}/output"
TEST_OUT_DIR="${OUTPUT_DIR}/test"
BASE_OUT_DIR="${ROOT_DIR}/internal/native"

CLANG_PATH="${1:-clang}"      # default clang
ARCHES=("sse" "avx2")

###############################################################################
# Utility functions
###############################################################################
compare_dirs() {
  local lhs="$1" rhs="$2" label="$3"
  if diff -r -q "$lhs" "$rhs" >/dev/null; then
    echo "✅  ${label}: match"
  else
    echo "❌  ${label}: mismatch, showing first differences:" >&2
    diff -r "$lhs" "$rhs" | head -n 20 >&2
    exit 1
  fi
}

compare_stage() {
  local stage="$1"
  for arch in "${ARCHES[@]}"; do
    compare_dirs "${BASE_OUT_DIR}/${arch}" "${TEST_OUT_DIR}/${arch}" "Stage ${stage} – ${arch}"
  done
}

# Silently run the specified script
run_quiet() {
  local script="$1"; shift
  "$script" "$@" > /dev/null 2>&1
}

###############################################################################
# 0) Build baseline (old script)
###############################################################################
echo "==== [0] build baseline with old script ===="
rm -rf "${OUTPUT_DIR}"
run_quiet "${OLD_SCRIPT}" "${CLANG_PATH}"

###############################################################################
# 1) Cache first build (clear output/, 0% hit rate)
###############################################################################
echo -e "\n==== [1] cache first build (clean output/) ===="
rm -rf "${OUTPUT_DIR}"
run_quiet "${CACHE_SCRIPT}" "${CLANG_PATH}" --test
compare_stage 1

###############################################################################
# 2) Cache rebuild (should achieve 100% hit rate)
###############################################################################
echo -e "\n==== [2] cache second build (expect full hit) ===="
run_quiet "${CACHE_SCRIPT}" "${CLANG_PATH}" --test
compare_stage 2

###############################################################################
# 3) Cache build with --no-cache (skip cache)
###############################################################################
echo -e "\n==== [3] cache build with --no-cache ===="
run_quiet "${CACHE_SCRIPT}" "${CLANG_PATH}" --test --no-cache
compare_stage 3

###############################################################################
# 4) Cache build with --force (clear and rebuild cache)
###############################################################################
echo -e "\n==== [4] cache build with --force ===="
run_quiet "${CACHE_SCRIPT}" "${CLANG_PATH}" --test --force
compare_stage 4

echo -e "\n🎉  All stages passed – cache script outputs match the old script."
