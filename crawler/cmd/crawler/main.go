package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/krasoffski/gomill/crawler"
)

func main() {
	workers := flag.Int("workers", 2, "number of workers (goroutines)")
	bufsize := flag.Int("bufsize", 0, "size of tasks' buffer")
	topsite := flag.Int("topsite", 0, "specify head of moz top sites, max 500")
	timeout := flag.Duration("timeout", 60*time.Second, "request timeout")
	threads := flag.Int("threads", 0, "numbers of OS threads, 0 runtime defaults")

	flag.Parse()

	var reader io.Reader = os.Stdin
	if *topsite > 0 {
		if *topsite < 500 {
			reader = crawler.MozReader(*topsite)
		} else {
			fmt.Fprintf(os.Stderr,
				"error value: %d, specify top sites within range [1, 500]\n",
				*topsite)
			os.Exit(1)
		}
	}
	// This call might be removed in the future.
	runtime.GOMAXPROCS(*threads)
	m := crawler.New(reader, *bufsize, &http.Client{Timeout: *timeout})
	start := time.Now()
	m.Run(*workers)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
