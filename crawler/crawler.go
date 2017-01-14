// Package implements basic asynchronous web sites crawler
// with ability to specify number of workers and tasks.
// Moreover, there is an option to check moztop 500 sites as example.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var result = map[bool]string{
	true:  color.GreenString("PASS"),
	false: color.RedString("FAIL"),
}

type HTTPTask struct {
	ok      bool
	url     string
	bsize   int64
	start   time.Time
	elapsed time.Duration
}

func (h *HTTPTask) Process(timeout time.Duration) {
	// Remove creating new client for each task.
	client := http.Client{Timeout: timeout}

	h.start = time.Now()
	resp, err := client.Get(h.url)
	h.elapsed = time.Since(h.start)

	if err != nil {
		h.ok = false
		return
	}
	defer resp.Body.Close()

	// TODO: Think about handling error for Copy.
	h.bsize, _ = io.Copy(ioutil.Discard, resp.Body)

	if resp.StatusCode == http.StatusOK {
		h.ok = true
		return
	}
	h.ok = false
}

func (h *HTTPTask) Output() {
	secs := color.BlueString("[%7.2fs]", h.elapsed.Seconds())
	bsize := color.YellowString("[%10db]", h.bsize)
	fmt.Printf("%s %s %s %s\n", result[h.ok], secs, bsize, h.url)
}

type Manufacture struct {
	source  io.Reader
	bufsize int
}

func (m *Manufacture) Bufsize() int {
	return m.bufsize
}

func (m *Manufacture) Create(url string) Task {
	h := new(HTTPTask)
	h.url = url
	return h
}

func (f *Manufacture) URLs() <-chan string {
	urls := make(chan string, f.bufsize)
	s := bufio.NewScanner(f.source)
	go func() {
		defer close(urls)
		for s.Scan() {

			line := strings.TrimSpace(s.Text())

			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			urls <- line
			// fmt.Printf("Send URL: %s\n", line)
		}
	}()
	if s.Err() != nil {
		log.Fatalf("error reading: %s", s.Err())
	}
	return urls
}

func NewManufacture(r io.Reader, bufsize int) *Manufacture {
	manufacture := new(Manufacture)
	manufacture.source = r
	manufacture.bufsize = bufsize
	return manufacture
}

func main() {
	workers := flag.Int("workers", 2, "number of workers")
	bufsize := flag.Int("bufsize", 0, "size of tasks buffer")
	example := flag.Int("example", 0, "specify head of moz top, max 500")
	timeout := flag.Duration("timeout", 5*time.Second, "request timeout")

	flag.Parse()

	var reader io.Reader = os.Stdin
	var client http.Client = http.Client{Timeout: *timeout}
	if *example > 0 {
		reader = MozReader(*example, &client)
	}
	m := NewManufacture(reader, *bufsize)
	Run(m, *workers, *timeout)
}
