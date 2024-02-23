package handlers

import "practice-telegram-bot/pkg/botApi"

func Hello(event botApi.Update, api botApi.BotAPI) error {
	message := botApi.OngoingMessage{ChatId: event.Message.Chat.ID, Text: "hello command"}

	_, err := api.SendMessage(message)

	return err
}
