package katty

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func Config() {
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

	defer katty.Close()

	// idk what's this - but it actually keeps the app running
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}