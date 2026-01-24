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

add_executable(asm2arm_tool main.cpp streamer_JIT.cpp streamer_SL.cpp dump_elf.cpp utils.cpp cal_depth.cpp mc_bundle.cpp go_func_parser.cpp)
target_compile_definitions(asm2arm_tool PRIVATE ${LLVM_DEFINITIONS})
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fno-rtti -ffunction-sections -fdata-sections -g -march=armv8-a+sve")
set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -Wl,--gc-sections")
target_link_libraries(asm2arm_tool PRIVATE
    lldELF
    lldCommon

    ${LLVM_LIBS}

    pthread
    dl
    z
)