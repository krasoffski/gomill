// race contains data race with resetting the timer.
// Issue "go run -race race.go" to get warring about data race.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Version of race example.
var Version = "0.1.0"

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

func main() {
	version := flag.Bool("version", false, "show version")
	flag.Parse()

	if *version {
		fmt.Printf("Version: %s\n", Version)
		return
	}

	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}
