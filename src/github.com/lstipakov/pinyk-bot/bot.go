package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("305991746:AAFgGOVVhcuh7kMX77C0BQtWfVpYLDkbZLQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	log.Printf("Allowed to listen on " + os.Getenv("IP") + ":" + os.Getenv("PORT"))

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://pinykbot-lstipakov.c9users.io:" + os.Getenv("PORT") + "/" + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(os.Getenv("IP")+":"+os.Getenv("PORT"), nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
