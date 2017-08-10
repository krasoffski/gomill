// Package implements asynchronous web sites crawler
// with ability to specify number of workers and size of tasks' buffer.
// Moreover, there is an option to check moztop 500 sites as an example.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

var result = map[bool]string{
	true:  color.GreenString("PASS"),
	false: color.RedString("FAIL"),
}

// httpTask represents HTTP task with required for processing fields.
// ok idecates that url is fetched without issues within timeout
// url is URL to fetch
// bsize is body size of response
// start is start time of processing
// client is http.Client with timeout
// elapsed is elapsed time for processing
type httpTask struct {
	ok      bool
	url     string
	bsize   int64
	start   time.Time
	client  *http.Client
	elapsed time.Duration
}

// Process processes and fills required fields of HTTPTask.
func (h *httpTask) Process() {
	h.start = time.Now()
	resp, err := h.client.Get(h.url)
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
func (h *httpTask) Output() {
	secs := color.BlueString("[%7.2fs]", h.elapsed.Seconds())
	bsize := color.YellowString("[%10db]", h.bsize)
	fmt.Printf("%s %s %s %s\n", result[h.ok], secs, bsize, h.url)
}

type manufacture struct {
	source     io.Reader
	bufSize    int
	httpClient *http.Client
}

func (m *manufacture) Bufsize() int {
	return m.bufSize
}

func (m *manufacture) Create(url string) Task {
	h := new(httpTask)
	h.url = url
	h.client = m.httpClient
	return h
}

func (m *manufacture) URLs() <-chan string {
	urls := make(chan string, m.bufSize)
	s := bufio.NewScanner(m.source)
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

func NewManufacture(r io.Reader, bufSize int, client *http.Client) *manufacture {
	m := new(manufacture)
	m.source = r
	m.bufSize = bufSize
	m.httpClient = client
	return m
}
