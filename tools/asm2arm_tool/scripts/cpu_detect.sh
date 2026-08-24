#!/bin/bash

# These scripts have two very different host requirements, and they used to
# share a single gate that hard-checked for a Kunpeng CPU:
#
#   - Generating the natives (build_tool.sh, generate_native_go.sh,
#     fuzz_generate_native_go.sh) is a pure cross-compilation. The clang/lld
#     that build_tool.sh builds is a cross compiler targeting aarch64, and
#     asm2arm_tool only parses text assembly and emits Go source. Neither
#     executes a single arm64 -- let alone SVE -- instruction, so any host will
#     do. Use check_build_host.
#
#   - Running the generated natives (test_native_recover.sh,
#     test_encoder_api.sh) really does execute them, so it needs arm64
#     hardware, and the sve_* suites additionally need SVE. Use check_test_host.
#
# Keeping one gate for both is what made the difference invisible. It is also
# actively unsafe for the test scripts: `go test` on an x86_64 host happily
# runs the amd64 avx2/sse paths and reports a pass without having exercised any
# of this code.

# 检测是否是鲲鹏CPU
is_kunpeng_cpu() {
    if [ -f "/proc/cpuinfo" ]; then
        if grep -q "CPU part" "/proc/cpuinfo"; then
            local cpu_part=$(grep "CPU part" "/proc/cpuinfo" | head -1 | awk -F: '{print $2}' | tr -d ' ')
            if [ "$cpu_part" = "0xd02" ] || [ "$cpu_part" = "0xd06" ]; then
                return 0
            fi
        fi
    fi
    return 1
}

is_arm64_host() {
    local machine
    machine="$(uname -m)"
    [ "$machine" = "aarch64" ] || [ "$machine" = "arm64" ]
}

# Does this host support SVE?
#
# The "Features" line in /proc/cpuinfo is the kernel's HWCAP/HWCAP2, decoded at
# boot from the CPU capability registers (ID_AA64ZFR0_EL1, ID_AA64PFR0_EL1).
# Userspace cannot read those registers directly, so this is the standard way to
# query SVE on Linux.
has_sve() {
    if [ -f "/proc/cpuinfo" ]; then
        grep "^Features" /proc/cpuinfo | head -1 | grep -qw "sve"
    else
        return 1
    fi
}

# Generation does not need arm64 *hardware* -- clang cross-compiles and
# asm2arm_tool only parses text assembly -- but it does need headers for the
# target. The natives include glibc via native/parsing.h, and cross-compiling
# without an aarch64 sysroot silently yields a 32-bit ssize_t and miscompiled
# natives. generate_native_go.sh locates a sysroot and hard-fails on a bad
# header environment, so the gate here is informational only.
check_build_host() {
    echo ">>> Build host: $(uname -m)"
    if ! is_arm64_host; then
        echo ">>> Cross-generating: an aarch64 sysroot is required (see generate_native_go.sh),"
        echo "    and this host cannot run the result. Test on arm64 via test_native_recover.sh."
    fi
}

# Running the generated natives needs the real thing.
check_test_host() {
    if ! is_arm64_host; then
        echo "Error: these tests execute the generated arm64 natives and cannot run on $(uname -m)."
        echo "       On a non-arm64 host 'go test' silently exercises the amd64 paths instead."
        exit 1
    fi
    if is_kunpeng_cpu; then
        echo "检测到鲲鹏CPU"
    elif has_sve; then
        echo ">>> SVE-capable arm64 CPU detected"
    else
        echo ">>> arm64 CPU without SVE detected: the neon suites are meaningful here,"
        echo "    but the sve_linkname/sve_wrapgoc suites will not exercise their natives."
    fi
}
