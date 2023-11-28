package lyrics

import (

)

func FindLyrics(song string, artist string) (string, error) {
	s := "ask Google - I don't know the lyrics for: \nsong: " + song + "\n artist: " + artist
	return s, nil
}
