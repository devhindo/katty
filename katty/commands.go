package katty

import (
	"errors"
	"fmt"
	"strings"

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
					Name: "prompt",
					Description: "follow this format: \"song-artist\"",
					Required:    true,
				},
		},
	},
	}
)


func processLryicsCommand(p string) (string, string, error) {
	fmt.Println(p)
	s := strings.Split(p, "-")

	if len(s) != 2 {
		fmt.Println(len(s), s)
		return "", "", errors.New("invalid prompt | follow this format: `song-artist`")
	}

	song, artist := s[0], s[1]

	return song, artist, nil
}