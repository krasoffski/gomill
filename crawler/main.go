package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	workers := flag.Int("workers", 2, "number of workers")
	bufsize := flag.Int("bufsize", 0, "size of tasks buffer")
	topsite := flag.Int("topsite", 0, "specify head of moz top sites, max 500")
	timeout := flag.Duration("timeout", 60*time.Second, "request timeout")

	flag.Parse()

	var reader io.Reader = os.Stdin
	if *topsite > 0 {
		if *topsite < 500 {
			reader = MozReader(*topsite)
		} else {
			fmt.Fprintf(os.Stderr,
				"error value: %d, specify top sites within range [1, 500]\n",
				*topsite)
			os.Exit(1)
		}
	}
	m := NewManufacture(reader, *bufsize, &http.Client{Timeout: *timeout})
	Run(m, *workers)
}
