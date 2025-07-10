package main

import (
	"botik/handlers"
	feature "botik/internal/servises"
	"botik/models"

	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	feature.LoadUsers()
	var tkn models.Token

	if err := cleanenv.ReadConfig("../config/.env", &tkn); err != nil {
		log.Fatalf("Ошибка чтения конфигурации: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(tkn.TGtoken)
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	bot.Debug = true



	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	feature.InitScheduler(bot, loc)
	handlers.ListenUpdates(bot)
}
