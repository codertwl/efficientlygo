# estring
<br>
## concats
```
goos: linux
goarch: amd64
Benchmark_Native-16               500000            337071 ns/op
Benchmark_Native2-16              500000            338291 ns/op
Benchmark_Native3-16            30000000                95.3 ns/op
Benchmark_Join-16               10000000               196 ns/op
Benchmark_BytesBuffer-16        30000000                46.1 ns/op
Benchmark_StringsBuilder-16     30000000                47.2 ns/op
Benchmark_Sprint-16             10000000               135 ns/op
PASS
ok      command-line-arguments  347.324s
```

