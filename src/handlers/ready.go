package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateGameStatus(0, "ORDEM")
	fmt.Printf("[+] %s is online!\n", s.State.User)
	fmt.Printf("[*] Id: %d\n", s.State.User.ID)
	fmt.Printf("[*] Running on: %d servers\n", len(r.Guilds))
}
