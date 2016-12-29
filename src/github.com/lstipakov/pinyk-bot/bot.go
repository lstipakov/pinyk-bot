package main

import (
    "log"
    "gopkg.in/telegram-bot-api.v4"
)

func main() {
    bot, err := tgbotapi.NewBotAPI("305991746:AAFgGOVVhcuh7kMX77C0BQtWfVpYLDkbZLQ")
    if err != nil {
        log.Panic(err)
    }
    
    bot.Debug = true
    
    log.Printf("Authorized on account %s", bot.Self.UserName)
    
    msg := tgbotapi.NewMessage(290659018, "Hello from Cloud9!")
    bot.Send(msg)
}