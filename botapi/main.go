package botapi

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
)

type Task struct {
	Text string
}

func mapTasks(vs []Task, f func(Task) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// GetBot Created new bot
func GetBot() (*tgbotapi.BotAPI, error) {
	log.Println("TOKEN", os.Getenv("TELEGRAM_APITOKEN"))
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot, nil
}

func getSender(bot *tgbotapi.BotAPI, chatID int64) func(string) {
	return func(text string) {
		msgConfig := tgbotapi.NewMessage(chatID, text)
		bot.Send(msgConfig)
	}
}

// PlayEcho Makee bot send request as response
func PlayEcho(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	tasks := make(map[int64][]Task)
	for update := range updates {
		sendMessage := getSender(bot, update.Message.Chat.ID)
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "add":
				tasks[update.Message.Chat.ID] = append(tasks[update.Message.Chat.ID], Task{strings.Replace(update.Message.Text, "/add ", "", 1)})
			case "show":
				sendMessage(strings.Join(mapTasks(tasks[update.Message.Chat.ID], func(task Task) string {
					return task.Text
				}), "\n"))
			}
			continue
		}
		sendMessage(update.Message.Text)
	}
}
