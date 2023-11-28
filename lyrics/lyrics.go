package lyrics

import (
	"errors"

	"github.com/devhindo/katty/genius"
	"github.com/devhindo/katty/musixmatch"
)

func FindLyrics(song string, artist string) (string, error) {
	lyrics, err := genius.GetLyrics(song, artist)
	if err != nil {
		lyrics, err = musixmatch.GetLyrics(song, artist)
		if err != nil {
			return "", errors.New("no lyrics found (sorry)")
		}
	}
	return lyrics, nil
}
