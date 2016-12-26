package main

import "sync"

func Consumer(in chan Tasker, wg *sync.WaitGroup, workers int) {
	out := make(chan Tasker)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.Process()
				out <- t
			}
		}()
		wg.Done()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.Output()
	}
}
