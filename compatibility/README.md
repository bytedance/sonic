# Go 1.27 JSON compatibility tests

This module records observed semantic differences between Sonic and the JSON
implementations available with Go 1.27. It does not change Sonic's production
behavior.

See [docs/sonic-go127-compatibility.md](../docs/sonic-go127-compatibility.md)
for the case matrix, interpretation rules, and release boundary.

## Files

| File | Responsibility |
|---|---|
| `catalog_test.go` | Sonic versus the active `encoding/json` backend |
| `jsonv2_test.go` | v1 facade versus experimental `encoding/json/v2` defaults |
| `backend_jsonv1_test.go` | Selects legacy v1 expectations |
| `backend_jsonv2_test.go` | Selects Go 1.27 default-backend expectations |

All Sonic-versus-standard-library comparisons call `sonic.ConfigStd`
directly. The only `ConfigDefault` calls are in tests that explicitly compare
Sonic presets.

## Run

```sh
# Go 1.27 default: encoding/json v1 API backed by jsonv2
env -u GOROOT GOTOOLCHAIN=go1.27rc2 \
  go test ./compatibility/ -count=1

# Go 1.27 legacy encoding/json backend
env -u GOROOT GOTOOLCHAIN=go1.27rc2 GOEXPERIMENT=nojsonv2 \
  go test ./compatibility/ -count=1

# Sonic optimized decoder, FastMap, and encoder VM
env -u GOROOT GOTOOLCHAIN=go1.27rc2 \
  SONIC_USE_OPTDEC=1 SONIC_USE_FASTMAP=1 SONIC_ENCODER_USE_VM=1 \
  go test ./compatibility/ -count=1
```

`compatibility/` is a separate workspace module. A root-level `go test ./...`
does not run it.
