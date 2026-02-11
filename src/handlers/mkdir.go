package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var dirs = make(map[string]bool)

func Mkdir(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	parts := strings.Fields(m.Content)
	if len(parts) < 2 {
		s.ChannelMessageSendReply(m.ChannelID, "Error!\nUsage: `!mkdir (directory name)`", m.Reference())
	}
	dirName := parts[1]

	if strings.Contains(dirName, "@") {
		s.ChannelMessageSendReply(
			m.ChannelID,
			"Dont try to bypass me.\n||Youre not elliot alderson||",
			m.Reference(),
		)
		return
	}

	if dirs[dirName] {
		s.ChannelMessageSendReply(
			m.ChannelID,
			"Theres another directory with that name!",
			m.Reference(),
		)
		return
	}

	dirs[dirName] = true
	response := fmt.Sprintf("You created a directory!\nName: %s", dirName)
	s.ChannelMessageSendReply(m.ChannelID, response, m.Reference())
}
