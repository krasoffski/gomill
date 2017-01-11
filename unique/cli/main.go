package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/krasoffski/gomill/unique"
)

func run() int {
	seq := os.Args[1:]
	if len(seq) == 0 {
		os.Stderr.WriteString("got empry sequence of strings\n")
		return 1
	}
	fmt.Println(strings.Join(unique.Strings(seq), " "))
	return 0
}

func main() {
	os.Exit(run())
}
