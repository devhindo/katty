package musixmatch

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv" // todo delete later


	mxm "github.com/milindmadhukar/go-musixmatch"
	"github.com/milindmadhukar/go-musixmatch/params"
)

/*
	track.search

*/

// var url string = "https://api.musixmatch.com/ws/1.1/"

func GetLyrics(song string, artist string) (string, error) {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	client := mxm.New(os.Getenv("MUSIXMATCH_API_KEY"), http.DefaultClient)



	lyrics, err := client.GetMatcherLyrics(context.Background(), params.QueryTrack(song), params.QueryArtist(artist))

	if err != nil {
		log.Fatal(err)
	}


	return lyrics.Body , nil 
}

