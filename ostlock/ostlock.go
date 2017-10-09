// Package contains os thread lock example for goroutine execution.
package main

/*
#cgo LDFLAGS: -L${SRCDIR}
#cgo LDFLAGS: -lcheck_thread_id -lpthread

extern void init_id();
extern void check_id();
*/
import "C"
import (
	"bytes"
	"fmt"
	"math"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
)

func init() {
	runtime.GOMAXPROCS(1)
	// 	runtime.LockOSThread()
}

func main() {
	fmt.Println("main thread = ", syscall.Gettid())

	done := make(chan bool, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		fmt.Println("application received signal: ", sig)
		done <- true
	}()

	runtime.LockOSThread()

	C.init_id()
	defer C.check_id()

	fmt.Println("main thread = ", syscall.Gettid())
	defer fmt.Println("finish main thread = ", syscall.Gettid())

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(number int, wg *sync.WaitGroup) {
			longCalculation(number)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()

	fmt.Println("main thread = ", syscall.Gettid())

	defer runtime.UnlockOSThread()

	longCalculation(1)

	fmt.Println("main thread = ", syscall.Gettid())
	fmt.Println("Press CTRL-C...")
	<-done
	return
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// Emulation of long calculation function.
func longCalculation(number int) {
	fmt.Printf("+ number = %d, goid = %v, thread = ", number, getGID())
	fmt.Println(syscall.Gettid())

	score := 1.0
	for i := 0.0; i < 10000; i++ {
		for j := 0.0; j < 10000; j++ {
			a := math.Hypot(i, j)
			if a > 1 {
				score *= i
				_ = score
			}
			_ = a
		}
	}

	fmt.Printf("- number = %d, thread = ", number)
	fmt.Println(syscall.Gettid())
}
