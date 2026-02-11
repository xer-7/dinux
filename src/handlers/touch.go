package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var files = make(map[string]bool)

func Touch(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	parts := strings.Fields(m.Content)
	if len(parts) < 2 {
		s.ChannelMessageSendReply(m.ChannelID, "Error!\nUsage: `!touch (file name)`", m.Reference())
	}
	fileName := parts[1]
	if strings.Contains(fileName, "@") {
		s.ChannelMessageSendReply(
			m.ChannelID,
			"Dont try to bypass me.\n||Youre not elliot alderson||",
			m.Reference(),
		)
		return
	}
	if files[fileName] {
		s.ChannelMessageSendReply(
			m.ChannelID,
			"Theres another directory with that name!",
			m.Reference(),
		)
		return
	}
	files[fileName] = true
	response := fmt.Sprintf("You created a file!\nName: %s", fileName)
	s.ChannelMessageSendReply(m.ChannelID, response, m.Reference())
}
