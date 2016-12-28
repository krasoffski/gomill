package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
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
	secs := color.BlueString("%03fs", h.elapsed.Seconds())
	fmt.Printf("%s %s %s\n", result[h.ok], secs, h.url)
}

type Manufacture struct{}

func (f *Manufacture) Create(line string) Task {
	h := new(HTTPTask)
	h.url = line
	return h
}

func main() {
	workers := flag.Int("workers", 10, "number of workers")
	flag.Parse()
	m := new(Manufacture)
	Run(m, os.Stdin, *workers)
}
