package main

import "time"

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

var tracks = []*Track {
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m35s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type byTwoKeys struct {
	Tracks []*Track
	PrimaryKey string
	SecondaryKey string
}

func (btk byTwoKeys) Len() int {
	return len(btk.Tracks)
}

func (btk byTwoKeys) Less(i, j int) bool {
	return true
}

func main() {
	//TODO
}