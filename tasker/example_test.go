package tasker_test

import (
	"fmt"
	"strings"

	"github.com/krasoffski/gomill/tasker"
)

type upperTask struct {
	in, out string
}

// Process processes and fills required fields of upperTask.
func (u *upperTask) Process() {
	u.out = strings.ToUpper(u.in)
}

// Output prints out Task result to standart output.
func (u *upperTask) Output() {
	fmt.Printf("%s ", u.out)
}

type taskBuilder struct {
	words   []string
	bufSize int
}

func (tb *taskBuilder) BufSize() int {
	return tb.bufSize
}

func (tb *taskBuilder) Create(s string) tasker.Task {
	return &upperTask{in: s}
}

// Items prepares strings from slice and send them to chan.
func (tb *taskBuilder) Items() <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, word := range tb.words {
			if word == "" {
				continue
			}
			ch <- word
		}
	}()
	return ch
}

// Run performs all tasks using provided number of goroutines and buffer size.
func (tb *taskBuilder) Run(workers int) {
	tasker.Run(tb, workers)
}

// New creates and initializes new task builder.
func New(input []string, bufSize int) tasker.Builder {
	return &taskBuilder{words: input, bufSize: bufSize}
}

func Example() {
	input := []string{"apple", "", "orange", "", "cherry"}
	u := New(input, 10)
	// NOTE: the output order might be different due to async processing.
	u.Run(5)
	// Output: APPLE ORANGE CHERRY
}
