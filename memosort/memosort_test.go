package memosort

import (
	"testing"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Bohemian Rhapsody", "Queen", "A Night at the Opera", 1975, length("6m6s")},
	{"Smells Like Teen Spirit", "Nirvana", "Nevermind", 1972, length("4m37s")},
	{"Imagine", "John Lennon", "Imagine", 1971, length("3m54s")},
	{"Hotel California", "Eagles", "Hotel California", 1976, length("6m40s")},
	{"One", "Metallica", "And Justice for All", 1988, length("7m23s")},
	{"Comfortably Numb", "Pink Floyd", "The Wall", 1979, length("6m53s")},
	{"Like a Rolling Stone", "Bob Dylan", "Highway 61 Revisited", 1965, length("6m20s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func TestStableSimple(t *testing.T) {

}
