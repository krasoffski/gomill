// Package implements asynchronous web sites crawler
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

// HTTPTask represents HTTP task with required for processing fields.
// ok idecates that url is fetched without issues within timeout
// url is URL to fetch
// bsize is body size of responce
// start is start time of processing
// timeout is timeout for get request
// elapsed is elapsed time for processing
type HTTPTask struct {
	ok      bool
	url     string
	bsize   int64
	start   time.Time
	timeout time.Duration
	elapsed time.Duration
}

// Process processes and fills required fields of HTTPTask.
func (h *HTTPTask) Process() {
	// Remove creating new client for each task in case of global timeout.
	client := http.Client{Timeout: h.timeout}

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

// Output prints out HTTPTask result to standart output.
func (h *HTTPTask) Output() {
	secs := color.BlueString("[%7.2fs]", h.elapsed.Seconds())
	bsize := color.YellowString("[%10db]", h.bsize)
	fmt.Printf("%s %s %s %s\n", result[h.ok], secs, bsize, h.url)
}

type Manufacture struct {
	source      io.Reader
	bufsize     int
	taskTimeout time.Duration
}

func (m *Manufacture) Bufsize() int {
	return m.bufsize
}

func (m *Manufacture) Create(url string) Task {
	h := new(HTTPTask)
	h.url = url
	h.timeout = m.taskTimeout
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

func NewManufacture(r io.Reader, bufsize int, taskTimeout time.Duration) *Manufacture {
	manufacture := new(Manufacture)
	manufacture.source = r
	manufacture.bufsize = bufsize
	manufacture.taskTimeout = taskTimeout
	return manufacture
}

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
				"error value: %d, specify top sites within range [1, 500]\n", *topsite)
			os.Exit(1)
		}
	}
	m := NewManufacture(reader, *bufsize, *timeout)
	Run(m, *workers)
}
