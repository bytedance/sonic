cmake_minimum_required(VERSION 3.16)
project(asm2arm_tool LANGUAGES CXX C)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

find_package(LLVM 19 REQUIRED CONFIG)
message(STATUS "Found LLVM ${LLVM_PACKAGE_VERSION}")
message(STATUS "Using LLVMConfig.cmake in ${LLVM_DIR}")
message(STATUS "Using LLVM libraries in ${LLVM_LIBRARY_DIRS}")

include_directories(${LLVM_INCLUDE_DIRS})
link_directories(${LLVM_LIBRARY_DIRS})
separate_arguments(LLVM_DEFINITIONS_LIST NATIVE_COMMAND ${LLVM_DEFINITIONS})
add_definitions(${LLVM_DEFINITIONS_LIST})
llvm_map_components_to_libnames(LLVM_LIBS all support option lto
AllTargetsInfos AllTargetsCodeGens AllTargetsAsmParsers AllTargetsDescs AllTargetsDisassemblers)

include(CheckCXXCompilerFlag)
include(CheckIncludeFileCXX)
check_cxx_compiler_flag("-march=armv8-a+sve" COMPILER_SUPPORTS_SVE)
if(COMPILER_SUPPORTS_SVE)
    check_include_file_cxx("arm_sve.h" HAVE_ARM_SVE_H)
    if(HAVE_ARM_SVE_H)
        set(ARCH_FLAGS "-march=armv8-a+sve")
        message(STATUS "SVE is supported, enabling SVE extensions")
    else()
        message(FATAL_ERROR "arm_sve.h header not found, SVE support is required")
    endif()
else()
    message(FATAL_ERROR "Compiler does not support SVE, SVE support is required")
endif()

add_executable(asm2arm_tool main.cpp streamer_JIT.cpp streamer_SL.cpp dump_elf.cpp utils.cpp cal_depth.cpp mc_bundle.cpp go_func_parser.cpp)
target_compile_definitions(asm2arm_tool PRIVATE ${LLVM_DEFINITIONS})
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fno-rtti -ffunction-sections -fdata-sections -g ${ARCH_FLAGS}")
set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -Wl,--gc-sections")
target_link_libraries(asm2arm_tool PRIVATE
    lldELF
    lldCommon

    ${LLVM_LIBS}

    pthread
    dl
    z
)