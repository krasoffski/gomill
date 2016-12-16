package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	factor := flag.Int("factor", 0, "multiplier for enlarger")
	flag.Parse()

	ctx := NewContext()

	counter := NewCounter(ctx, &wg)
	c := counter.GetSource()

	enlarger := NewEnlarger(ctx, &wg, c, *factor)
	r := enlarger.GetSource()

	fmt.Println(<-r)
	fmt.Println(<-r)
	fmt.Println(<-r)
	ctx.Stop()
	wg.Wait()
}
