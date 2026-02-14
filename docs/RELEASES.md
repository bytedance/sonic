# GitHub Release Labels Explanation

English | [中文](RELEASES_ZH_CN.md)

## Overview

This document explains the meaning of different label types on the GitHub Release homepage and their impact on dependency management tools.

## Release Label Types

### 1. Latest Release

**Label**: Release with a green "Latest" badge

**Meaning**:
- This is the latest stable version of the project
- GitHub automatically marks the most recent non-pre-release version as "Latest"
- Recommended for production use

**Impact on Dependency Management**:
- **Go Modules**: When you run `go get github.com/bytedance/sonic@latest`, it fetches the version marked as "Latest"
- **Dependabot**: By default, Dependabot suggests upgrading to the "Latest" version
- **Other Tools**: Most dependency management tools (like Renovate) default to recommending the "Latest" version

### 2. Pre-release

**Label**: Release with an orange "Pre-release" badge

**Meaning**:
- This is a test version that may contain experimental features or incompletely tested code
- Typically used for alpha, beta, rc (release candidate), etc.
- Not recommended for production use

**Impact on Dependency Management**:
- **Go Modules**: By default, `go get` commands will **NOT** fetch pre-release versions unless explicitly specified
- **Dependabot**: By default, Dependabot will **NOT** suggest upgrading to pre-release versions
- **Semantic Versioning**: Pre-release versions typically contain suffixes like `-alpha`, `-beta`, `-rc`, e.g., `v1.2.0-beta.1`

**Example**:
```bash
# Will NOT fetch pre-release versions
go get github.com/bytedance/sonic@latest

# Must explicitly specify to get a pre-release version
go get github.com/bytedance/sonic@v1.2.0-beta.1
```

### 3. Unmarked Releases

**Label**: Releases that are neither "Latest" nor "Pre-release"

**Meaning**:
- These are published stable versions, but not the latest
- Usually historical versions or versions superseded by newer releases

**Impact on Dependency Management**:
- **Go Modules**: Can be fetched by explicitly specifying the version number, e.g., `go get github.com/bytedance/sonic@v1.0.0`
- **Dependabot**: Will not proactively suggest upgrading to older versions
- **Use Case**: Suitable for scenarios requiring a fixed specific version

## Version Selection Recommendations

### For Library Users

1. **Production Environment**: Always use the "Latest" marked version
2. **Testing New Features**: Try "Pre-release" versions, but perform thorough testing
3. **High Stability Requirements**: Consider pinning to a verified historical version

### For Library Maintainers

1. **Publishing Stable Versions**: Do NOT check the "Set as a pre-release" option
2. **Publishing Test Versions**: Check the "Set as a pre-release" option and use semantic versioning (e.g., `v1.2.0-beta.1`)
3. **Version Numbering**: Follow [Semantic Versioning 2.0.0](https://semver.org/)

## Go Modules Specific Notes

### Version Selection Rules

Go modules use the following rules for version selection:

1. If no version is specified, `@latest` selects:
   - The latest tagged version (excluding pre-releases)
   - If no tags exist, uses the latest commit from the default branch

2. Pre-release version identification:
   - Versions containing `-` are treated as pre-releases (e.g., `v1.0.0-beta`)
   - Only selected when explicitly specified

### Examples

```bash
# Get the latest stable version
go get github.com/bytedance/sonic@latest

# Get a specific version
go get github.com/bytedance/sonic@v1.0.0

# Get a pre-release version
go get github.com/bytedance/sonic@v1.2.0-beta.1

# Upgrade to the latest patch version
go get github.com/bytedance/sonic@v1.0
```

## Dependabot Configuration

Dependabot does not suggest updates to pre-release versions by default. Currently, Dependabot does not have a direct configuration option to enable pre-release version updates. If you need to use pre-release versions, you can:

1. Manually update the version number in your `go.mod` file to the pre-release version
2. Use other dependency management tools (like Renovate) that provide more flexible pre-release version configuration options

Basic Dependabot configuration example:

```yaml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
```

## Frequently Asked Questions

### Q: How do I know if I'm using the latest version?

A: Check the version number in your `go.mod` file and compare it with the "Latest" version on the GitHub Release page.

### Q: Are pre-release versions safe?

A: Pre-release versions may contain insufficiently tested code and are not recommended for production environments. However, they can be used in testing environments to try new features.

### Q: How do I pin a specific version to avoid automatic updates?

A: Explicitly specify the version number in `go.mod` and avoid using `go get -u` in your CI/CD pipeline.

## References

- [GitHub Releases Documentation](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases)
- [Semantic Versioning Specification](https://semver.org/)
- [Go Modules Reference](https://go.dev/ref/mod)
- [Dependabot Configuration Reference](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file)
