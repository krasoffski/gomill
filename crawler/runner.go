package main

import (
	"fmt"
	"sync"
)

type Task interface {
	Process()
	Output()
}

type Manufacturer interface {
	Bufsize() int
	Create(line string) Task
	URLs() <-chan string
}

func Run(m Manufacturer, workers int) {
	var wg sync.WaitGroup
	in := make(chan Task, m.Bufsize())

	wg.Add(1)
	go func() {
		for url := range m.URLs() {
			in <- m.Create(url)
			fmt.Printf("Created task for: %s\n", url)
		}
		close(in)
		wg.Done()
	}()

	out := make(chan Task)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.Process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.Output()
	}
}
