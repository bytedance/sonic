# Sonic and Go 1.27 Compatibility

Last updated: July 30, 2026
Test module: [`compatibility/`](../compatibility/)

## Status

Sonic treats Go 1.27 as a supported target. The current implementation and
validation evidence use Go 1.27rc2; the Go project still marks the 1.27 release
notes as draft and expects the final release in August 2026.

The support claim currently means:

- Sonic builds and passes its main test suites with Go 1.27rc2.
- JIT, OptDec, FastMap, encoder VM, and loader paths are exercised.
- Runtime ABI changes used by the loader and map iterator have been reviewed.
- Both Go 1.27 `encoding/json` backends are covered.

It does not mean that unreleased runtime internals are frozen. The final
toolchain must be validated again before the release-candidate qualifier can
be removed from this evidence.

## Compatibility Surfaces

Go 1.27 affects Sonic in four separate areas:

| Surface | Sonic adaptation | Current evidence | Remaining risk |
|---|---|---|---|
| JIT and native build tags | Go 1.27 uses the supported implementation instead of unsupported stubs | Default and optimized main-module tests pass | Final compiler or linker changes |
| Loader runtime ABI | `loader/funcdata_go127.go` follows the Go 1.27 `moduledata` layout | Loader race tests pass | `moduledata` is private runtime state |
| Map iteration | `GoMapIterator` retains legacy trailing storage | Go 1.25 noswissmap and Go 1.27 tests pass | Linkname shims are not public APIs |
| JSON semantics | Default jsonv2 backend and legacy backend are tested separately | Compatibility and issue suites pass in both modes | RC behavior may change before final |

## Runtime ABI Notes

### `moduledata`

The Go 1.27 runtime layout adds `typedesclen`, `itaboffset`, and `itabsize`,
removes the old `typelinks` and `itablinks` layout, and changes `typemap` to
`map[*_type]*_type`. Sonic mirrors that private layout in
`loader/funcdata_go127.go`.

This is an ABI adaptation, not a public Go API integration. A passing loader
test is necessary but does not replace a source-level comparison against the
final Go runtime.

### Map iterator storage

`internal/rt.GoMapIterator` keeps a trailing `ClearSeq uint64` field:

- Go 1.24 and Go 1.25 with `GOEXPERIMENT=noswissmap` require `clearSeq` in the
  legacy `hiter` layout.
- Go 1.27 provides legacy linkname shims for pre-Go 1.24 iterator callers.
- The newer shim does not need the trailing field, but safely ignores it.

`ClearSeq` therefore remains for older supported toolchains. It is not a new
Go 1.27 field.

## JSON Compatibility Model

The tests distinguish four entry points:

| Name | Entry point | Role |
|---|---|---|
| SonicStd | `sonic.ConfigStd` | Classic `encoding/json` v1 compatibility target |
| GoStd127 | `encoding/json` on Go 1.27 | v1 API backed by jsonv2 by default |
| LegacyStd | `encoding/json` with `GOEXPERIMENT=nojsonv2` | Temporary legacy v1 backend |
| JSONV2 | `encoding/json/v2` | New API with stricter defaults |

GoStd127 and JSONV2 are not interchangeable. GoStd127 keeps the v1 API and is
intended to preserve v1 marshal and unmarshal behavior. The direct JSONV2 API
intentionally uses new defaults and is not a `ConfigStd` compatibility target.

Sonic continues to target classic v1 behavior. It does not add a production
switch to reproduce release-candidate backend quirks. A future v2-style Sonic
API, if needed, should be a separate opt-in contract.

## Test Layout

The standalone compatibility module has four test files:

```text
compatibility/
  catalog_test.go          # Sonic versus the active encoding/json backend
  jsonv2_test.go           # v1 facade versus encoding/json/v2 defaults
  backend_jsonv1_test.go   # !goexperiment.jsonv2 expectations
  backend_jsonv2_test.go   # goexperiment.jsonv2 expectations
```

The two backend files only select the `stdUsesJSONV2` constant through build
tags. No runtime version parser, profile registry, or forwarding layer is
required.

`jsonv2_test.go` is compiled only with `go1.27 && goexperiment.jsonv2`, so
older toolchains and `GOEXPERIMENT=nojsonv2` never resolve
`encoding/json/v2`.

Every Sonic-versus-standard-library assertion in `catalog_test.go` calls
`sonic.ConfigStd` directly. `ConfigDefault` appears only in the two preset
tests that explicitly compare Sonic configurations.

## GoStd127 Compatibility Differences

These results were observed with Go 1.27rc2. They are regression evidence, not
a promise that the final Go 1.27 implementation will keep the same behavior.

| Case and test | Input or type | SonicStd | GoStd127 | LegacyStd | Compatibility impact |
|---|---|---|---|---|---|
| [`F32-OVF`](../compatibility/catalog_test.go#L40) | Value above float32 range | Error; destination remains `0` | Error; destination becomes `+Inf` | Error; destination remains `0` | Error-time mutation |
| [`F64-OVF`](../compatibility/catalog_test.go#L58) | `1e1000` | Error; destination remains `0` | Error; destination becomes `+Inf` | Error; destination remains `0` | Error-time mutation |
| [`F64-LONG-OVF`](../compatibility/catalog_test.go#L793) | 1,712-digit mantissa followed by `e-913` | No error; incorrectly returns a finite value near `5.33e-114` | Range error; element becomes `+Inf` | Range error; element remains `nil` | Historical range-detection bug exposed by Go 1.27 |
| [`MAP-F64-ENC`](../compatibility/catalog_test.go#L80) | `map[float64]string` | Unsupported | Encodes successfully | Unsupported | Accepted type set |
| [`MAP-PTR-ENC`](../compatibility/catalog_test.go#L98) | `map[*int]*int` | Unsupported | Encodes successfully | Unsupported | Accepted type set |
| [`MAP-F64-DEC`](../compatibility/catalog_test.go#L118) | Float map keys | Decodes successfully | Decodes successfully | Error | Accepted type set |
| [`STR-NULL`](../compatibility/catalog_test.go#L144) | `,string` pointer fields receive `"null"` | Nil pointers; no error | `*string` path returns a type error | Matches SonicStd | Error and pointer state |
| [`UTF8-ENC`](../compatibility/catalog_test.go#L173) | Go string contains invalid UTF-8 | Escaped `\ufffd` | Literal U+FFFD in output | Matches SonicStd | Byte-level output |
| [`TMPTR-UNJ`](../compatibility/catalog_test.go#L219) | Named pointer whose element implements `UnmarshalJSON` | Uses default field decoding | Calls `UnmarshalJSON` | Matches SonicStd | Method dispatch |
| [`UNJ-PARTIAL`](../compatibility/catalog_test.go#L264) | Custom unmarshal fails before field `E` | JIT keeps `E=0`; OptDec writes `E=1` | Writes `E=1` | Keeps `E=0` | Error-time partial writes |
| [`TM-KEY-QUOTE`](../compatibility/catalog_test.go#L303) | `MarshalText` returns quoted text | Quotes delimit the key | Quotes become key content and are re-escaped | Matches SonicStd | Encoded map-key bytes |
| [`IFACE-SKIP-NONPTR`](../compatibility/catalog_test.go#L355) | Interface holds a non-pointer value | Error; preserves the dynamic type | Type error; interface becomes nil | Matches SonicStd | Interface state |
| [`IFACE-SKIP-NILPTR`](../compatibility/catalog_test.go#L379) | Interface holds a typed nil | Error; preserves the typed nil | Type error; interface becomes nil | Matches SonicStd | Interface state |
| [`SELF-REF`](../compatibility/catalog_test.go#L410) | Self-referential `interface{}` | Replaces the cycle with a map | Not called because rc2 recurses to stack overflow | Not compared | Process availability |

`UNJ-PARTIAL`, `SYN-PREMUT`, `STRING-GONUM`, `STRING-GONUM-PLUS`, and
`BASE64-ERROR` have explicit decoder-dependent expectations. They read
`internal/envs.UseOptDec` and pin both JIT and OptDec behavior.

The most consequential differences are error-time mutation, interface state,
and self-reference handling. They can change application state even when both
implementations return an error. Accepted map-key types and method dispatch
change whether an operation succeeds. UTF-8 and quoted map-key cases primarily
affect byte-for-byte comparisons, signatures, snapshots, and cache keys.

### Long-mantissa float overflow

`F64-LONG-OVF` is a correctness difference, not merely an error-time mutation.
The test number is approximately `5.33e798`, far above the maximum finite
`float64`. Go 1.18 through Go 1.26 and Sonic releases from v1.0.0 through
v1.15.2 instead returned a finite value near `5.33e-114` without an error.

Both implementations historically used an 800-byte decimal fallback buffer.
When the 1,712-digit integer mantissa was truncated, the fallback computed the
decimal point from the retained 800 digits, turning the exponent into roughly
`800 - 913 = -113`. Go 1.27's
[`strconv` scaling change](https://github.com/golang/go/commit/71300e80113c6ca56105aac524e9c1b0db43910f)
detects the full decimal magnitude before that fallback and returns
`ErrRange`. The Go 1.27 jsonv2 backend stores `+Inf` before returning the JSON
type error; `GOEXPERIMENT=nojsonv2` returns the same error without storing a
value.

This MR records the gap but intentionally does not change Sonic's decoding
semantics. A follow-up fix should preserve the full decimal-point position or
reject the value before the bounded fallback, then evaluate the compatibility
impact for users that currently receive the historical finite result.

## Additional Gaps Found by the v1 Source Audit

Auditing every option in `DefaultOptionsV1` found five compatibility cases that
were not represented in the original catalog:

| Case and test | SonicStd | GoStd127 | LegacyStd | Classification |
|---|---|---|---|---|
| [`TEXT-APPEND`](../compatibility/catalog_test.go#L774) | Calls `MarshalText`, producing `"MARSHAL"` | Prefers `AppendText`, producing `"APPEND"` | Calls `MarshalText` | Go 1.27 accepted behavior change |
| [`STRING-GONUM`](../compatibility/catalog_test.go#L578) | Independent subtests show JIT rejects `"00012"` and `"0x1_4p-2"`; OptDec decodes them as `12` and `5` | Decodes them as `12` and `5` | Same as GoStd127 | Sonic JIT v1 gap |
| [`SYN-PREMUT`](../compatibility/catalog_test.go#L537) | JIT writes `A=1` before the later syntax error; OptDec leaves the destination unchanged | Leaves the destination unchanged | Same as GoStd127 | Sonic JIT v1 gap |
| [`MAP-VALUE-REPLACE`](../compatibility/catalog_test.go#L561) | Merges an object into the existing map value | Replaces the map value before decoding | Same as GoStd127 | Existing SonicStd v1 gap |
| [`MARSHAL-ERR-WRAP`](../compatibility/catalog_test.go#L621) | Returns the error from `MarshalJSON` directly | Wraps it in `*json.MarshalerError` | Wraps it in `*json.MarshalerError` | Existing SonicStd error-contract gap |

`TEXT-APPEND` is the only new GoStd127-versus-LegacyStd difference in this
group. Go [issue #79852](https://github.com/golang/go/issues/79852) considered
whether the v1 facade should ignore `encoding.TextAppender`; it was closed
without changing Go 1.27. Sonic must therefore choose whether v2-oriented APIs
should recognize `AppendText`, but `ConfigStd` can retain the legacy method
order as a documented compatibility choice.

`STRING-GONUM` is confirmed by Go
[issue #75619](https://github.com/golang/go/issues/75619). The v1 API uses
`strconv` grammar inside `,string`, including zero padding, underscores, and
hexadecimal floating-point syntax. The jsonv2-backed facade was changed in
[CL 709615](https://go.dev/cl/709615) to preserve that behavior. Sonic still
parses the quoted content as a JSON number.

The remaining three cases follow directly from documented v1 options:

- [`ReportErrorsWithLegacySemantics`](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2_options.go#L436-L467)
  requires complete syntactic validation before semantic mutation and preserves
  v1 error wrapping.
- [`MergeWithLegacySemantics`](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2_options.go#L349-L376)
  specifies replacement of map values rather than recursive merge.

These are not reasons to add Go 1.27 build-tag behavior. They are concrete
`ConfigStd` correctness gaps that exist under both Go backends. Two of them are
specific to Sonic's JIT decoder and already align under OptDec. Syntax
prevalidation and map-value replacement have the highest application impact
because successful earlier writes can survive an error or stale state can
survive a successful decode.

### Additional GoStd127-versus-LegacyStd Drift

A second source-guided tri-diff found behavior that changes inside the v1 API
when only the backend is switched:

| Case and test | SonicStd | GoStd127 | LegacyStd | Impact |
|---|---|---|---|---|
| [`STRING-GONUM-PLUS`](../compatibility/catalog_test.go#L639) | Independent subtests show JIT rejects leading `+` and `.5`; OptDec accepts them | Accepts them | Rejects them | Accepted numeric grammar |
| [`MAP-STRING-TEXT`](../compatibility/catalog_test.go#L685) | Uses the underlying named-string key | Calls `MarshalText` | Uses the underlying key | Method dispatch and map-key bytes |
| [`BASE64-ERROR`](../compatibility/catalog_test.go#L702) | Clears the byte slice; JIT returns `base64.CorruptInputError`, OptDec returns a mismatch error | Preserves it and wraps the cause in `*json.UnmarshalTypeError` | Preserves it and returns `base64.CorruptInputError` directly | Destination state and error type |
| [`MARSHAL-ERR-WRAP`](../compatibility/catalog_test.go#L621) | Returns the original error | `MarshalerError.Type` is `*T` | `MarshalerError.Type` is `T` | Observable error metadata |
| [`UNSUPPORTED-VALUE-PAYLOAD`](../compatibility/catalog_test.go#L723) | `Value` remains populated | `Value` is an invalid `reflect.Value` | `Value` contains the unsupported value | Observable error metadata |
| [`UNMARSHAL-TYPE-ROOT`](../compatibility/catalog_test.go#L744) | Returns Sonic's own mismatch error | `Struct` names the root destination type | `Struct` names the immediate containing type | Error classification and logging |

The `,string` behavior follows the same `strconv` compatibility path discussed
in [#75619](https://github.com/golang/go/issues/75619), but leading `+` and
leading decimal points demonstrate an actual GoStd127-versus-LegacyStd split.
The nested error-type change is related to the long-standing ambiguity in
[#43126](https://github.com/golang/go/issues/43126). No exact upstream issue was
found for the named-string map-key dispatch, Base64 wrapper, pointer-valued
`MarshalerError.Type`, or empty `UnsupportedValueError.Value`.

GoStd127 also adds an `Err` field and `Unwrap` method to
`encoding/json.UnmarshalTypeError`; those members are absent when the same
toolchain is built with `GOEXPERIMENT=nojsonv2`. This is a conditional public
API difference, not merely a runtime error-text change. Sonic should not try to
mirror backend-dependent error fields. A stable Sonic error contract, with
documented `errors.Is` and `errors.As` behavior, is safer than exposing a
`Go127Semantics` switch.

### API Surface Gap

Sonic's public [`Decoder`](../api.go#L159) supports `Decode`, `Buffered`,
`DisallowUnknownFields`, `More`, and `UseNumber`, but does not expose the v1
`Token` or `InputOffset` methods. Its internal stream decoder has
`InputOffset`, but callers receiving the public interface cannot use it. This
is an existing API-coverage gap rather than a Go 1.27 semantic regression.
Supporting direct jsonv2 would additionally require deciding whether Sonic
offers a token/value streaming layer comparable to `encoding/json/jsontext`;
adding configuration flags alone would not close that surface difference.

### Resolution Triage

The cases should not all be handled by making `ConfigStd` match GoStd127:

| Action | Cases | Reason |
|---|---|---|
| Fix in `ConfigStd` | `SYN-PREMUT` (JIT), `MAP-VALUE-REPLACE`, `STRING-GONUM` (JIT), `BASE64-ERROR` destination state, `MARSHAL-ERR-WRAP` | GoStd127 and LegacyStd agree on the core v1 behavior, or Sonic's JIT and OptDec disagree. These are Sonic compatibility or internal-consistency gaps. |
| Keep Sonic's stable v1 behavior | `F32-OVF`, `F64-OVF`, `STR-NULL`, `UTF8-ENC`, `TMPTR-UNJ`, `TM-KEY-QUOTE`, interface clearing, `SELF-REF` | Sonic matches LegacyStd or has safer behavior. Matching GoStd127 would silently change `ConfigStd` or copy an unsafe implementation artifact. |
| Record GoStd127 drift; consider only for an opt-in v2 preset | broader map keys, `TEXT-APPEND`, `STRING-GONUM-PLUS`, `MAP-STRING-TEXT` | These differ between GoStd127 and LegacyStd. They should not change the classic v1 preset by toolchain version. |
| Do not chase backend-specific error metadata | pointer-valued `MarshalerError.Type`, empty `UnsupportedValueError.Value`, root-valued `UnmarshalTypeError.Struct`, conditional `UnmarshalTypeError.Err` | The Go 1.27 default and legacy backends already disagree. Sonic should document one stable error contract instead. |
| Evaluate as separate API work | `Decoder.Token`, public `Decoder.InputOffset`, jsontext-style streaming | These are API-surface gaps, not fixes for the Go 1.27 runtime backend. |

The first implementation batch should focus on destination correctness:
`MAP-VALUE-REPLACE`, `SYN-PREMUT`, and `BASE64-ERROR`. They can leave stale or
partially modified application state. Numeric grammar and error wrapping are
important compatibility work but have lower state-corruption risk.

## Official Upstream Evidence

### How Go Preserves v1 Semantics

The accepted [encoding/json/v2 proposal](https://github.com/golang/go/issues/71497)
does not retain two permanent implementations. Instead, `encoding/json` calls
the jsonv2 engine with
[`DefaultOptionsV1`](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2_options.go#L197-L229),
while `encoding/json/v2` uses the new defaults. Individual options can override
parts of either preset. The
[Go 1.27 release notes](https://go.dev/doc/go1.27#encodingjsonv2) state that v1
marshal and unmarshal behavior is preserved, while exact error text may differ.
They also describe `GOEXPERIMENT=nojsonv2` as a temporary implementation
fallback that is expected to be removed, not a long-term semantic selection
mechanism.

This gives upstream three ways to handle a reported difference:

1. fix a v1-facade regression when `DefaultOptionsV1` can reproduce the old
   behavior;
2. accept a behavior change when the old behavior was inconsistent, obscure,
   or only changed non-semantic output;
3. keep the stricter behavior in the direct v2 API and expose an option for
   callers that need v1 behavior.

### Disposition of the Sonic Catalog Cases

The table distinguishes an exact issue from the nearest official design or
source evidence. “No exact issue found” means that the cited issue does not
fully cover Sonic's reproducer and the case should not be described as fixed
upstream.

| Sonic cases | Official evidence and status | Upstream disposition |
|---|---|---|
| `F32-OVF`, `F64-OVF` | [#77666](https://github.com/golang/go/issues/77666) is closed and implemented by [CL 772360](https://go.dev/cl/772360). The working group chose `strconv.ParseFloat`-like results: an out-of-range float returns `+Inf` or `-Inf` together with an error. The rc2 [decoder assigns the parsed value before returning its range error](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2/arshal_default.go#L737-L767). | The numeric conversion rule is intentional v2 behavior. #77666 concerns token accessors, not v1 destination mutation; no exact issue was found for the `0` to `+Inf` mutation. Treat that mutation as an unresolved v1-facade difference. |
| `MAP-F64-ENC`, `MAP-PTR-ENC`, `MAP-F64-DEC` | [#79938](https://github.com/golang/go/issues/79938) reported broader map-key acceptance and was closed as not planned. A maintainer explained that v2 defers an invalid-key error until an actual element is encoded. The [v2 contract](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2/arshal.go#L124-L127) recursively encodes keys and only requires their JSON representation to be a string. | Broader map-key handling is part of the v2 model. #79938 only partially covers these exact float and pointer cases, and `DefaultOptionsV1` has no option that restores the legacy key-type allowlist. Current facade behavior is therefore accepted in practice, but not explicitly promised for these three inputs. |
| `STR-NULL` | [`StringifyWithLegacySemantics`](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2_options.go#L493-L520) and [`MergeWithLegacySemantics`](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2_options.go#L349-L376) are both in `DefaultOptionsV1`. | The official design intends to preserve v1 `,string` and null-merge behavior. No exact issue was found for the `*string` result, so the rc2 type error and pointer mutation remain an unreported v1-facade regression candidate. |
| `UTF8-ENC` | [#75163](https://github.com/golang/go/issues/75163) is closed as completed. The working group deliberately kept literal U+FFFD rather than escaped `\ufffd`; [CL 687116](https://go.dev/cl/687116) removed the proposed legacy escape option. Maintainers cited existing v1 inconsistencies, RFC 8785 formatting, and precedent for changing semantically equivalent JSON bytes. | Intended v1-facade byte-output change, not considered a Go 1 compatibility violation. Upstream said it may reconsider a switch only if later releases show substantial breakage. |
| `TMPTR-UNJ` | [#75361](https://github.com/golang/go/issues/75361) is still open. It covers the same named-pointer method-discovery mechanism and a worse recursive form. The working group chose to keep v2's consistent pointer-receiver dispatch; it recommends declaring a method-free value type and converting its pointer. [CL 738361](https://go.dev/cl/738361) explored rejecting recursive named-pointer calls, but the runtime guard was abandoned after a measured 3–9% cost; the planned vet check was not ready for Go 1.27. | Method dispatch is an intended v2 semantic that upstream decided not to emulate exactly in the v1 facade. Recursion detection remains unresolved; the exact non-recursive Sonic case is mechanism-aligned but not the issue's reproducer. |
| `UNJ-PARTIAL` | [#74614](https://github.com/golang/go/issues/74614) fixed a related facade regression where a custom unmarshaler ran before invalid trailing input was discovered; [CL 689919](https://go.dev/cl/689919) restored whole-value validation before mutation. | Upstream treats consequential error-time mutation as fixable when it can reproduce v1 behavior. That issue does not cover continuing to later fields after `UnmarshalJSON` itself returns an error. No exact issue was found for Sonic's case, so it remains unresolved. |
| `TM-KEY-QUOTE` | [`CallMethodsWithLegacySemantics`](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2_options.go#L233-L268) explicitly preserves special v1 method rules for map keys. Historical [#29732](https://github.com/golang/go/issues/29732) documents map-key method asymmetry but not already-quoted `MarshalText` output. | No exact official issue or CL was found. Since the active v1 compatibility option still fails to reproduce the legacy bytes, classify this as an unresolved v1-facade difference rather than an intended direct-v2 default. |
| `IFACE-SKIP-NONPTR`, `IFACE-SKIP-NILPTR` | `DefaultOptionsV1` includes `MergeWithLegacySemantics`, whose [documentation explicitly covers interface values](https://github.com/golang/go/blob/release-branch.go1.27/src/encoding/json/v2_options.go#L349-L376). Historical [#20994](https://github.com/golang/go/issues/20994) and [#21494](https://github.com/golang/go/issues/21494) show that interface replacement has long been compatibility-sensitive, but neither is this rc2 failure case. | Clearing the interface despite the v1 merge option is a v1-facade regression candidate. No exact upstream report or resolution was found for the non-pointer or typed-nil cases. |
| `SELF-REF` | The same interface merge option is the nearest official contract. No exact golang/go issue was found for `v = &v` causing a stack overflow under the Go 1.27 facade. [#75361](https://github.com/golang/go/issues/75361) is related only in showing that recursive method discovery must be bounded. | Unresolved and high priority to report upstream. A process-ending stack overflow is not an intentional v2 semantic and must not be copied into Sonic. |

Two related upstream decisions show the boundary more clearly:

- [#79659](https://github.com/golang/go/issues/79659) and
  [CL 785420](https://go.dev/cl/785420) fixed incorrect v1
  `SyntaxError`/`UnmarshalTypeError` offsets.
- [#79770](https://github.com/golang/go/issues/79770) accepted some type-name
  and containing-struct changes while fixing the offset error. This is
  consistent with the release-note warning that exact error presentation may
  change, while machine-consumed position data still needs compatibility.

### Official Mapping for Direct JSONV2 Defaults

These six cases are not facade regressions. They are documented v2 defaults
with explicit migration controls in the accepted proposal and
[`v2_options.go`](https://go.dev/src/encoding/json/v2_options.go):

| Direct JSONV2 default | v1 behavior restored by |
|---|---|
| Reject duplicate object names | `jsontext.AllowDuplicateNames(true)` |
| Match struct fields case-sensitively | `jsonv2.MatchCaseInsensitiveNames(true)` plus the v1 delimiter rule |
| Encode nil slices/maps as `[]`/`{}` | `jsonv2.FormatNilSliceAsNull(true)` and `FormatNilMapAsNull(true)` |
| Require an exact array length | `jsonv1.UnmarshalArrayFromAnyLength(true)` |
| Encode `[N]byte` as Base64 | `jsonv1.FormatByteArrayAsArray(true)` |
| Reject invalid UTF-8 | `jsontext.AllowInvalidUTF8(true)` |

Historical issues such as
[#14750](https://github.com/golang/go/issues/14750) for case-insensitive field
matching and [#37711](https://github.com/golang/go/issues/37711) /
[#27589](https://github.com/golang/go/issues/27589) for nil collection output
explain why configuration was requested. The final accepted resolution is
#71497: strict v2 defaults, a complete v1 preset, and composable options for
hybrid migration.

## Alignment Sentinels

The catalog also keeps cases that currently agree with the active v1 facade.
They catch accidental Sonic drift:

- Invalid UTF-8 decoding replaces invalid bytes with U+FFFD.
- Plain `TextMarshaler` map keys remain aligned.
- Duplicate object names are accepted and the last value wins.
- Oversized arrays are truncated when decoding into a fixed-size Go array.
- A `[3]byte` value encodes as `[1,2,3]`.
- A non-nil `*bool(false)` is not omitted by `omitempty`.
- Trailing non-whitespace input returns an error.

Separate preset cases document intentional Sonic configuration differences:

- `ConfigDefault` does not HTML-escape; `ConfigStd` does.
- A nil slice encodes as `null` by default.
- Field-name matching is case-insensitive unless `CaseSensitive` is enabled.

## Direct JSONV2 Defaults

`TestJSONV2Semantics` records intentional differences between the v1 facade
and direct `encoding/json/v2` use:

| Case and test | SonicStd / GoStd127 | JSONV2 |
|---|---|---|
| [Duplicate names](../compatibility/jsonv2_test.go#L34) | Accepted; last value wins | Error |
| [Case folding](../compatibility/jsonv2_test.go#L45) | `foo` matches `Foo` | Case-sensitive by default |
| [Nil slice](../compatibility/jsonv2_test.go#L58) | `null` | `[]` |
| [Oversized array](../compatibility/jsonv2_test.go#L72) | Accepted and truncated | Error |
| [Byte array](../compatibility/jsonv2_test.go#L83) | `[1,2,3]` | `"AQID"` |
| [Invalid UTF-8](../compatibility/jsonv2_test.go#L97) | Replaced with U+FFFD | Error |

These tests describe migration impact. They do not make JSONV2 defaults a
`ConfigStd` compatibility requirement.

## Sonic Response Options

The runtime ABI and JSON behavior require different strategies. Build tags are
appropriate for private runtime layouts, which are selected by the toolchain.
Public JSON semantics need a stable contract that does not change silently
when users rebuild the same source with a newer Go release.

### Option 1: Preserve the Existing API Semantics

Keep `ConfigStd` aligned with classic `encoding/json` v1 semantics across Go
versions. Record GoStd127 drift in compatibility tests, but do not reproduce
release-candidate implementation quirks.

Advantages:

- Existing applications retain the same accepted types, output, method
  dispatch, and error-time mutation behavior after a toolchain upgrade.
- One semantic contract applies to Go 1.18 through Go 1.27.
- Upstream regressions such as self-reference stack overflow are not copied
  into Sonic.

Costs:

- `ConfigStd` is not byte-for-byte or state-for-state identical to the Go 1.27
  implementation in every edge case.
- The compatibility catalog and issue snapshots must remain part of CI.
- Users comparing Sonic directly with the active standard-library backend must
  classify intentional differences instead of assuming equality.

This is the current behavior and the recommended default.

### Option 2: Select Go 1.27 Semantics with Build Tags

Add Go-version or experiment-specific files, for example
`go1.27 && goexperiment.jsonv2`, and make the existing API use GoStd127
behavior when built with that toolchain.

Advantages:

- The default Sonic behavior can track the active `encoding/json`
  implementation without a user-visible configuration change.
- Differential testing against the standard library becomes simpler.

Costs and risks:

- Rebuilding unchanged application source with Go 1.27 changes its JSON
  contract. There is no runtime opt-out inside the same binary.
- Each semantic path multiplies JIT, OptDec, VM, architecture, and experiment
  combinations in CI.
- Some observed differences are implementation artifacts, not desirable
  semantics. Sonic must not deliberately reproduce stack overflow or unstable
  partial-write behavior.
- Experiment build tags are suitable for test selection but are a fragile
  public product contract.

Build tags should continue to select runtime ABI implementations such as
`loader/funcdata_go127.go`. They are not recommended for silently changing the
semantics of `ConfigStd`.

### Option 3: Add Explicit Encoder and Decoder Options

Keep `ConfigStd` stable and expose selected v2-style behavior through an
opt-in preset and granular options. Existing options already cover part of the
space:

- `NoNullSliceOrMap` selects empty arrays and objects instead of `null`;
- `CaseSensitive` selects strict struct-field name matching;
- `ValidateString` provides stricter string validation, although it is not a
  complete JSONV2 compatibility switch.

Additional behavior would need explicit support, potentially including:

- preferring `encoding.TextAppender` over `encoding.TextMarshaler`;
- rejecting duplicate object names;
- requiring exact array lengths;
- encoding byte arrays as Base64;
- accepting additional map-key types;
- defining named-pointer marshaler and unmarshaler dispatch;
- rejecting invalid UTF-8 with a stable error contract.

The API should prefer a named preset such as `ConfigJSONV2` plus granular
encoder and decoder flags. A single `Go127Semantics` boolean would couple the
contract to a toolchain version and make later evolution ambiguous.

Not every observed difference should become an option. Error-time destination
mutation, interface clearing, and cycle handling should have one documented,
safe behavior. In particular, self-reference must be bounded by cycle or depth
detection rather than emulating a stack overflow.

### Decision Matrix

| Criterion | Preserve v1 semantics | Build-tag default | Explicit options |
|---|---|---|---|
| Existing user compatibility | Best | Poor | Best |
| Same-toolchain parity with GoStd127 | Partial | Best | Opt-in |
| Runtime selection | Not needed | Compile time only | Per API instance |
| Implementation and CI cost | Lowest | Highest | Medium to high |
| Risk of silent behavior change | Lowest | Highest | Low |
| Suitable for runtime ABI adaptation | No | Yes | No |
| Suitable for JSON behavior | Default contract | Not recommended | Recommended opt-in |

### Recommended Direction

1. Keep `ConfigStd` on classic v1 semantics.
2. Continue using build tags only for runtime ABI and experiment-specific test
   selection.
3. If user demand exists, add an explicit JSONV2-oriented preset and implement
   stable granular options incrementally.
4. Do not expose RC-only bugs or accidental partial-write behavior as supported
   configuration.
5. Treat any future `ConfigJSONV2` contract as opt-in and test it against both
   direct `encoding/json/v2` and the Go 1.27 v1 facade.
6. Fix backend-independent `ConfigStd` gaps such as syntax prevalidation,
   map-value replacement, quoted numeric parsing, and marshaler error wrapping
   independently of any JSONV2 API design.

## Running the Tests

`compatibility/` is a separate workspace module. A root-level `go test ./...`
does not run it.

```sh
# Go 1.27 default encoding/json backend
env -u GOROOT GOTOOLCHAIN=go1.27rc2 \
  go test ./compatibility/ -count=1

# Go 1.27 legacy encoding/json backend
env -u GOROOT GOTOOLCHAIN=go1.27rc2 GOEXPERIMENT=nojsonv2 \
  go test ./compatibility/ -count=1

# Sonic optimized paths
env -u GOROOT GOTOOLCHAIN=go1.27rc2 \
  SONIC_USE_OPTDEC=1 SONIC_USE_FASTMAP=1 SONIC_ENCODER_USE_VM=1 \
  go test ./compatibility/ -count=1

# Sonic optimized paths with the legacy encoding/json backend
env -u GOROOT GOTOOLCHAIN=go1.27rc2 GOEXPERIMENT=nojsonv2 \
  SONIC_USE_OPTDEC=1 SONIC_USE_FASTMAP=1 SONIC_ENCODER_USE_VM=1 \
  go test ./compatibility/ -count=1
```

The broader Go 1.27 validation set is:

```sh
env -u GOROOT GOTOOLCHAIN=go1.27rc2 GOMAXPROCS=4 \
  go test ./... -count=1

env -u GOROOT GOTOOLCHAIN=go1.27rc2 GOMAXPROCS=4 \
  SONIC_USE_OPTDEC=1 SONIC_USE_FASTMAP=1 SONIC_ENCODER_USE_VM=1 \
  go test ./... -count=1

env -u GOROOT GOTOOLCHAIN=go1.27rc2 GOMAXPROCS=4 \
  go test -race ./loader/... -count=1

env -u GOROOT GOTOOLCHAIN=go1.27rc2 GOMAXPROCS=4 \
  go test -race ./issue_test -count=1

env -u GOROOT GOTOOLCHAIN=go1.27rc2 GOEXPERIMENT=nojsonv2 GOMAXPROCS=4 \
  go test -race ./issue_test -count=1
```

## CI Coverage and Limits

The Go 1.27 x86 job explicitly runs:

- issue regressions with `GOEXPERIMENT=nojsonv2`;
- compatibility tests with the default backend;
- compatibility tests with the legacy backend;
- compatibility tests with OptDec, FastMap, and encoder VM enabled under both
  backends.

The repository matrix also covers Linux amd64, Linux arm64, macOS arm64, and
Windows amd64. These jobs prove the tested source state against rc2. They do
not prove that final runtime internals will remain unchanged, and they do not
replace validation with application workloads.

## Adding a Compatibility Case

A new case should state:

1. Whether it compares Sonic with GoStd127 or the v1 facade with JSONV2.
2. The result under both the default and legacy backends.
3. Whether JIT, OptDec, FastMap, or encoder VM changes the result.
4. The exact output, error presence, and mutated destination state.
5. Whether calling one implementation is unsafe, such as the self-reference
   stack-overflow case.

For interfaces and pointers, assert dynamic types, pointer layers, typed nils,
and partial writes directly. Re-marshaling the value is not sufficient because
different internal states can produce identical JSON.

## Go 1.27 Final Acceptance

After the final toolchain is published:

1. Replace or supplement `1.27.0-rc.2` in CI with the final version.
2. Compare `moduledata` and the map iterator linkname shim with final sources.
3. Rerun the main, optimized, loader race, issue race, and compatibility
   suites.
4. Confirm Linux amd64, Linux arm64, macOS arm64, and Windows amd64 jobs.
5. Reduce any remaining GoStd127 v1-facade differences to upstream reports
   before changing Sonic semantics.

If final changes only error wording, tests should continue to assert error
types or presence rather than complete strings. Changes to output, dynamic
types, or destination state require a new compatibility decision.

## Primary Sources

- [Go 1.27 release notes](https://go.dev/doc/go1.27)
- [Go 1.27 runtime `moduledata`](https://github.com/golang/go/blob/release-branch.go1.27/src/runtime/symtab.go)
- [Go 1.27 legacy map linkname shim](https://github.com/golang/go/blob/release-branch.go1.27/src/runtime/linkname_shim.go)
- [`encoding/json/v2` package documentation](https://pkg.go.dev/encoding/json/v2)
- [A new experimental Go API for JSON](https://go.dev/blog/jsonv2-exp)
