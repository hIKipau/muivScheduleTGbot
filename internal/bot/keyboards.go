package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var DefaultKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Сегодня"),
		tgbotapi.NewKeyboardButton("Выбрать день"),
	),
)
var ChooseKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Понедельник"),
		tgbotapi.NewKeyboardButton("Вторник"),
		tgbotapi.NewKeyboardButton("Среда"),
		tgbotapi.NewKeyboardButton("Четверг"),
		tgbotapi.NewKeyboardButton("Пятница"),
		tgbotapi.NewKeyboardButton("Суббота"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться"),
	),
)
