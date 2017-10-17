// Package crawler implements asynchronous web sites crawler using interfaces
// from tasker and ability to specify number of workers and size of tasks' buffer.
package crawler

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
	"github.com/krasoffski/gomill/tasker"
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

// Output prints out Task result to standart output.
func (h *httpTask) Output() {
	secs := color.BlueString("[%7.2fs]", h.elapsed.Seconds())
	bsize := color.YellowString("[%10db]", h.bsize)
	fmt.Printf("%s %s %s %s\n", result[h.ok], secs, bsize, h.url)
}

type taskBuilder struct {
	source     io.Reader
	bufSize    int
	httpClient *http.Client
}

func (tb *taskBuilder) BufSize() int {
	return tb.bufSize
}

func (tb *taskBuilder) Create(url string) tasker.Task {
	h := new(httpTask)
	h.url = url
	h.client = tb.httpClient
	return h
}

func (tb *taskBuilder) Items() <-chan string {
	urls := make(chan string, tb.bufSize)
	s := bufio.NewScanner(tb.source)
	go func() {
		defer close(urls)
		for s.Scan() {

			line := strings.TrimSpace(s.Text())

			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			urls <- line
		}
	}()
	if s.Err() != nil {
		log.Fatalf("error reading: %s", s.Err())
	}
	return urls
}

// Run processes the tasks which are created internally using Create method.
// It blocks execution and waits till all tasks are completed.
func (tb *taskBuilder) Run(workers int) {
	tasker.Run(tb, workers)
}

// New creates and initializes new task builder.
func New(r io.Reader, bufSize int, client *http.Client) tasker.Builder {
	m := new(taskBuilder)
	m.source = r
	m.bufSize = bufSize
	m.httpClient = client
	return m
}
