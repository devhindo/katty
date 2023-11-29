package genius

import (

)

func GetLyrics(song string, artist string) (string, error) {
	return "genius: ask Google for now" + "`" + song +"-" + artist + "`" , nil
}