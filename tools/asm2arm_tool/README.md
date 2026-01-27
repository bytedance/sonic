# asm2arm_tool

asm2arm_tool 是一个用于将 ARM64 汇编代码转换为 Go 语言可调用代码的工具，支持 JIT 和静态链接两种模式。

## 构建与测试

### 工具依赖

- **LLVM19**

### 脚本说明

1. **build.sh**：负责拉取LLVM、simde依赖并构建LLVM和工具本身
2. **build_go.sh**：负责编译native目录下的C源码为汇编并输入给工具生成对应模式的Go文件
3. **test.sh**：测试生成文件功能正确性的脚本

### 构建步骤

1. **构建工具**

   ```bash
   cd tools/asm2arm_tool/scripts

   bash build.sh
   ```

   构建脚本会：
   - 克隆LLVM仓库到 `tools/asm2arm_tool/llvm-project` 
   - 在`tools/asm2arm_tool/build/llvm-build`构建LLVM（仅启用clang和lld）
   - 安装LLVM到 `tools/asm2arm_tool/build/llvm-install` 目录
   - 在`tools/asm2arm_tool/build`构建asm2arm_tool工具

2. **运行测试**

   ```bash
   cd tools/asm2arm_tool/scripts

   bash test.sh
   ```

   测试脚本会：
   - 拷贝 `internal/native/neon` 下的文件到 `output/neon` 目录
   - 拷贝 `internal/native/sve_linkname` 下的文件到 `output/sve_linkname` 目录
   - 拷贝 `internal/native/sve_wrapgpc` 下的文件到 `output/sve_wrapgpc` 目录
   - 调用 `build_go.sh` 生成对应平台的go代码文件到上述`output/neon`、`output/sve_linkname`、`output/sve_wrapgpc`目录
   - 在三个输出目录下分别运行go的测试

### 注意事项

- 首次构建会拉取和编译LLVM，过程可能需要较长时间和较大磁盘空间
- 构建过程依赖网络连接（拉取LLVM和子模块）
- build_go.sh脚本执行工具时，默认开启了--debug选项，会将debug输出重定向到对应的log文件


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
./build/asm2arm_tool --mode JIT --source input.s --output output_dir --link-ld scripts/link.ld --package sve_wrapgpc --tmpl input.tmpl
```

执行完成后，会在output_dir下生成对应的`input_subr.go`、`input_text_arm64.go`、`input.go`三个文件

#### 静态链接模式

```bash
./build/asm2arm_tool --mode SL --source input.s --output output_dir --link-ld scripts/link.ld --package sve_linkname --goproto input_arm64.go
```

执行完成后，会在output_dir下生成对应的`input_subr_arm64.go`、`input_arm64.s`两个文件
