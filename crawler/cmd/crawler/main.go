package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/krasoffski/gomill/crawler"
)

// mozTop creates new reader which represents list of sites addresses from
// https://moz.com/top500 delimeted by new line charaster.
// Request timeout is 60s.
func mozTop(head int) io.Reader {
	// TODO: Create a better solution for providing top sites.

	client := http.Client{Timeout: time.Second * 60}
	fmt.Printf("Getting csv with top 500 from moz.com... ")
	start := time.Now()
	resp, err := client.Get("https://moz.com/top500/domains/csv")

	if err != nil {
		log.Fatalf("error performing request %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error getting moz top 500: %s\n", resp.Status)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/csv") {
		log.Fatalln("error getting csv data, wrong conten type")
	}
	defer resp.Body.Close()

	urls := make([]string, head)
	reader := csv.NewReader(resp.Body)

	for i := 0; i <= head; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading csv %s\n", err)
		}
		url := "http://" + strings.TrimRight(record[1], "/")

		urls = append(urls, url)
	}

	// Do not count strings join and new reader creation.
	fmt.Printf("done in %.2fs\n", time.Since(start).Seconds())
	// Skipping csv header.
	return strings.NewReader(strings.Join(urls[1:], "\n"))
}

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
			reader = mozTop(*topsite)
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
