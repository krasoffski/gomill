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

[binstale]: https://github.com/shurcooL/binstale
[go-torch]: https://github.com/uber/go-torch
[errcheck]: https://github.com/kisielk/errcheck
[interfacer]: https://github.com/mvdan/interfacer/