package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	var isCommand bool
	if m.Author.Bot {
		return
	}
	switch m.Content {
		case "!ping":
 			isCommand = true
    		s.ChannelMessageSend(m.ChannelID, "Pong!")
      		break
	}
	fmt.Printf("-----\nAutor: %s\nMensagem: %s\nIsCommand: %v\n-----", m.Author.Username, m.Content, isCommand)
}
