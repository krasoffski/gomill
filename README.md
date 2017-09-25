# Golang experimental modules and utilities 

[![GoDoc](https://godoc.org/github.com/krasoffski/gomill?status.svg)](https://godoc.org/github.com/krasoffski/gomill)
[![Go Report Card](https://goreportcard.com/badge/github.com/krasoffski/gomill)](https://goreportcard.com/report/github.com/krasoffski/gomill)

## Notes about useful tools

### To lint code and tools
 * [binstale] - verifies the binaries in your GOPATH/bin are stale or up to date
 * [go-torch] - flame graph profiler for Go programs
 * [errcheck] - checks for unchecked errors in Go programs
 * [interfacer] - linter that suggests interface types

### Checking coverage
```sh
$ go test -coverprofile cover.report
$ go tool cover -html=cover.report -o cover.html
```

### Inspecting package
```sh
$ go list -f '{{ .Name }}: {{ .Doc }}'
unique: Package unique provides a simple function for removing...
```
```sh
$ go list -f '{{ .Imports }}'
[flag fmt sync]
```
```sh
$ go list -f '{{ .Imports }}' fmt
[errors io math os reflect strconv sync unicode/utf8]
```
```sh
$ go list -f '{{ join .Imports "\n" }}' fmt
errors
io
math
os
reflect
strconv
sync
unicode/utf8
```

### Getting documentation
```sh
$ go doc 'github.com/krasoffski/gomill/unique' Strings
func Strings(s []string) []string
    Strings removes duplicated strings from a slice of strings. It returns a new
    slice of strings without duplicates.
```

### Performing CPU profiling
Enable profiling in your code.

```go
package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("multiplier.cpuprofile")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Your program here
}
```
Than you can analyze report.
```sh
$ go tool pprof multiplier multiplier.cpuprofile
Entering interactive mode (type "help" for commands)
(pprof) top
3150ms of 5640ms total (55.85%)
Dropped 18 nodes (cum <= 28.20ms)
Showing top 10 nodes out of 95 (cum >= 3180ms)
      flat  flat%   sum%        cum   cum%
    1700ms 30.14% 30.14%     1920ms 34.04%  syscall.Syscall
     440ms  7.80% 37.94%     1290ms 22.87%  runtime.selectgoImpl
     250ms  4.43% 42.38%      250ms  4.43%  runtime/internal/atomic.Xchg
     150ms  2.66% 45.04%      150ms  2.66%  runtime.futex
     130ms  2.30% 47.34%      130ms  2.30%  runtime/internal/atomic.Cas
     120ms  2.13% 49.47%      120ms  2.13%  runtime.usleep
     110ms  1.95% 51.42%      160ms  2.84%  fmt.(*fmt).integer
     100ms  1.77% 53.19%     2670ms 47.34%  fmt.Fprintln
      80ms  1.42% 54.61%       80ms  1.42%  runtime/internal/atomic.Load
      70ms  1.24% 55.85%     3180ms 56.38%  main.main
(pprof) web
(pprof) list multiplier
```
What is more, you are able to create a [Flame Graph].
```sh
$ go-torch --file=output.svg multiplier multiplier.cpuprofile
INFO[14:49:02] Run pprof command: go tool pprof -raw -seconds 30 multiplier multiplier.cpuprofile
INFO[14:49:02] Writing svg to output.svg
$ open torch.svg
```
_Note: without application load CPU profile might be empty._

### Tracing with `go tool trace`
Enable tracing in your code.

```go
package main

import (
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// Your program here
}
```
Than you can open web brouser for investigation
```sh
$ go tool trace trace.out
```

### Setting version
Version defined in the source code of data [race] example.
```go
// Version of race example.
var Version = "0.1.0"
```
Inject new version using `-ldflags`.
```sh
$ go build -ldflags="-X main.Version=0.2.1"
$ ./race -version
Version: 0.2.1
```

### Small binary file
Remove the debugging information included in the executable
binary file using `-ldflags`. Also re-pack binary using `upx`.

```sh
$ go build
$ du -h crawler
5.4M	crawler
```
```sh
$ go build -ldflags="-s -w"
$ du -h crawler
3.6M	crawler
```
```sh
$ upx crawler
$ du -h crawler
1.4M	crawler
```

### Running simple benchmark test
Create simple benchmark test like bellow
```go
package popcount

import "testing"

func BenchmarkPopcount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(275032803564053945)
	}
}
```
Run test with options like `go test -bench=.` or `go test -bench=BenchmarkPopcount1`:
```sh
$ go test -bench=.
PASS
BenchmarkPopcount1-2	200000000	         9.36 ns/op
ok  	github.com/krasoffski/gomill/gopl/ch02/popcount	2.824s
```


[binstale]: https://github.com/shurcooL/binstale
[go-torch]: https://github.com/uber/go-torch
[errcheck]: https://github.com/kisielk/errcheck
[interfacer]: https://github.com/mvdan/interfacer/
[Flame Graph]: https://github.com/brendangregg/FlameGraph

[race]: https://godoc.org/github.com/krasoffski/gomill/race
