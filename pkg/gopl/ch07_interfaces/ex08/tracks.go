package ex08

import (
	"fmt"
	"github.com/golang/glog"
	"os"
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

type MultiTierSort struct{
	t []*Track
	recentClick []int
}

func(x MultiTierSort)GetTrack()[]*Track{
	return x.t
}

func(x MultiTierSort)GetRecent()[]int{
	return x.recentClick
}

func(x *MultiTierSort)SetRecent(y []int){
	x.recentClick = y
}

func(x *MultiTierSort)SetTrack(y []*Track){
	x.t = y
}

func(x MultiTierSort)Len()int{
	return len(x.t)
}

func(x MultiTierSort)Swap(i,j int){
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func(x MultiTierSort)Less(i,j int)bool{
	for _, value:=range x.recentClick{
		switch value{
		case 1:
			if x.t[i].Title == x.t[j].Title{
				continue
			}else{
				return x.t[i].Title <x.t[j].Title
			}
		case 2:
			if x.t[i].Artist == x.t[j].Artist{
				continue
			}else{
				return x.t[i].Artist <x.t[j].Artist
			}
		case 3:
			if x.t[i].Album == x.t[j].Album{
				continue
			}else{
				return x.t[i].Album <x.t[j].Album
			}
		case 4:
			if x.t[i].Year == x.t[j].Year{
				continue
			}else{
				return x.t[i].Year <x.t[j].Year
			}
		case 5:
			if x.t[i].Length == x.t[j].Length{
				continue
			}else{
				return x.t[i].Length <x.t[j].Length
			}
		default:
			glog.Error("Switch run to default")
			return false
		}
	}
	return false
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}