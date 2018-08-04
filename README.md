bench cpu
```bash
go test -v -run Testxxx
go test -v -bench BenchmarkQuickSort -run NONE
```

bench memory
```bash
$ go test -v -bench BenchmarkMalloc -benchmem -gcflags "-N -l"
```
