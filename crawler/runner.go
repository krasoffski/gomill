package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

type Task interface {
	Process()
	Output()
}

type Manufacturer interface {
	Create(lint string) Task
}

func Run(m Manufacturer, workers int) {
	var wg sync.WaitGroup
	in := make(chan Task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			text := strings.TrimSpace(s.Text())

			if text == "" || strings.HasPrefix(text, "#") {
				continue
			}
			in <- m.Create(text)
		}
		if s.Err() != nil {
			log.Fatalf("error reading STDIN: %s", s.Err())
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
