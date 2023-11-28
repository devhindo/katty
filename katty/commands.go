package katty

import (

	"github.com/bwmarrin/discordgo"
)

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