package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	ctx *Context
	c   chan int
	i   int
}

func NewConunter(ctx *Context, wg *sync.WaitGroup) *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.ctx = ctx

	wg.Add(1)
	go func() {
		defer wg.Done()
		done := counter.ctx.GetDone()
		for {
			select {
			case counter.c <- counter.i:
				counter.i += 1
			case <-done:
				// Add logging here.
				fmt.Println("Counter terminated")
				return
			}
		}
	}()

	return counter
}

func (c *Counter) GetSource() <-chan int {
	return c.c
}
