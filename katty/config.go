package katty

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/devhindo/katty/lyrics"
)

func config() {
	token := os.Getenv("TOKEN")
	katty, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	
	activity := discordgo.Activity{
		Name: "/lyrics song-artist",
		Type: discordgo.ActivityTypeCustom,
	}
	
	status := discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{&activity},
	}
	
	err = katty.UpdateStatusComplex(status)
	if err != nil {
		log.Fatal(err)
	}
	
	katty.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = katty.Open()
	if err != nil {
		log.Fatal(err)
	}


	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}

	_, err = katty.ApplicationCommandBulkOverwrite(os.Getenv("APP_ID"), "", commands)

	if err != nil {
		log.Fatal(err)
	}

	katty.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Handle interaction here
		deferredResponse := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		}
		s.InteractionRespond(i.Interaction, deferredResponse)

		var prompt string

		for _, v := range i.ApplicationCommandData().Options {
			if v.Name == "prompt" {
				prompt = v.StringValue()
			}
		}
		log.Println(prompt)
		song, artist, err := processLryicsCommand(prompt)

		var response string

		if err != nil {
			response = fmt.Sprintf("%s", err)
		} else {
			lyrics, err := lyrics.FindLyrics(song, artist)

			if err != nil {
				response = "no lyrics found (sorry)"
			}
			response = lyrics
		}


		s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: response,
		})
		
		log.Println("interaction received")
	})

	defer katty.Close()

	// idk what's this - but it actually keeps the app running
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}