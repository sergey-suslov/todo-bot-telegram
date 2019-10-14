package main

import (
	"log"
	"github.com/sergey-suslov/todo-bot-telegram/db"
)


func main()  {
	log.Println("Hello")
	_, err := db.GetConnection()
	if err != nil {
		log.Panic(err)
	}
}