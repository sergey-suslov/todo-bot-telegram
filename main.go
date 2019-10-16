package main

import (
	"log"
	"github.com/sergey-suslov/todo-bot-telegram/botapi"
	"github.com/sergey-suslov/todo-bot-telegram/db"
)

func main() {
	log.Println("Hello")
	_, err := db.GetConnection()
	if err != nil {
		log.Println("Couldn't connect to DB")
	}
	bot, err := botapi.GetBot()
	if err != nil {
		log.Fatal(err)
	}
	botapi.PlayEcho(bot)
}
