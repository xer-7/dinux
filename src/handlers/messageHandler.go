package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	switch m.Content {
		case "!ping":
    		s.ChannelMessageSend(m.ChannelID, "Pong!")
      		break
        case "!help":
	       	s.ChannelMessageSendReply(
				m.ChannelID,
				"`!mkdir (directory name)` > creates a new directory\n`!rmdir (directory name)` > removes an existing directory\n`!touch (file name)` > creates a new file\n`!rm (file name)` > removes an existing file",
				m.Reference(),
			)
      		break
        default:
	        if strings.HasPrefix(m.Content, "!mkdir ") {
				Mkdir(s, m) // ←
			} else if strings.HasPrefix(m.Content, "!rmdir ") {
				Rmdir(s, m)
			} else if strings.HasPrefix(m.Content, "!touch ") {
				Touch(s, m)
			} else if strings.HasPrefix(m.Content, "!rm ") {
				Rm(s, m)
			}
	}
	fmt.Printf("-----\nAutor: %s\nMensagem: %s\n", m.Author.Username, m.Content)
}
