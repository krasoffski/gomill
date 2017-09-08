package memosort

import (
	"sort"
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

var tracks = []Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Comfortably Numb", "Pink Floyd", "The Wall", 1979, length("6m53s")},
	{"Like a Rolling Stone", "Bob Dylan", "Highway 61 Revisited", 1965, length("6m20s")},
	{"Bohemian Rhapsody", "Queen", "A Night at the Opera", 1975, length("6m6s")},
	{"Smells Like Teen Spirit", "Nirvana", "Nevermind", 1972, length("4m37s")},
	{"Imagine", "John Lennon", "Imagine", 1971, length("3m54s")},
	{"Hotel California", "Eagles", "Hotel California", 1976, length("6m40s")},
	{"One", "Metallica", "And Justice for All", 1988, length("7m23s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func TestMemoSimple(t *testing.T) {
	te := make([]Track, len(tracks)) // tracks expected
	tm := make([]Track, len(tracks)) // tracks memosort
	copy(te, tracks)
	copy(tm, tracks)

	ms := New() // memo sort
	ms.By(func(i, j int) bool { return tm[i].Title < tm[j].Title })
	ms.By(func(i, j int) bool { return tm[i].Year < tm[j].Year })
	ms.By(
		func(i, j int) bool { return tm[i].Length < tm[j].Length },
		func(i, j int) bool { return tm[i].Artist < tm[j].Artist },
		func(i, j int) bool { return tm[i].Album < tm[j].Album },
	)
	sort.Slice(tm, ms.Less)

	sort.Slice(te, func(i, j int) bool {
		if te[i].Title != te[j].Title {
			return te[i].Title < te[j].Title
		}
		if te[i].Year != te[j].Year {
			return te[i].Year < te[j].Year
		}
		if te[i].Length != te[j].Length {
			return te[i].Length < te[j].Length
		}
		if te[i].Artist != te[j].Artist {
			return te[i].Artist < te[j].Artist
		}
		if te[i].Album != te[j].Album {
			return te[i].Album < te[j].Album
		}
		return false
	})

	for i, track := range te {
		if track != tm[i] {
			t.Errorf("i: %v, expected %v, got %v", i, track, tm[i])
		}
	}
}
