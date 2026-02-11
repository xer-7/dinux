package handlers

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Rmdir(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	parts := strings.Fields(m.Content)
	if len(parts) < 2 {
		s.ChannelMessageSendReply(
			m.ChannelID,
		    "Error!\nUsage: `!rmdir (directory name)`",
			m.Reference(),
		)
		return
	}

	dirName := parts[1]

	if dirs[dirName] {
		delete(dirs, dirName)
		s.ChannelMessageSendReply(
			m.ChannelID,
			"Sucessfully deleted directory!",
			m.Reference(),
		)
	} else {
		s.ChannelMessageSendReply(
			m.ChannelID,
			"Theres no directory with that name!",
			m.Reference(),
		)
	}
}
