package tgbot

import (
	"log"

	"github.com/Ksenofontovas/solution_cup/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	botAPI  *tgbotapi.BotAPI
	timeout int
	service *service.Service
}

func NewTgBot(api string, debug bool, timeout int, service *service.Service) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		return nil, err
	}
	bot.Debug = debug

	return &TgBot{
		botAPI:  bot,
		timeout: timeout,
		service: service,
	}, nil
}

func (tg *TgBot) GetUpdates() {

	log.Printf("Authorized on account %s", tg.botAPI.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.botAPI.GetUpdatesChan(u)
	// Loop through each update.
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			switch update.Message.Command() {
			case "create":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.Text = `Для создания работы пришли сообщение формата:
							Новая работа
							Начало: 15.04.2023
							Конец: 15.04.2023
							Длительность: 01:10
							Зона: 1
							Срочная: нет
							Критическая: да
							`
				tg.botAPI.Send(msg)
			}
			if tg.service.ValidateTask(update.Message.Text) {
				// /strings.Contains(update.Message.Text, "Новая работа")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Работа создана")
				tg.botAPI.Send(msg)
			}
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID
			// msg.ReplyMarkup = mainInlineKeyboard

			// tg.botAPI.Send(msg)
		}
		if update.CallbackQuery != nil {
			//var editMessage tgbotapi.EditMessageTextConfig
			log.Println(update.CallbackQuery.Data)
		}
	}
}
