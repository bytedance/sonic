#!/bin/bash

set -e

# 清理选项
CLEAN="false"

# 解析命令行参数
while getopts "c" opt; do
  case $opt in
    c)
      CLEAN="true"
      ;;
    *)
      echo "Usage: $0 [-c]"
      echo "  -c: Clean generated files (*.o, *.elf, *.log)"
      exit 1
      ;;
  esac
done

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TOOL_DIR="$(dirname "${SCRIPT_DIR}")"   # asm2arm_tool
PROJECT_DIR="$(dirname $(dirname "${TOOL_DIR}"))"   # sonic

BUILD_DIR="${TOOL_DIR}/build"
TOOL_PATH="${BUILD_DIR}/asm2arm_tool"
LLVM_INSTALL_DIR="${BUILD_DIR}/llvm-install"
CLANG_PATH="${LLVM_INSTALL_DIR}/bin/clang"

SRC_DIR="${PROJECT_DIR}/native"
TMPL_DIR="${PROJECT_DIR}/internal/native"
OUTPUT_DIR="${TOOL_DIR}/output"

# 清理函数
function clean_files() {
  echo ">>> Cleaning generated files..."
  
  # 清理 neon 目录下的 .o 和 .log 文件
  find "${OUTPUT_DIR}/neon" -name "*.o" -type f -delete 2>/dev/null || true
  find "${OUTPUT_DIR}/neon" -name "*.log" -type f -delete 2>/dev/null || true
  
  # 清理 sve_linkname 目录下的 .o 和 .log 文件
  find "${OUTPUT_DIR}/sve_linkname" -name "*.o" -type f -delete 2>/dev/null || true
  find "${OUTPUT_DIR}/sve_linkname" -name "*.log" -type f -delete 2>/dev/null || true
  
  # 清理 sve_wrapgoc 目录下的 .o 和 .log 文件
  find "${OUTPUT_DIR}/sve_wrapgoc" -name "*.o" -type f -delete 2>/dev/null || true
  find "${OUTPUT_DIR}/sve_wrapgoc" -name "*.log" -type f -delete 2>/dev/null || true
  
  echo ">>> Clean completed!"
}

# 创建输出目录
mkdir -p "${OUTPUT_DIR}/neon"
mkdir -p "${OUTPUT_DIR}/sve_linkname"
mkdir -p "${OUTPUT_DIR}/sve_wrapgoc"
mkdir -p "${OUTPUT_DIR}/asm/neon"
mkdir -p "${OUTPUT_DIR}/asm/sve"

NEON_OUTPUT="${OUTPUT_DIR}/neon"
SVE_LINKNAME_OUTPUT="${OUTPUT_DIR}/sve_linkname"
SVE_WRAPGOC_OUTPUT="${OUTPUT_DIR}/sve_wrapgoc"
NEON_ASM_DIR="${OUTPUT_DIR}/asm/neon"
SVE_ASM_DIR="${OUTPUT_DIR}/asm/sve"

# 检查simde子模块是否已拉取
echo ">>> Checking simde submodule..."
SIMDE_DIR="${PROJECT_DIR}/tools/simde"
SIMDE_INCLUDE_DIR="${PROJECT_DIR}/tools/simde/simde"
if [ ! -d "${SIMDE_DIR}" ] || [ ! -d "${SIMDE_DIR}/.git" ]; then
    echo ">>> simde submodule not found or not initialized."
    echo ">>> Initializing and updating simde submodule..."
    cd "${PROJECT_DIR}"
    git submodule update --init --recursive tools/simde
    if [ $? -ne 0 ]; then
        echo "Error: Failed to initialize simde submodule."
        exit 1
    fi
    cd "${SCRIPT_DIR}"
else
    echo ">>> simde submodule is already initialized."
fi

# 检查simde头文件是否存在
if [ ! -d "${SIMDE_INCLUDE_DIR}" ]; then
    echo "Error: simde include directory not found: ${SIMDE_INCLUDE_DIR}"
    echo "Please check if simde submodule is correctly initialized."
    exit 1
fi
echo ">>> simde include directory: ${SIMDE_INCLUDE_DIR}"
echo ""

echo ">>> Using ${CLANG_PATH} compiler"
echo ">>> Tool path: ${TOOL_PATH}"
echo ">>> Output directory: ${OUTPUT_DIR}"

# 检查工具是否存在
if [ ! -f "${TOOL_PATH}" ]; then
    echo "Error: Tool not found. Please run build_tool.sh first."
    exit 1
fi

# 检查clang是否存在
if [ ! -f "${CLANG_PATH}" ]; then
    echo "Error: Clang not found. Please run build_tool.sh first."
    exit 1
fi

# 遍历native目录下的.c文件
echo ""  
echo ">>> Processing native directory files..."
if [ -d "${SRC_DIR}" ]; then
    for src_file in "${SRC_DIR}"/*.c; do
        if [ -f "${src_file}" ]; then
            base_name="$(basename "${src_file}" .c)"
            
            echo ""
            echo ">>> Processing ${src_file}..."
            
            # 处理neon目录
            NEON_FILE="${PROJECT_DIR}/internal/native/neon/${base_name}_arm64.go"
            if [ -f "${NEON_FILE}" ]; then
                echo ""
                echo ">>> Processing for neon..."
                asm_file="${NEON_ASM_DIR}/${base_name}.s"
                cerr_log="${NEON_OUTPUT}/${base_name}.log"
                
                # 编译生成汇编文件（neon版本）
                echo ">>> Compiling to assembly (neon)... --> ${asm_file}"
                ${CLANG_PATH} \
                -g0 -fverbose-asm -fstack-usage -fsigned-char -Wa,--no-size-directive -fno-ident -fno-jump-tables \
                -ffixed-x28 -ffixed-x9 -Wno-error -Wno-nullability-completeness -Wno-incompatible-pointer-types \
                -mllvm=--go-frame -mllvm=--enable-shrink-wrap=0 -mno-red-zone \
                -fno-stack-protector -nostdlib -O3 -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions \
                -march=armv8-a+simd -I${SIMDE_INCLUDE_DIR} -S -o "${asm_file}" "${src_file}"
                
                # 检查汇编文件是否生成
                if [ ! -f "${asm_file}" ]; then
                    echo "Error: Assembly file not generated for neon."
                else
                    echo ">>> Execute SL mode for neon..."
                    ${TOOL_PATH} --debug --mode=SL --source=${asm_file} --goproto=${NEON_FILE} --output=${NEON_OUTPUT} --link-ld=${SCRIPT_DIR}/link.ld \
                    --package=neon 2>${cerr_log}
                    
                    if [ $? -eq 0 ]; then
                        echo ">>> Tool execution succeeded for neon ${base_name}"
                    else
                        echo ">>> Warning: Tool execution failed for neon ${base_name}. Check ${cerr_log} for details."
                    fi
                fi
            fi
            
            # 处理sve_linkname目录
            SVE_LINKNAME_FILE="${PROJECT_DIR}/internal/native/sve_linkname/${base_name}_arm64.go"
            if [ -f "${SVE_LINKNAME_FILE}" ]; then
                echo ""
                echo ">>> Processing for sve_linkname..."
                asm_file="${SVE_ASM_DIR}/${base_name}.s"
                cerr_log="${SVE_LINKNAME_OUTPUT}/${base_name}.log"

                echo ">>> Compiling to assembly (sve)... --> ${asm_file}"
                ${CLANG_PATH} \
                -g0 -fverbose-asm -fstack-usage -fsigned-char -Wa,--no-size-directive -fno-ident -fno-jump-tables \
                -ffixed-x28 -ffixed-x9 -Wno-error -Wno-nullability-completeness -Wno-incompatible-pointer-types\
                -mllvm -disable-constant-hoisting -mllvm=--go-frame -fno-addrsig -no-integrated-as \
                -mno-red-zone -fno-stack-protector -nostdlib -O3 -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions \
                -march=armv8-a+sve+aes -I${SIMDE_INCLUDE_DIR} -D__SVE__ -S -o "${asm_file}" "${src_file}"
                
                # 检查汇编文件是否生成
                if [ ! -f "${asm_file}" ]; then
                    echo "Error: Assembly file not generated for sve_linkname."
                else
                    echo ">>> Execute SL mode for sve_linkname..."
                    ${TOOL_PATH} --debug --mode=SL --source=${asm_file} --goproto=${SVE_LINKNAME_FILE} --output=${SVE_LINKNAME_OUTPUT} --link-ld=${SCRIPT_DIR}/link.ld \
                    --package=sve_linkname --features=+sve,+aes 2>${cerr_log}
                    
                    if [ $? -eq 0 ]; then
                        echo ">>> Tool execution succeeded for sve_linkname ${base_name}"
                    else
                        echo "Warning: Tool execution failed for sve_linkname ${base_name}. Check ${cerr_log} for details."
                    fi
                fi
            fi
            
            # 处理sve_wrapgoc目录
            SVE_WRAPGOC_FILE="${PROJECT_DIR}/internal/native/sve_wrapgoc/${base_name}.go"
            SVE_WRAPGOC_TMPL="${PROJECT_DIR}/internal/native/${base_name}.tmpl"
            if [ -f "${SVE_WRAPGOC_FILE}" ] && [ -f "${SVE_WRAPGOC_TMPL}" ]; then
                echo ""
                echo ">>> Processing for sve_wrapgoc..."
                asm_file="${SVE_ASM_DIR}/${base_name}.s"
                cerr_log="${SVE_WRAPGOC_OUTPUT}/${base_name}.log"
                
                echo ">>> Compiling to assembly (sve)... --> ${asm_file}"
                ${CLANG_PATH} \
                -g0 -fverbose-asm -fstack-usage -fsigned-char -Wa,--no-size-directive -fno-ident -fno-jump-tables \
                -ffixed-x28 -ffixed-x9 -Wno-error -Wno-nullability-completeness -Wno-incompatible-pointer-types\
                -mllvm -disable-constant-hoisting -mllvm=--go-frame -fno-addrsig -no-integrated-as \
                -mno-red-zone -fno-stack-protector -nostdlib -O3 -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions \
                -march=armv8-a+sve+aes -I${SIMDE_INCLUDE_DIR} -D__SVE__ -S -o "${asm_file}" "${src_file}"
                
                # 检查汇编文件是否生成
                if [ ! -f "${asm_file}" ]; then
                    echo "Error: Assembly file not generated for sve_wrapgoc."
                else
                    echo ">>> Execute JIT mode for sve_wrapgoc..."
                    ${TOOL_PATH} --debug --mode=JIT --source=${asm_file} --output=${SVE_WRAPGOC_OUTPUT} --link-ld=${SCRIPT_DIR}/link.ld --tmpl=${SVE_WRAPGOC_TMPL} \
                    --package=sve_wrapgoc --features=+sve,+aes 2>${cerr_log}
                    
                    if [ $? -eq 0 ]; then
                        echo ">>> Tool execution succeeded for sve_wrapgoc ${base_name}"
                    else
                        echo "Warning: Tool execution failed for sve_wrapgoc ${base_name}. Check ${cerr_log} for details."
                    fi
                fi
            fi
        fi
    done
else
    echo "Warning: native directory not found: ${SRC_DIR}"
fi

echo ""
echo ">>> All files processed!"
echo ">>> Output files are in: ${OUTPUT_DIR}"

# 如果指定了清理选项，执行清理
if [ "$CLEAN" = "true" ]; then
  clean_files
fi
