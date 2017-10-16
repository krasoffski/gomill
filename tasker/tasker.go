// Package tasker provides high level interfaces for asynchronous
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
	Run(workers int)
}

// Run creates tasks from slice of strings and performs this task.
// After task completion method Output is invoked for each tasks.
// Argument workers respresents number of goroutines for performing tasks.
// This function can used as a reference implemention of Run method for
// satisfying Builder interface.
func Run(b Builder, workers int) {
	var wg sync.WaitGroup
	in := make(chan Task, b.BufSize())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for url := range b.Items() {
			in <- b.Create(url)
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
