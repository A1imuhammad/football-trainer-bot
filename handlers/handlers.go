package handlers

import (
	data "botik/internal/data"
	feature "botik/internal/servises"
	"botik/keyboards"
	"botik/models"
	"encoding/json"
	"os"
	"strings"

	"fmt"
	"log"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("Обрабатываем команду /start от: %d", update.Message.Chat.ID)
	feature.AddUser(update.Message.Chat.ID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Привет, футбольный фанат! ⚽

Я твой личный помощник, чтобы помочь тебе стать лучше на поле! 🌟
Этот бот создан для развития твоих навыков, улучшения тактики и вдохновения через интересные факты о футболе. 🎉

Я предложу ежедневные упражнения для отработки техники и дам советы, которые пригодятся в игре. 💪
А ещё поделюсь крутыми историями из мира футбола, чтобы ты всегда был мотивирован! 🚀

Ты сможешь следить за своим прогрессом и видеть, как становишься сильнее. 🌱
Начни прямо сейчас — футбол ждёт! ⚡`)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Что я умею", "function_commands")),
	)

	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func handleGeneralCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.CallbackQuery.Data {
	case "tactic":
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, `Какой совет ты хочешь получить?`)
		msg.ReplyMarkup = keyboards.TacticKeyboard
		bot.Send(msg)

	case "fact":
		facts, err := data.Facts()
		if err != nil {
			log.Println(err)
			return
		}
		reply := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Давай следующий", "fact"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Назад", "function_commands"),
			),
		)
		edit := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			facts[rand.Intn(len(facts))].Fact,
			reply,
		)
		bot.Send(edit)

	case "drill":
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, `Для какой игровой ситуации ты хочешь получить упражнение?`)
		msg.ReplyMarkup = keyboards.DrillKeyboard
		bot.Send(msg)
	case "progress":
		userID := update.CallbackQuery.From.ID
		report := getProgressText(int64(userID))
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, report)
		bot.Send(msg)
	}

}

func handleGeneralCommands2(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "tactic":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Какой совет ты хочешь получить?`)
		msg.ReplyMarkup = keyboards.TacticKeyboard
		bot.Send(msg)

	case "fact":
		facts, err := data.Facts()
		if err != nil {
			log.Println(err)
			return
		}
		reply := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Давай следующий", "fact"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Назад", "function_commands"),
			),
		)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, facts[rand.Intn(len(facts))].Fact)
		msg.ReplyMarkup = reply
		bot.Send(msg)

	case "drill":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Для какой игровой ситуации ты хочешь получить упражнение?`)
		msg.ReplyMarkup = keyboards.DrillKeyboard
		bot.Send(msg)
	case "progress":
		userID := update.Message.From.ID
		report := getProgressText(int64(userID))
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, report)
		bot.Send(msg)
	}
}

func handleTacticCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	tactics, err := data.Tactics()
	if err != nil {
		log.Println(err)
		return
	}

	switch update.CallbackQuery.Data {
	case "attack":
		reply := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Давай следующий", "attack"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Назад", "tactic"),
			),
		)
		edit := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			tactics.Tactics.Attack[rand.Intn(len(tactics.Tactics.Attack))],
			reply,
		)
		bot.Send(edit)

	case "defense":
		reply := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Давай следующий", "defense"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Назад", "tactic"),
			),
		)
		edit := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			tactics.Tactics.Defense[rand.Intn(len(tactics.Tactics.Defense))],
			reply,
		)
		bot.Send(edit)

	case "position":
		reply := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Давай следующий", "position"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Назад", "tactic"),
			),
		)
		edit := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			tactics.Tactics.Positioning[rand.Intn(len(tactics.Tactics.Positioning))],
			reply,
		)
		bot.Send(edit)

	}
}

func handleDrillCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.CallbackQuery.Data {
	case "attacker":
		edit := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			`Для какой позиции ты хочешь получить упражнение?`,
			keyboards.ForAttackerKeyboard,
		)
		bot.Send(edit)
	case "defend":
		edit := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			`Для какой позиции ты хочешь получить упражнение?`,
			keyboards.ForDefendceKeyboard,
		)
		bot.Send(edit)
	}
}

const progressFile = "user_progress.json"

var userProgress = make(map[int64]map[string]int)

func handleTrainingCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	att, def, err := data.Training()
	if err != nil {
		log.Println(err)
		return
	}

	exerciseMap := map[string][]models.Exercise{
		"general":              att.General,
		"forward":              att.Forward,
		"winger":               att.Winger,
		"attackingMidfielder":  att.AttackingMidfielder,
		"centralMidfielder":    att.CentralMidfielder,
		"defensiveMidfielder":  att.DefensiveMidfielder,
		"laterale":             att.Laterale,
		"defender":             att.Defender,
		"goalkeeper":           att.Goalkeeper,
		"general2":             def.General,
		"forward2":             def.Forward,
		"attackingMidfielder2": def.AttackingMidfielder,
		"centralMidfielder2":   def.CentralMidfielder,
		"defensiveMidfielder2": def.DefensiveMidfielder,
		"laterale2":            def.Laterale,
		"defender2":            def.Defender,
		"goalkeeper2":          def.Goalkeeper,
	}
	if strings.HasPrefix(update.CallbackQuery.Data, "done_") {
		positionKey := strings.TrimPrefix(update.CallbackQuery.Data, "done_")
		positionKey = strings.TrimSuffix(positionKey, "2") // на всякий случай
		if userProgress[int64(update.CallbackQuery.From.ID)] == nil {
			userProgress[int64(update.CallbackQuery.From.ID)] = make(map[string]int)
		}
		userProgress[int64(update.CallbackQuery.From.ID)][positionKey]++

		if err := saveProgress(); err != nil {
			log.Printf("Ошибка при сохранении прогресса: %v", err)
		}

		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Отлично! Прогресс сохранён ✅")
		bot.Send(msg)
		return
	}
	if exercises, ok := exerciseMap[update.CallbackQuery.Data]; ok {
		if len(exercises) == 0 {
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Нет доступных упражнений.")
			bot.Send(msg)
			return
		}

		var backCallback string
		if strings.Contains(update.CallbackQuery.Data, "2") {
			backCallback = "defend"
		} else {
			backCallback = "attacker"
		}

		reply := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("✅Выполнено", "done_"+update.CallbackQuery.Data),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Давай другое", update.CallbackQuery.Data),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Назад", backCallback),
			),
		)

		ex := exercises[rand.Intn(len(exercises))]
		text := fmt.Sprintf("Упражнение: %s\nОписание: %s\nФокус: %s", ex.Name, ex.Description, ex.Focus)
		edit := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			text,
			reply,
		)
		bot.Send(edit)
	} else {
		log.Println("Неизвестная команда:", update.CallbackQuery.Data)
	}
}

func getProgressText(userID int64) string {
	stats, ok := userProgress[userID]
	if !ok || len(stats) == 0 {
		return `😔 Пока нет прогресса, но это отличный повод начать! 
⚽ Выбери упражнение и стань лучше уже сегодня!`
	}

	positionTranslations := map[string]string{
		"general":             "Общие",
		"forward":             "Нападающий",
		"winger":              "Вингер",
		"attackingMidfielder": "Атакующий полузащитник",
		"centralMidfielder":   "Центральный полузащитник",
		"defensiveMidfielder": "Опорный полузащитник",
		"laterale":            "Фланговый защитник",
		"defender":            "Центральный защитник",
		"goalkeeper":          "Вратарь",
	}

	totalExercises := 0
	for _, count := range stats {
		totalExercises += count
	}
	var sb strings.Builder
	sb.WriteString("📊 Ваша статистика выполненных упражнений по позициям:\n\n")
	for position, count := range stats {
		// Получаем переведённое название позиции, если оно есть, иначе используем исходное
		positionName := positionTranslations[position]
		if positionName == "" {
			positionName = position // На случай, если позиция не найдена в словаре
		}
		sb.WriteString(fmt.Sprintf("- %s: %d упражнений\n", positionName, count))
	}
	sb.WriteString(fmt.Sprintf("\n🔥 Всего выполнено: %d %s\n", totalExercises, getExerciseWord(totalExercises)))
	return sb.String()
}

func getExerciseWord(count int) string {
	if count%10 == 1 && count%100 != 11 {
		return "упражнение"
	} else if count%10 >= 2 && count%10 <= 4 && (count%100 < 10 || count%100 > 20) {
		return "упражнения"
	}
	return "упражнений"
}

func loadProgress() error {
	data, err := os.ReadFile(progressFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Файл прогресса не существует, создаётся новый")
			return nil
		}
		return fmt.Errorf("ошибка чтения файла %s: %v", progressFile, err)
	}

	if err := json.Unmarshal(data, &userProgress); err != nil {
		return fmt.Errorf("ошибка декодирования JSON: %v", err)
	}
	log.Println("Прогресс успешно загружен из", progressFile)
	return nil
}

func saveProgress() error {
	data, err := json.MarshalIndent(userProgress, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка кодирования JSON: %v", err)
	}

	if err := os.WriteFile(progressFile, data, 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл %s: %v", progressFile, err)
	}
	log.Println("Прогресс успешно сохранён в", progressFile)
	return nil
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	if update.CallbackQuery != nil {
		bot.Send(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))

		if update.CallbackQuery.Data == "function_commands" {
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,
				`Я помогу тебе стать лучше на поле с помощью ежедневных упражнений, тактических советов и отображением прогресса.
А ещё я буду делиться с тобой интересными фактами из мира футбола! 📚`,
			)

			msg.ReplyMarkup = keyboards.FunctionalKeyboard
			bot.Send(msg)
		}
		handleTrainingCommands(bot, update)
		handleTacticCommands(bot, update)
		handleDrillCommands(bot, update)
		handleGeneralCommands(bot, update)

	}
	if update.Message == nil {
		return
	}
	handleGeneralCommands2(bot, update)
	if update.Message.Text == "/start" {
		handleStart(bot, update)
	}

}

func ListenUpdates(bot *tgbotapi.BotAPI) {
	if err := loadProgress(); err != nil {
		log.Printf("Ошибка при загрузке прогресса: %v", err)
	}

	offset := lastUpdatesID(bot)

	updateConfig := tgbotapi.NewUpdate(offset)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		handleUpdate(bot, update)
	}
}

func lastUpdatesID(bot *tgbotapi.BotAPI) int {
	updates, err := bot.GetUpdates(tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   1000,
		Timeout: 0,
	})
	if err != nil {
		log.Printf("Ошибка при получении обновлений: %v", err)
		return 0
	}
	if len(updates) > 0 {
		return updates[len(updates)-1].UpdateID + 1
	}
	return 0
}
