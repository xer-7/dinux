package handlers

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Rm(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	parts := strings.Fields(m.Content)
	if len(parts) < 2 {
		s.ChannelMessageSendReply(
			m.ChannelID,
		    "Error!\nUsage: `!rm (file name)`",
			m.Reference(),
		)
		return
	}
	fileName := parts[1]
	if files[fileName] {
		delete(files, fileName)
		s.ChannelMessageSendReply(
			m.ChannelID,
			"Sucessfully deleted file!",
			m.Reference(),
		)
	} else {
		s.ChannelMessageSendReply(
			m.ChannelID,
			"Theres no file with that name!",
			m.Reference(),
		)
	}
}
