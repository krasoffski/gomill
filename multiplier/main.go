package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ctx := NewContext()
	counter := NewConunter(ctx, &wg)
	c := counter.GetSource()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	ctx.Stop()
	wg.Wait()
}
