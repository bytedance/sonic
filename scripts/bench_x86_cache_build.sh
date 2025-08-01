#!/usr/bin/env bash
#
# Benchmark build-x86.sh VS build-x86_cache.sh
#
# Usage: 
#   ./scripts/bench-build.sh            # default clang
#   ./scripts/bench-build.sh /path/clang-17
#
set -euo pipefail

###############################################################################
# Settings 
###############################################################################
ROOT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
OLD_SCRIPT="${ROOT_DIR}/scripts/build-x86.sh"
NEW_SCRIPT="${ROOT_DIR}/scripts/build-x86_cache.sh"

CLANG="${1:-clang}"          
RUNS=5                       

###############################################################################
# Timing function
###############################################################################
run_once() {
  local script="$1"; shift
  local t_start t_end duration
  t_start=$(date +%s)
  "$script" "$@" > /dev/null 2>&1
  t_end=$(date +%s)
  duration=$(( t_end - t_start ))
  echo "${duration}"
}

###############################################################################
# Run benchmark
###############################################################################
echo "Compiler      : ${CLANG}"
echo "Old script    : ${OLD_SCRIPT}"
echo "Cache script  : ${NEW_SCRIPT}"
echo "Runs per case : ${RUNS}"
echo

min_old=999999
min_new=999999

echo "----- Old script (${OLD_SCRIPT}) -----"
for i in $(seq 1 "${RUNS}"); do
  d=$(run_once "${OLD_SCRIPT}" "${CLANG}")
  echo "Run ${i}: ${d} s"
  (( d < min_old )) && min_old=${d}
done
echo "Shortest (old): ${min_old} s"
echo

echo "----- New script (${NEW_SCRIPT}) -----"
for i in $(seq 1 "${RUNS}"); do
  if [[ ${i} -eq 1 ]]; then
    d=$(run_once "${NEW_SCRIPT}" "${CLANG}" --force)
    echo "Run ${i} (with --force): ${d} s"
  else
    d=$(run_once "${NEW_SCRIPT}" "${CLANG}")
    echo "Run ${i}: ${d} s"
  fi
  (( d < min_new )) && min_new=${d}
done
echo "Shortest (new): ${min_new} s"
echo

###############################################################################
# Calculate percentage
###############################################################################
if [[ ${min_old} -eq 0 ]]; then
  echo "The old script took 0 seconds, so the percentage cannot be calculated." >&2
  exit 1
fi

percent=$( echo "scale=2; (${min_old} - ${min_new}) * 100 / ${min_old}" | bc )
echo "The new script reduced build time by up to ${percent}%"
