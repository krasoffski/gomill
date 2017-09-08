package memosort

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"text/tabwriter"
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

	// {"Comfortably Numb", "Pink Floyd", "The Wall", 1979, length("6m53s")},
	// {"Like a Rolling Stone", "Bob Dylan", "Highway 61 Revisited", 1965, length("6m20s")},
	// {"Bohemian Rhapsody", "Queen", "A Night at the Opera", 1975, length("6m6s")},
	// {"Smells Like Teen Spirit", "Nirvana", "Nevermind", 1972, length("4m37s")},
	// {"Imagine", "John Lennon", "Imagine", 1971, length("3m54s")},
	// {"Hotel California", "Eagles", "Hotel California", 1976, length("6m40s")},
	// {"One", "Metallica", "And Justice for All", 1988, length("7m23s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	fmt.Println()
	tw.Flush()
}

func TestStableSimple(t *testing.T) {
	tracksMemo := make([]Track, len(tracks))
	copy(tracksMemo, tracks)

	printTracks(tracksMemo)

	m := New()
	m.By(func(i, j int) bool { return tracksMemo[i].Title < tracksMemo[j].Title })
	m.By(func(i, j int) bool { return tracksMemo[i].Year < tracksMemo[j].Year })
	m.By(func(i, j int) bool { return tracksMemo[i].Length < tracksMemo[j].Length })
	sort.Slice(tracksMemo, m.Less)
	printTracks(tracksMemo)
}
