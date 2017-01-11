// Package implement example of context based
// asynchronous number multiplier.
package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	factor := flag.Int("factor", 0, "multiplier for enlarger")
	number := flag.Int("number", 1, "how many numbers to print")
	flag.Parse()

	ctx := NewContext()

	counter := NewCounter(ctx, &wg)
	c := counter.GetSource()

	enlarger := NewEnlarger(ctx, &wg, c, *factor)
	r := enlarger.GetSource()

	for i := 0; i < *number; i++ {
		fmt.Println(<-r)
	}
	ctx.Stop()
	wg.Wait()
}
