# Miscellaneous Golang experiment library.

## Notes about useful tools

### To lint code and tools
 * [binstale] - verifies the binaries in your GOPATH/bin are stale or up to date
 * [go-torch] - flame graph profiler for Go programs
 * [errcheck] - checks for unchecked errors in go programs
 * [interfacer] - linter that suggests interface types

### To check coverage
```sh
$ go test -coverprofile cover.report
$ go tool cover -html=cover.report -o cover.html
```

### To inspect `package`
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
 
### To get documentation
```sh
$ go doc 'github.com/krasoffski/gomill/unique' Strings
func Strings(s []string) []string
    Strings removes duplicated strings from a slice of strings. It returns a new
    slice of strings without duplicates.
```

### To perform CPU profiling
Enable profiling in your code.

```go
package main

import (
    // imports
	"os"
	"runtime/pprof"
)

func main() {
	f, _ := os.Create("multiplier.cpuprofile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
    // code
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
```
What is more, you are able to create a [Flame Graph].
```sh
$ go-torch --file=output.svg multiplier multiplier.cpuprofile 
INFO[14:49:02] Run pprof command: go tool pprof -raw -seconds 30 multiplier multiplier.cpuprofile
INFO[14:49:02] Writing svg to output.svg
$ open torch.svg
```
_Note:_ without application load CPU profile might be empty.

[binstale]: https://github.com/shurcooL/binstale
[go-torch]: https://github.com/uber/go-torch
[errcheck]: https://github.com/kisielk/errcheck
[interfacer]: https://github.com/mvdan/interfacer/
[Flame Graph]: https://github.com/brendangregg/FlameGraph