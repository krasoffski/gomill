package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/krasoffski/gomill/lru1"
)

func main() {
	//	capacity := flag.Int("size", 1000, "size of the cache")
	//	flag.Parse()

	cache := lru1.NewCache(1000)

	count := 0
	miss := 0

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		word := in.Text()
		if len(word) > 3 {
			three := word[1:4]
			if cache.Get(three) == nil {
				cache.Put(three, word)
				miss++
			}
			count++
		}
	}
	fmt.Printf("%d total %d missed\n", count, miss)
}
