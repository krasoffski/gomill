package main

import (
	"fmt"
	"net/http"
)

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
	fmt.Printf("%s %t\n", h.url, h.ok)
}

type Manufacture struct{}

func (f *Manufacture) Create(line string) Task {
	h := new(HTTPTask)
	h.url = line
	return h
}

func main() {
	m := new(Manufacture)
	Run(m, 100)
}
