# GitHub Release 标记说明

[English](RELEASES.md) | 中文

## 概述

本文档解释了 GitHub Release 主页中不同标记类型的含义，以及它们对依赖管理工具的影响。

## Release 标记类型

### 1. Latest Release (最新版本)

**标记**: 带有 "Latest" 绿色标签的 release

**含义**:
- 这是项目的最新稳定版本
- GitHub 会自动将最新的非 pre-release 版本标记为 "Latest"
- 推荐用于生产环境

**对依赖管理的影响**:
- **Go Modules**: 当你运行 `go get github.com/bytedance/sonic@latest` 时，会获取标记为 "Latest" 的版本
- **Dependabot**: 默认情况下，Dependabot 会建议升级到 "Latest" 版本
- **其他依赖管理工具**: 大多数工具（如 Renovate）默认推荐 "Latest" 版本

### 2. Pre-release (预发布版本)

**标记**: 带有 "Pre-release" 橙色标签的 release

**含义**:
- 这是一个测试版本，可能包含实验性功能或未完全测试的代码
- 通常用于 alpha、beta、rc（候选版本）等
- 不推荐在生产环境中使用

**对依赖管理的影响**:
- **Go Modules**: 默认情况下，`go get` 命令**不会**获取 pre-release 版本，除非明确指定版本号
- **Dependabot**: 默认情况下，Dependabot **不会**建议升级到 pre-release 版本
- **语义化版本**: Pre-release 版本通常包含 `-alpha`、`-beta`、`-rc` 等后缀，如 `v1.2.0-beta.1`

**示例**:
```bash
# 不会获取 pre-release 版本
go get github.com/bytedance/sonic@latest

# 必须明确指定才能获取 pre-release 版本
go get github.com/bytedance/sonic@v1.2.0-beta.1
```

### 3. 无标记的 Release

**标记**: 既不是 "Latest" 也不是 "Pre-release" 的 release

**含义**:
- 这些是已发布的稳定版本，但不是最新的
- 通常是历史版本或被后续版本取代的版本

**对依赖管理的影响**:
- **Go Modules**: 可以通过明确指定版本号来获取，如 `go get github.com/bytedance/sonic@v1.0.0`
- **Dependabot**: 不会主动建议升级到旧版本
- **用途**: 适用于需要固定特定版本的场景

## 版本选择建议

### 对于库的使用者

1. **生产环境**: 始终使用 "Latest" 标记的版本
2. **测试新功能**: 可以尝试 "Pre-release" 版本，但要做好充分测试
3. **稳定性要求高**: 可以固定使用某个已验证的历史版本

### 对于库的维护者

1. **发布稳定版本**: 不要勾选 "Set as a pre-release" 选项
2. **发布测试版本**: 勾选 "Set as a pre-release" 选项，并使用语义化版本命名（如 `v1.2.0-beta.1`）
3. **版本号规范**: 遵循[语义化版本 2.0.0](https://semver.org/lang/zh-CN/)

## Go Modules 特定说明

### 版本选择规则

Go modules 使用以下规则选择版本：

1. 如果没有指定版本，`@latest` 会选择：
   - 最新的 tagged 版本（不包括 pre-release）
   - 如果没有 tags，则使用默认分支的最新 commit

2. Pre-release 版本的识别：
   - 包含 `-` 的版本被视为 pre-release（如 `v1.0.0-beta`）
   - 只有明确指定时才会被选择

### 示例

```bash
# 获取最新稳定版本
go get github.com/bytedance/sonic@latest

# 获取特定版本
go get github.com/bytedance/sonic@v1.0.0

# 获取 pre-release 版本
go get github.com/bytedance/sonic@v1.2.0-beta.1

# 升级到最新的 patch 版本
go get github.com/bytedance/sonic@v1.0
```

## Dependabot 配置

Dependabot 默认不会建议更新到 pre-release 版本。目前，Dependabot 没有直接的配置选项来启用 pre-release 版本更新。如果需要使用 pre-release 版本，你需要：

1. 手动更新 `go.mod` 文件中的版本号到 pre-release 版本
2. 使用其他依赖管理工具（如 Renovate），它们提供了更灵活的 pre-release 版本配置选项

基本的 Dependabot 配置示例：

```yaml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
```

## 常见问题

### Q: 如何知道当前使用的是否是最新版本？

A: 查看 `go.mod` 文件中的版本号，并与 GitHub Release 页面的 "Latest" 版本对比。

### Q: Pre-release 版本安全吗？

A: Pre-release 版本可能包含未经充分测试的代码，不建议在生产环境中使用。但可以在测试环境中试用新功能。

### Q: 如何固定某个版本避免自动更新？

A: 在 `go.mod` 中明确指定版本号，并在 CI/CD 中不使用 `go get -u`。

## 参考资料

- [GitHub Releases 文档](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases)
- [语义化版本规范](https://semver.org/lang/zh-CN/)
- [Go Modules 参考](https://go.dev/ref/mod)
- [Dependabot 配置参考](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file)
