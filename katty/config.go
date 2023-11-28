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

	log.Println("adding commands..")

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := katty.ApplicationCommandCreate(katty.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
		log.Printf("Command '%v' registered", cmd.Name)
	}
	defer katty.Close()

	// idk what's this - but it actually keeps the app running
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
