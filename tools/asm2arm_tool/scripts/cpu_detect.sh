#!/bin/bash

# 检测是否是鲲鹏CPU
is_kunpeng_cpu() {
    # 尝试从 /proc/cpuinfo 读取CPU信息
    if [ -f "/proc/cpuinfo" ]; then
        # 检查CPU part字段
        if grep -q "CPU part" "/proc/cpuinfo"; then
            local cpu_part=$(grep "CPU part" "/proc/cpuinfo" | head -1 | awk -F: '{print $2}' | tr -d ' ')
            if [ "$cpu_part" = "0xd02" ] || [ "$cpu_part" = "0xd06" ]; then
                return 0
            fi
        fi
    fi
    return 1
}

check_kunpeng_cpu() {
    if is_kunpeng_cpu; then
        echo "检测到鲲鹏CPU"
    else
        echo "不支持的CPU"
        exit 1
    fi
}
