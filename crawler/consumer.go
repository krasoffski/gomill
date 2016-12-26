package crawler

import "sync"

func consumer(in chan task, wg *sync.WaitGroup, workers int) {
	out := make(chan task)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
		}()
		wg.Done()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}
