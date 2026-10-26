# Repository Guidelines

## Project Structure & Module Organization

Sonic is a Go JSON library with JIT, SIMD, and native-code acceleration. Public APIs live at the repository root and in `ast/`, `decoder/`, `encoder/`, `option/`, `unquote/`, and `utf8/`. Implementation details live under `internal/`, including JIT, native dispatch, runtime helpers, caching, and encoder/decoder internals. C sources are in `native/`; generated Go native wrappers are under `internal/native/`. The workspace in `go.work` includes the root module plus `compatibility/`, `external_jsonlib_test/`, `fuzz/`, `generic_test/`, `issue_test/`, and `loader/`. Test data and examples live in `testdata/`, `examples/`, `fuzz/corpus/`, and package-local `*_test.go` files.

## Build, Test, and Development Commands

- `go test ./...`: run the main module tests.
- `SONIC_USE_OPTDEC=1 SONIC_USE_FASTMAP=1 SONIC_ENCODER_USE_VM=1 go test ./...`: exercise the optimized/VM path used in CI.
- `cd loader && go test ./...`: test the loader submodule.
- `go test ./issue_test` and `go test ./generic_test`: run regression and compatibility suites from the workspace.
- `go test ./compatibility`: run the Go 1.27 `encoding/json` semantic drift catalog (see `docs/sonic-go127-compatibility.md`).
- `./scripts/test_race.sh`: run the expected race-detection regression check.
- `./scripts/fuzz.sh fuzz`, then `./scripts/fuzz.sh run` or `runopt`: initialize and run fuzzing.
- `./scripts/bench.sh`: run core encoder, decoder, AST, and external JSON benchmark suites.
- `./scripts/build-x86.sh [clang]`: regenerate x86 native outputs after changing `native/*.c`.

## Coding Style & Naming Conventions

Use `gofmt` for all Go files and follow Go Code Review Comments. Keep generated native output synchronized with its C source. Go tests use `TestXxx`, `BenchmarkXxx`, and `FuzzXxx`. Issue regressions follow `issue_test/issueNNN_test.go`. Shell scripts should be executable, use bash consistently, and keep environment toggles explicit. Write repository documentation in English by default; use another language only when the target audience or existing document explicitly requires it.

## Testing Guidelines

Add focused package tests beside changed Go code. For parser, encoder, JIT, SIMD, or loader changes, run both default and optimized/VM paths. For user-reported behavior fixes, add or update an `issue_test` case. For fuzz-sensitive parsing changes, seed `fuzz/corpus/` or `fuzz/testdata/` when useful.

## Commit & Pull Request Guidelines

History and the PR template use Conventional Commits, for example `feat: support Go 1.26`, `fix(decoder): handle ...`, `opt: ast node set`, `ci: add ...`, or `chore: update go mod`. Branch names should use prefixes such as `feature/`, `bugfix/`, `optimize/`, `doc/`, `ci/`, `test/`, or `refactor/`; validate with `./scripts/check_branch_name.sh`. PRs should describe user-visible impact, link issues when applicable, include tests or benchmark data for performance changes, and update user documentation when behavior changes.

## Security & Compatibility Notes

Do not disclose security bugs in public issues; follow `CONTRIBUTING.md` and contact the maintainers. Supported Go versions are documented in `README.md`; be careful with linkname behavior on newer Go releases and use `scripts/go_flags.sh` patterns when needed.
