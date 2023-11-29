package main

import (
	//"github.com/devhindo/katty/katty"
	"fmt"

	"github.com/devhindo/katty/lyrics"
)

func main() {
	//katty.Run()
	tests := []struct {
		song   string
		artist string
	}{
		{"Meds", "Shehab"},
		//{"", ""},
		//{"", ""},
		//{"", ""},
		//{"", ""},
	}

	for _, test := range tests {
		lyrics, err := lyrics.FindLyrics(test.song, test.artist)
		if err != nil {
			println(err.Error())
		}
		fmt.Println(lyrics)
	}
}
