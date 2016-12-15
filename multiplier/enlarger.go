package main

import (
	"fmt"
	"sync"
)

type Enlarger struct {
	in  <-chan int
	out chan int
}

func (e *Enlarger) GetSource() <-chan int {
	return e.out
}

func NewEnlarger(ctx *Context, wg *sync.WaitGroup,
	in <-chan int, scale int) *Enlarger {
	e := new(Enlarger)
	e.in = in
	e.out = make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		done := ctx.GetDone()
		for {
			select {
			case i := <-e.in:
				select {
				case e.out <- i * scale:
				case <-done:
					fmt.Println("Enlarger terminated in")
					return
				}
			case <-done:
				fmt.Println("Enlarger terminated out")
				return
			}
		}
	}()
	return e
}
