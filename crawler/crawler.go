package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
	start   time.Time
	elapsed time.Duration
}

func (h *HTTPTask) Process() {

	h.start = time.Now()
	resp, err := http.Get(h.url)
	h.elapsed = time.Since(h.start)

	if err != nil {
		h.ok = false
		return
	}
	if resp.StatusCode == http.StatusOK {
		h.ok = true
		return
	}
	h.ok = false
}

func (h *HTTPTask) Output() {
	secs := color.BlueString("[%6.2fs]", h.elapsed.Seconds())
	fmt.Printf("%s %s %s\n", result[h.ok], secs, h.url)
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
			fmt.Printf("Send URL: %s\n", line)
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
	workers := flag.Int("workers", 10, "number of workers")
	bufsize := flag.Int("bufsize", 10, "size of tasks buffer")
	flag.Parse()
	// TODO: Add builder for mufacture?
	m := NewManufacture(os.Stdin, *bufsize)
	Run(m, *workers)
}
