// Package tasker provide high level interfaces for creating asynchronous
// task creation and execution.
package tasker

import (
	"sync"
)

// Task describes minimal interface for simple tasks.
type Task interface {
	Process()
	Output()
}

// Builder describes interface for creating tasks from slice of strings.
type Builder interface {
	BufSize() int
	Create(line string) Task
	Items() <-chan string
}

// Run creates tasks from slice of strings and performs this task.
// After task completion method Output is invoked for each tasks.
// Argument workers respresents number of goroutines for performing tasks.
func Run(m Builder, workers int) {
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
