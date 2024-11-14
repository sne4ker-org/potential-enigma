package main

import (
	"log"

	"github.com/gtuk/discordwebhook"
)

func main() {
	var username = "BotUser"
	var content = "This is a test message"
	var url = "https://discord.com/api/webhooks/1173335846292836493/-PbVVG-1fSDCWOLalBh0DFK0ihA4KyjXpA2c5I76tOQ5fFiPk7ImuIet8oyf807ay3r_"

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}
