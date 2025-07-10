package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var FunctionalKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Случайное упражнение", "drill"),
		tgbotapi.NewInlineKeyboardButtonData("Дать совет по тактике", "tactic"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Показать прогресс", "progress"),
		tgbotapi.NewInlineKeyboardButtonData("Футбольный факт", "fact"),
	),
)

var TacticKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("По атаке", "attack"),
		tgbotapi.NewInlineKeyboardButtonData("По защите", "defense"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("По позиционированию", "position"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "function_commands"),
	),
)

var DrillKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Атака", "attacker"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Oборона", "defend"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "function_commands"),
	),
)

var ForAttackerKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Общее для ситуации", "general"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ЦН", "forward"),
		tgbotapi.NewInlineKeyboardButtonData("ПрН/ЛН", "winger"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("АПЗ", "attackingMidfielder"),
		tgbotapi.NewInlineKeyboardButtonData("ЦПЗ", "centralMidfielder"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ОПЗ", "defensiveMidfielder"),
		tgbotapi.NewInlineKeyboardButtonData("ПрЗ/ЛЗ", "laterale"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ЦЗ", "defender"),
		tgbotapi.NewInlineKeyboardButtonData("ВРТ", "goalkeeper"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "drill"),
	),
)

var ForDefendceKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Общее для ситуации", "general2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ЦН", "forward2"),
		tgbotapi.NewInlineKeyboardButtonData("ПрН/ЛН", "winger2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("АПЗ", "attackingMidfielder2"),
		tgbotapi.NewInlineKeyboardButtonData("ЦПЗ", "centralMidfielder2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ОПЗ", "defensiveMidfielder2"),
		tgbotapi.NewInlineKeyboardButtonData("ПрЗ/ЛЗ", "laterale2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ЦЗ", "defender2"),
		tgbotapi.NewInlineKeyboardButtonData("ВРТ", "goalkeeper2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "drill"),
	),
)
