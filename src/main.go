package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"log"
	"github.com/joho/godotenv"
	"os/signal"
	"syscall"
	"bot/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}


	var Token string = os.Getenv("DISCORD_TOKEN")
	if Token == "" {
		fmt.Println("[-] Erro: Nenhum token fornecido. Verifique a .env")
		return
	}
	root, err := discordgo.New("Bot " + Token)

	root.AddHandler(handlers.OnReady)
	root.AddHandler(handlers.HandleMessage)

	err = root.Open()
	if err != nil {
		log.Fatal("Erro ao criar sessão: ", err)
		return
	}
	defer root.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	defer root.Close()
}
