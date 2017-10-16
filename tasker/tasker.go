package tasker

import (
	"sync"
)

// Task describes
type Task interface {
	Process()
	Output()
}

// Manufacturer describes
type Manufacturer interface {
	BufSize() int
	Create(line string) Task
	Items() <-chan string
}

// Run executes
func Run(m Manufacturer, workers int) {
	var wg sync.WaitGroup
	in := make(chan Task, m.BufSize())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for url := range m.Items() {
			in <- m.Create(url)
		}
		close(in)
	}()

	out := make(chan Task)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range in {
				t.Process()
				out <- t
			}
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
