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

func (h *HTTPTask) Process() {
	// fmt.Printf("Got task for: %s\n", h.url)
	h.start = time.Now()
	// TODO: Add timeout for http get request.
	resp, err := http.Get(h.url)
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

	flag.Parse()

	var reader io.Reader = os.Stdin
	if *example > 0 {
		reader = MozReader(*example)
	}
	m := NewManufacture(reader, *bufsize)
	Run(m, *workers)
}
