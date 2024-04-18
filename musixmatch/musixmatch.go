package musixmatch

import (
	"context"
	"fmt"
	"net/http"
	"os"

	//"github.com/joho/godotenv" // todo delete later

	mxm "github.com/milindmadhukar/go-musixmatch"
	"github.com/milindmadhukar/go-musixmatch/params"
)

/*
	track.search

*/

// var url string = "https://api.musixmatch.com/ws/1.1/"

func GetLyrics(song string, artist string) (string, error) {

  
	client := mxm.New(os.Getenv("MUSIXMATCH_API_KEY"), http.DefaultClient)



	lyrics, err := client.GetMatcherLyrics(context.Background(), params.QueryTrack(song), params.QueryArtist(artist))

	if err != nil {
		return "no lyrics found for this song", nil
	}

	fmt.Println("Lyrics found for the song: ", song, " by ", artist)

	return lyrics.Body , nil 
}

