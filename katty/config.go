package katty

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func config() {
	token := os.Getenv("TOKEN")
	katty, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	katty.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = katty.Open()
	if err != nil {
		log.Fatal(err)
	}

	var (
		commands = []*discordgo.ApplicationCommand{
			{
				Name: "lyrics",
				Description: "Get lyrics of a song",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type: discordgo.ApplicationCommandOptionString,
						Name: "song",
						Description: "Name of the song",
						Required:    true,
					},
			},
		},
		}
	)

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

		var song string
		
		for _, v := range i.ApplicationCommandData().Options {
			if v.Name == "song" {
				song = v.StringValue()
			}
		}
		log.Println(song)
		// respone with message contains the song name
		// response with a message that contains the song name

		response := "ask Google - I don't know the lyrics for: " + song

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