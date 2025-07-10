package servises

import (
	data "botik/internal/data"
	"botik/models"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)



const usersFile = "../users.json"

func InitScheduler(bot *tgbotapi.BotAPI, loc *time.Location) {
	c := cron.New(cron.WithLocation(loc))
	c.AddFunc("0 12 * * *", func() { sendDailyFact(bot) })
	c.Start()

	log.Println("Планировщик запущен.")
}

func LoadUsers() {
	data, err := os.ReadFile(usersFile)
	if err != nil {
		if os.IsNotExist(err) {
			models.UsersID = []int64{}
			log.Println("Файл users.json не найден, создаём новый.")
			return
		}
		log.Fatalf("Ошибка чтения users.json: %v", err)
	}

	log.Printf("Загружаем users.json: %s\n", data)

	if err := json.Unmarshal(data, &models.UsersID); err != nil {
		log.Fatalf("Ошибка парсинга users.json: %v", err)
	}
	log.Printf("Загружено пользователей: %d\n", len(models.UsersID))
}

func saveUsers() {
	data, err := json.MarshalIndent(models.UsersID, "", "  ")
	if err != nil {
		log.Printf("Ошибка сериализации users: %v", err)
		return
	}

	if err := os.WriteFile(usersFile, data, 0644); err != nil {
		log.Printf("Ошибка записи users.json: %v", err)
	}
}

func AddUser(chatID int64) {
	for _, id := range models.UsersID {
		if id == chatID {
			log.Println("Пользователь уже есть:", chatID)
			return
		}
	}
	log.Println("Добавляю нового пользователя:", chatID)
	models.UsersID = append(models.UsersID, chatID)
	saveUsers()
}

func sendDailyFact(bot *tgbotapi.BotAPI) {
	facts, err := data.Facts()
	if err != nil {
		log.Println(err)
		return
	}

	fact := facts[rand.Intn(len(facts))].Fact
	text := fmt.Sprintf("Интересный факт: \n%s", fact)

	for _, chatID := range models.UsersID {
		msg := tgbotapi.NewMessage(chatID, text)
		bot.Send(msg)
	}

}
