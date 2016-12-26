package crawler

import (
	"fmt"
	"net/http"
)

type task interface {
	process()
	output()
}

type factory interface {
	create(lint string) task
}

type HTTPTask struct {
	url string
	ok  bool
}

func (h *HTTPTask) process() {
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

func (h *HTTPTask) output() {
	fmt.Printf("%s %t\n", h.url, h.ok)
}

type Factory struct{}

func (f *Factory) create(line string) task {
	h := new(HTTPTask)
	h.url = line
	return h
}
