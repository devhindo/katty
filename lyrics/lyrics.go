package lyrics

import (
	"errors"
	"strings"

	"github.com/devhindo/katty/genius"
	"github.com/devhindo/katty/musixmatch"
)

func FindLyrics(song string, artist string) (string, error) {

	lyrics, err := musixmatch.GetLyrics(song, artist)
	if err != nil {
		lyrics, err = genius.GetLyrics(song, artist)
		if err != nil {
			return "", errors.New("no lyrics found (sorry) - `" + song + "-" + artist +"`")
		}
	}
	
	i := strings.Index(lyrics, "...")
	if i == -1 {
		return lyrics, nil
	}

	

	return lyrics[:i], nil
}
