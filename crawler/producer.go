package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func Producer(m Manufacturer, wg *sync.WaitGroup) chan Tasker {
	in := make(chan Tasker)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- m.Create(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()
	return in
}
