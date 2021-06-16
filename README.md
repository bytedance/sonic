# Sonic

A blazingly fast JSON serializing &amp; deserializing library, accelerated by JIT(just-in-time compiling) and SIMD(single-instruction-multi-data).

**WARNING: This is still in alpha stage, use with care !**

## Benchmarks
For all sizes of json and all scenes of usage, Sonic performs almost best.
- Small (400B, 11 keys, 3 levels)
![small benchmarks](bench-400B.png)
- Medium (110KB, 300+ keys, 3 levels, with many quoted-json values)
![medium benchmarks](bench-110KB.png)
- Large (550KB, 10000+ key, 6 levels)
![large benchmarks](bench-550KB.png)

For a 13KB [TwitterJson](https://github.com/bytedance/sonic/blob/main/decoder/testdata_test.go#L19)(cpu i9-9880H, goarch amd64), Sonic is **1.5x** faster than [json-iterator](https://github.com/json-iterator/go) in decoding, **2.5x** faster in encoding.

```powershell
BenchmarkDecoder_Generic_Sonic-16                          10000             54309 ns/op         240.01 MB/s       46149 B/op        303 allocs/op
BenchmarkDecoder_Generic_StdLib-16                         10000            135268 ns/op          96.36 MB/s       50899 B/op        772 allocs/op
BenchmarkDecoder_Generic_JsonIter-16                       10000             96701 ns/op         134.80 MB/s       55791 B/op       1068 allocs/op
BenchmarkDecoder_Binding_Sonic-16                          10000             29478 ns/op         442.20 MB/s       26062 B/op         34 allocs/op
BenchmarkDecoder_Binding_StdLib-16                         10000            119348 ns/op         109.22 MB/s       10560 B/op        207 allocs/op
BenchmarkDecoder_Binding_JsonIter-16                       10000             37646 ns/op         346.25 MB/s       14673 B/op        385 allocs/op
BenchmarkEncoder_Generic_Sonic-16                          10000             25894 ns/op         503.39 MB/s       19096 B/op         42 allocs/op
BenchmarkEncoder_Generic_JsonIter-16                       10000             50275 ns/op         259.27 MB/s       13432 B/op         77 allocs/op
BenchmarkEncoder_Generic_StdLib-16                         10000            154901 ns/op          84.15 MB/s       48173 B/op        827 allocs/op
BenchmarkEncoder_Binding_Sonic-16                          10000              7373 ns/op        1768.04 MB/s       13861 B/op          4 allocs/op
BenchmarkEncoder_Binding_JsonIter-16                       10000             23223 ns/op         561.31 MB/s        9489 B/op          2 allocs/op
BenchmarkEncoder_Binding_StdLib-16                         10000             19512 ns/op         668.07 MB/s        9477 B/op          1 allocs/op
```
More detail see [ast/search_test.go](https://github.com/bytedance/sonic/blob/main/ast/search_test.go), [decoder/decoder_test.go](https://github.com/bytedance/sonic/blob/main/decoder/decoder_test.go), [encoder/encoder_test.go](https://github.com/bytedance/sonic/blob/main/encoder/encoder_test.go),

## Usage

### Marshal/Unmarshal

The behaviors are mostly consistent with encoding/json, except some uncommon escaping and key sorting (see [issue4](https://github.com/bytedance/sonic/issues/4))
 ```go
import "github.com/bytedance/sonic"

// Marshal
output, err := sonic.Marshal(&data) 
// Unmarshal
err := sonic.Unmarshal(input, &data) 
 ```

### Get

Search partial json by given pathes, which must be non-negative integer or string or nil
```go
import "github.com/bytedance/sonic"

input := []byte(`{"key1":[{},{"key2":{"key3":[1,2,3]}}]}`)

// no path, returns entire json
root, err := sonic.Get(input)
raw := root.Raw() // == string(input)

// multiple pathes
root, err := sonic.Get(input, "key1", 1, "key2")
sub := root.Get("key3").Index(2).Int64() // == 3
```
Returned ast.Node supports：
- secondary search: `Get()`, `Index()`, `GetByPath()`
- type assignment: `Int64()`, `Float64()`, `String()`, `Number()`, `Bool()`, `Map()`, `Array()`
- children traversal: `Values()`, `Properties()`
- supplement: `Set()`, `SetByIndex()`, `Add()`, `Cap()`, `Len()`

### Use Number/Use Int64
 ```go
import "github.com/bytedance/sonic/decoder"

input := `1`
var data interface{}

// default float64
dc := decoder.NewDecoder(input) 
dc.Decode(&data) // data == float64(1)
// use json.Number
dc = decoder.NewDecoder(input)
dc.UseNumber()
dc.Decode(&data) // data == json.Number("1")
// use int64
dc = decoder.NewDecoder(input)
dc.UseInt64()
dc.Decode(&data) // data == int64(1)

root, err := sonic.GetFromString(input)
// Get json.Number
jn := root.Number()
jm := root.InterfaceUseNumber().(json.Number) // jn == jm
// Get float64
fn := root.Float64()
fm := root.Interface().(float64) // jn == jm
 ```

## Tips

### Pretouch
Since Sonic uses JIT(just-in-time) compiling for decoder/encoder, huge schema may cause request-timeout. For better stability, we suggest to use `Pretouch()` for more-than-10000-field schema(struct) before `Marshal()/Unmarshal()`.
```go
import (
    "reflect"
    "github.com/bytedance/sonic"
)

func init() {
    var v HugeStruct
    err := sonic.Pretouch(reflect.TypeOf(v))
}
```

### Pass string or []byte?
For alignment to encoding/json, we provide API to pass `[]byte` as arguement, but the string-to-bytes copy is conducted at the same time considering safety, which may lose performance when origin json is huge. Therefore, you can use `UnmarshalString`, `GetFromString` to pass string, as long as your origin data is string or **nocopy-cast** is safe for your []byte.

### Avoid repeating work
`Get()` overlapping pathes from the same root may cause repeating parsing. Instead of using `Get()` several times, you can use parser and searcher together like this:
```go
import "github.com/bytedance/sonic"

root, err := sonic.GetByString(_TwitterJson, "statuses", 3, "user")
a = root.GetByPath( "entities","description")
b = root.GetByPath( "entities","url")
c = root.GetByPath( "created_at")
```
No need to worry about the overlaping or overparsing of a, b and c, because the inner parser of their root is lazy-loaded.
### Better performance for generic deserializing
In most cases of fully-load generic json, `Unmarshal()` performs better than `ast.Loads()`. But if you only want to search a partial json and convert it into `interface{}` (or `map[string]interface{}`, `[]interface{}`), we advise you to combine `Get()` and `Unmarshal()`:
```go
import "github.com/bytedance/sonic"

node, err := sonic.GetByString(_TwitterJson, "statuses", 3, "user")
var user interface{}
err = sonic.UnmarshalString(node.Raw(), &user)
```
