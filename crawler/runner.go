package crawler

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func run(f factory, workers int) {
	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.create(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()
}