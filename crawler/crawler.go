package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

var result = map[bool]string{
	true:  color.GreenString("PASS"),
	false: color.RedString("FAIL"),
}

type HTTPTask struct {
	url string
	ok  bool
}

func (h *HTTPTask) Process() {
	resp, err := http.Get(h.url)
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
	fmt.Printf("%s %s\n", result[h.ok], h.url)
}

type Manufacture struct{}

func (f *Manufacture) Create(line string) Task {
	h := new(HTTPTask)
	h.url = line
	return h
}

func main() {
	workers := flag.Int("workers", 1000, "number of workers")
	flag.Parse()
	m := new(Manufacture)
	Run(m, *workers)
}
