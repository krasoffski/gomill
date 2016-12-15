package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ctx := NewContext()

	counter := NewCounter(ctx, &wg)
	c := counter.GetSource()

	enlarger := NewEnlarger(ctx, &wg, c, 2)
	r := enlarger.GetSource()

	fmt.Println(<-r)
	fmt.Println(<-r)
	fmt.Println(<-r)
	ctx.Stop()
	wg.Wait()
}
