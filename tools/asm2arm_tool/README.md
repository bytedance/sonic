# asm2arm_tool

asm2arm_tool 是一个用于将 ARM64 汇编代码转换为 Go 语言可调用代码的工具，支持 JIT 和静态链接两种模式。

## 构建与测试

### 工具依赖

- **LLVM19**：依赖特定版本的LLVM，需要在生成汇编时同时支持goframe与SVE指令，见**build_tool.sh**中的链接
- 测试依赖sonic的go环境

### 脚本说明

1. **build_tool.sh**：负责拉取LLVM、simde依赖并构建LLVM和工具本身
2. **generate_native_go.sh**：负责编译native目录下的C源码为汇编并输入给工具生成对应模式的Go文件
3. **test_native_recover.sh**：测试生成文件功能正确性的脚本
4. **test_encoder_api.sh**：测试encoder与api功能正确性的脚本

### 构建步骤

0. **构建依赖**

   编译器需要支持C++17与完整的SVE支持，推荐GCC9或者CLANG9以上版本

1. **构建工具**

   ```bash
   cd tools/asm2arm_tool/scripts

   bash build_tool.sh
   ```

   构建脚本会：
   - 克隆LLVM仓库到 `tools/asm2arm_tool/llvm-project` 
   - 在`tools/asm2arm_tool/build/llvm-build`构建LLVM（仅启用clang和lld）
   - 安装LLVM到 `tools/asm2arm_tool/build/llvm-install` 目录
   - 在`tools/asm2arm_tool/build`构建asm2arm_tool工具

2. **运行测试**

   `native`与`recover`测试：
   ```bash
   cd tools/asm2arm_tool/scripts

   bash test_native_recover.sh
   ```

   `test_native_recover.sh`测试脚本会：
   - 拷贝 `internal/native/neon` 下的文件到 `output/neon` 目录
   - 拷贝 `internal/native/sve_linkname` 下的文件到 `output/sve_linkname` 目录
   - 拷贝 `internal/native/sve_wrapgoc` 下的文件到 `output/sve_wrapgoc` 目录
   - 调用 `generate_native_go.sh` 生成对应平台的go代码文件到上述`output/neon`、`output/sve_linkname`、`output/sve_wrapgoc`目录
   - 在三个输出目录下分别运行go的测试

   `encoder`与`api`测试：
   ```bash
   cd tools/asm2arm_tool/scripts

   bash test_encoder_api.sh
   ```
   `test_encoder_api.sh`测试脚本会：
   - 将生成文件拷贝到internal/native下的目录中
   - 执行encoder、api的测试

### 注意事项

- 首次构建会拉取和编译LLVM，过程可能需要较长时间和较大磁盘空间
- 构建过程依赖网络连接（拉取LLVM和子模块）
- generate_native_go.sh脚本执行工具时，默认开启了--debug选项，会将debug输出重定向到对应的log文件


## 使用方法

### 命令行选项

#### 通用选项

| 选项 | 描述 | 必需 |
|------|------|------|
| `--mode` | 工具模式：`JIT` 或 `SL` | 是 |
| `--source` | 输入汇编文件路径 | 是 |
| `--output` | 生成的 Go 文件输出路径 | 是 |
| `--link-ld` | 链接脚本路径 | 是 |
| `--package` | 生成的 Go 文件所属的包名 | 是 |
| `--features` | 特性，如 `+sve,+aes` | 否 |
| `--debug` | 启用调试输出 | 否 |

#### JIT 模式选项

| 选项 | 描述 | 必需 |
|------|------|------|
| `--tmpl` | 模板文件路径 | 是 |

#### 静态链接 (SL) 模式选项

| 选项 | 描述 | 必需 |
|------|------|------|
| `--goproto` | Go 函数接口声明文件路径 | 是 |

### 使用示例

#### 输出帮助

```bash
cd tools/asm2arm_tool

./build/asm2arm_tool --help
```

#### JIT 模式

```bash
./build/asm2arm_tool --mode JIT --source input.s --output output_dir --link-ld scripts/link.ld --package sve_wrapgoc --tmpl input.tmpl
```

执行完成后，会在output_dir下生成对应的`input_subr.go`、`input_text_arm64.go`、`input.go`三个文件

#### 静态链接模式

```bash
./build/asm2arm_tool --mode SL --source input.s --output output_dir --link-ld scripts/link.ld --package sve_linkname --goproto input_arm64.go
```

执行完成后，会在output_dir下生成对应的`input_subr_arm64.go`、`input_arm64.s`两个文件

#### 汇编编译选项

输入给工具的汇编文件，在生成时可参考下列选项：

```bash
${CLANG_PATH} \
   -g0 -fverbose-asm -fstack-usage -fsigned-char -Wa,--no-size-directive -fno-ident -fno-jump-tables \
   -ffixed-x28 -ffixed-x9 -Wno-error -Wno-nullability-completeness -Wno-incompatible-pointer-types \
   -mllvm=--go-frame -mllvm=--enable-shrink-wrap=0 -mno-red-zone \
   -fno-stack-protector -nostdlib -O3 -fno-asynchronous-unwind-tables -fno-builtin -fno-exceptions \
   -march=armv8-a -I${SIMDE_INCLUDE_DIR} -S -o "${asm_file}" "${src_file}"
```

- clang在-O0、-O1的优化级别下，生成的汇编可能存在栈空间过大的非法汇编，推荐开启-O2及以上优化等级
