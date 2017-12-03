//go:generate go build unique.go
//go:generate ./unique string
package main

import (
	"fmt"
	"os"
	"strings"
)

func run() int {
	seq := os.Args[1:]
	if len(seq) == 0 {
		os.Stderr.WriteString("got empty sequence of strings\n")
		return 1
	}
	fmt.Println(strings.Join(Dedup(seq), " "))
	return 0
}

func main() {
	os.Exit(run())
}
