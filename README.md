# Sonic

A blazingly fast JSON serializing &amp; deserializing library, accelerated by JIT (just-in-time compiling) and SIMD (single-instruction-multiple-data).

## Requirement
- Go 1.15/1.16/1.17
- Linux/darwin OS
- Amd64 CPU with AVX instruction set

## Features
- Runtime object binding without code generation
- Complete APIs for JSON value manipulation
- Fast, fast, fast!

## Benchmarks
For **all sizes** of json and **all scenarios** of usage, **Sonic performs best**.
- [Medium](https://github.com/bytedance/sonic/blob/main/decoder/testdata_test.go#L19) (13KB, 300+ key, 6 layers)
```powershell
goversion: 1.17.1
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkEncoder_Generic_Sonic                        25181 ns/op         517.65 MB/s       13035 B/op          4 allocs/op
BenchmarkEncoder_Generic_JsonIter                     43765 ns/op         297.84 MB/s       13433 B/op         77 allocs/op
BenchmarkEncoder_Generic_StdLib                      108776 ns/op         119.83 MB/s       49137 B/op        827 allocs/op
BenchmarkEncoder_Binding_Sonic                         6282 ns/op        2075.01 MB/s       13765 B/op          4 allocs/op
BenchmarkEncoder_Binding_JsonIter                     20740 ns/op         628.51 MB/s        9487 B/op          2 allocs/op
BenchmarkEncoder_Binding_StdLib                       16661 ns/op         782.34 MB/s        9479 B/op          1 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic-16             4072 ns/op        3200.89 MB/s       11052 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Generic_JsonIter-16         11379 ns/op        1145.52 MB/s       13458 B/op         77 allocs/op
BenchmarkEncoder_Parallel_Generic_StdLib-16           50635 ns/op         257.43 MB/s       49183 B/op        827 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic-16             1304 ns/op        9994.64 MB/s       10925 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_JsonIter-16          6072 ns/op        2146.76 MB/s        9505 B/op          2 allocs/op
BenchmarkEncoder_Parallel_Binding_StdLib-16            3510 ns/op        3713.89 MB/s        9481 B/op          1 allocs/op

BenchmarkDecoder_Generic_Sonic                        53843 ns/op         242.09 MB/s       49779 B/op        313 allocs/op
BenchmarkDecoder_Generic_StdLib                      130402 ns/op          99.96 MB/s       50868 B/op        772 allocs/op
BenchmarkDecoder_Generic_JsonIter                     92810 ns/op         140.45 MB/s       55788 B/op       1068 allocs/op
BenchmarkDecoder_Binding_Sonic                        29793 ns/op         437.52 MB/s       24778 B/op         34 allocs/op
BenchmarkDecoder_Binding_StdLib                      121206 ns/op         107.54 MB/s       10576 B/op        208 allocs/op
BenchmarkDecoder_Binding_JsonIter                     36099 ns/op         361.09 MB/s       14674 B/op        385 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic-16            10319 ns/op        1263.21 MB/s       49423 B/op        313 allocs/op
BenchmarkDecoder_Parallel_Generic_StdLib-16           58526 ns/op         222.72 MB/s       50875 B/op        772 allocs/op
BenchmarkDecoder_Parallel_Generic_JsonIter-16         60156 ns/op         216.69 MB/s       55812 B/op       1068 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic-16             7265 ns/op        1794.18 MB/s       24952 B/op         34 allocs/op
BenchmarkDecoder_Parallel_Binding_StdLib-16           44000 ns/op         296.25 MB/s       10575 B/op        208 allocs/op
BenchmarkDecoder_Parallel_Binding_JsonIter-16         21029 ns/op         619.86 MB/s       14678 B/op        385 allocs/op

BenchmarkGetOne_Sonic                                 17070 ns/op         762.94 MB/s          29 B/op          1 allocs/op
BenchmarkGetOne_Gjson                                 19714 ns/op         660.59 MB/s           0 B/op          0 allocs/op
BenchmarkGetOne_Jsoniter                              99281 ns/op         131.17 MB/s       27936 B/op        647 allocs/op
BenchmarkSetOne_Sonic                                 23730 ns/op         548.80 MB/s        1883 B/op         17 allocs/op
BenchmarkSetOne_Sjson                                 57680 ns/op         225.78 MB/s       52180 B/op          9 allocs/op
BenchmarkSetOne_Jsoniter                             104018 ns/op         125.20 MB/s       45859 B/op        964 allocs/op
BenchmarkGetOne_Parallel_Sonic-16                      2010 ns/op        6479.41 MB/s         114 B/op          1 allocs/op
BenchmarkGetOne_Parallel_Gjson-16                      1815 ns/op        7176.39 MB/s           0 B/op          0 allocs/op
BenchmarkGetOne_Parallel_Jsoniter-16                  23261 ns/op         559.86 MB/s       27942 B/op        647 allocs/op
BenchmarkSetOne_Parallel_Sonic-16                      2007 ns/op        6487.78 MB/s        2202 B/op         17 allocs/op
BenchmarkSetOne_Parallel_Sjson-16                     12422 ns/op        1048.40 MB/s       52180 B/op          9 allocs/op
BenchmarkSetOne_Parallel_Jsoniter-16                  39204 ns/op         332.18 MB/s       45889 B/op        964 allocs/op
```        
- [Small](https://github.com/bytedance/sonic/blob/main/testdata/small.go) (400B, 11 keys, 3 layers)
![small benchmarks](bench-small.jpg)
- [Large](https://github.com/bytedance/sonic/blob/main/testdata/twitter.json) (635KB, 10000+ key, 6 layers)
![large benchmarks](bench-large.jpg)

See [bench.sh](https://github.com/bytedance/sonic/blob/main/bench.sh) for benchmark codes.

## How it works
See [INTRODUCTION.md](INTRODUCTION.md).

## Usage

### Marshal/Unmarshal

Their behaviors are mostly consistent with `encoding/json`, except two escaping form (see [issue4](https://github.com/bytedance/sonic/issues/4)) that is **NOT** in conformity to [RFC8259](https://datatracker.ietf.org/doc/html/rfc8259).
 ```go
import "github.com/bytedance/sonic"

var data YourSchema
// Marshal
output, err := sonic.Marshal(&data) 
// Unmarshal
err := sonic.Unmarshal(output, &data) 
 ```

### Use Number/Use Int64
 ```go
import "github.com/bytedance/sonic/decoder"

var input = `1`
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

### Sort Keys
On account of the performance loss from sorting (roughly 10%), sonic doesn't enable this feature by default. If your component depends on it to work (like [zstd](https://github.com/facebook/zstd)), Use it like this:
```go
import "github.com/bytedance/sonic/encoder"

m := map[string]interface{}{}
v, err := encoder.Encode(m, encoder.SortMapKeys)
```

### Print Syntax Error
```go
import "github.com/bytedance/sonic"
import "github.com/bytedance/sonic/decoder"

var data interface{}
err := sonic.Unmarshal("[[[}]]", &data)
if err != nil {
    /*one line by default*/
    println(e.Error())) // "Syntax error at index 3: invalid char\n\n\t[[[}]]\n\t...^..\n"
    /*pretty print*/
    if e, ok := err.(decoder.SyntaxError); ok {
        /*Syntax error at index 3: invalid char

            [[[}]]
            ...^..
        */
        print(e.Description())
    }
}
```

### Ast.Node
Sonic/ast.Node is a completely self-contained AST for JSON. It implements serialization and deserialization both, and provides robust APIs for obtaining and modification of generic data.
#### Get/Index
Search partial JSON by given paths, which must be non-negative integer or string or nil
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
**Tip**: since `Index()` uses offset to locate data, which is faster much than scanning like `Get()`, we suggest you use it as much as possible. And sonic also provides another API `IndexOrGet()` to underlying use offset as well as ensuring the key is matched.

#### Set/Unset
Modify the json content by Set()/Unset()
```go
import "github.com/bytedance/sonic"

// Set
exist, err := root.Set("key4", NewBool(true)) // exist == false
alias1 := root.Get("key4") 
println(alias1.Valid()) // true
alias2 := root.Index(1)
println(alias1 == alias2) // true

// Unset
exist, err := root.UnsetByIndex(1) // exist == true
println(root.Get("key4").Check()) // "value not exist"
```

#### Serialize
To encode `ast.Node` as json, use `MarshalJson()` or `json.Marshal()` (MUST pass the node's pointer)
```go
import (
    "encoding/json"
    "github.com/bytedance/sonic"
)

buf, err := root.MarshalJson()
println(string(buf))                // {"key1":[{},{"key2":{"key3":[1,2,3]}}]}
exp, err := json.Marshal(&root)     // WARN: use pointer
println(string(buf) == string(exp)) // true
```

#### APIs
- validation: `Check()`, `Error()`, `Valid()`, `Exist()`
- searching: `Index()`, `Get()`, `IndexPair()`, `IndexOrGet()`, `GetByPath()`
- go-type casting: `Int64()`, `Float64()`, `String()`, `Number()`, `Bool()`, `Map[UseNumber|UseNode]()`, `Array[UseNumber|UseNode]()`, `Interface[UseNumber|UseNode]()`
- go-type packing: `NewRaw()`, `NewNumber()`, `NewNull()`, `NewBool()`, `NewString()`, `NewObject()`, `NewArray()`
- iteration: `Values()`, `Properties()`
- modification: `Set()`, `SetByIndex()`, `Add()`, `Cap()`, `Len()`

## Tips

### Pretouch
Since Sonic uses [golang-asm](https://github.com/twitchyliquid64/golang-asm) as a JIT assembler, which is NOT very suitable for runtime compiling, first-hit running of a huge schema may cause request-timeout or even process-OOM. For better stability, we advise to **use `Pretouch()` for huge-schema or compact-memory application** before `Marshal()/Unmarshal()`.
```go
import (
    "reflect"
    "github.com/bytedance/sonic"
    "github.com/bytedance/sonic/option"
 )
 
 func init() {
     var v HugeStruct
    // For most large types (nesting depth <= 5)
     err := sonic.Pretouch(reflect.TypeOf(v))
    // If the type is too deep nesting (nesting depth > 5),
    // you can set compile recursive depth in Pretouch for better stability in JIT.
    err := sonic.Pretouch(reflect.TypeOf(v), option.WithCompileRecursiveDepth(depth))
```

### Pass string or []byte?
For alignment to `encoding/json`, we provide API to pass `[]byte` as an argument, but the string-to-bytes copy is conducted at the same time considering safety, which may lose performance when origin JSON is huge. Therefore, you can use `UnmarshalString` and `GetFromString` to pass a string, as long as your origin data is a string or **nocopy-cast** is safe for your []byte.

### Better performance for generic data
In **fully-parsed** scenario, `Unmarshal()` performs better than `Get()`+`Node.Interface()`. But if you only have a part of schema for specific json, you can combine `Get()` and `Unmarshal()` together:
```go
import "github.com/bytedance/sonic"

node, err := sonic.GetFromString(_TwitterJson, "statuses", 3, "user")
var user User // your partial schema...
err = sonic.UnmarshalString(node.Raw(), &user)
```
Even if you don't have any schema, use `ast.Node` as the container of generic values instead of `map` or `interface`:
```go
import "github.com/bytedance/sonic"

root, err := sonic.GetFromString(_TwitterJson)
user := root.GetByPath("statuses", 3, "user")  // === root.Get("status").Index(3).Get("user")
err = user.Check()

// err = user.LoadAll() // only call this when you want to use 'user' concurrently...
go someFunc(user)
```
Why? Because `ast.Node` stores its children using `array`: 
- `Map`'s performance degrades a lot once rehashing triggered, but `ast.Node` doesn't have this concern;
- **Hashing** (`map[x]`) is not as efficient as **Indexing** (`array[x]`), which `ast.Node` can conduct on **both array and object**.
- Using `Interface()`/`Map()` means Sonic must parse all the underlying values, while in most cases you don't need them all;

**CAUTION:** `ast.Node` **DOESN'T** ensure concurrent security directly, due to its **lazy-load** design. However, your can call `Node.Load()`/`Node.LoadAll()` to achieve that, which may bring performance reduction while it still works faster than converting to `map` or `interface{}` 