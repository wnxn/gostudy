package main

import (
	"github.com/wnxn/gostudy/pkg/gopl/ch07_interfaces/ex08"
	"sort"
	"time"
)

var tracks = []*ex08.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var recent = []int{1,4,3,2,5}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	ex08.PrintTracks(tracks)
	obj := ex08.MultiTierSort{}
	obj.SetTrack(tracks)
	obj.SetRecent(recent)
	sort.Sort(obj)
	ex08.PrintTracks(obj.GetTrack())
}
