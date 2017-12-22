package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track tracks recoders
type Track struct {
	Title, Artist, Album string
	Year                 int
	Length               time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", " From the Roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type byArtist []*Track

func (artists byArtist) Len() int {
	return len(artists)
}

func (artists byArtist) Less(i, j int) bool {
	return artists[i].Artist < artists[j].Artist
}

func (artists byArtist) Swap(i, j int) {
	artists[i], artists[j] = artists[j], artists[i]
}

func main() {
	fmt.Println("ORIGIN TRACKS")
	printTracks(tracks)

	sort.Sort(byArtist(tracks))
	fmt.Println("Sorted by Artist")
	printTracks(tracks)

	sort.Sort(sort.Reverse(byArtist(tracks)))
	fmt.Println("Reverse Sorted by Artist")
	printTracks(tracks)

	fmt.Println("Custom Sort")
	customFunc := func(a, b *Track) bool {
		if a.Title != b.Title {
			return a.Title < b.Title
		}
		if a.Year != b.Year {
			return a.Year < b.Year
		}
		if a.Length != b.Length {
			return a.Length < b.Length
		}
		return false
	}
	custom := customSort{
		tracks,
		customFunc,
	}
	sort.Sort(custom)
	printTracks(tracks)
}

func printTracks(ts []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "----", "----", "----", "----", "----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int {
	return len(c.t)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}
