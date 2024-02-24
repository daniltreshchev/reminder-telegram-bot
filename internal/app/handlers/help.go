package handlers

import (
	"practice-telegram-bot/pkg/api"
	"practice-telegram-bot/pkg/types"
)

func Hello(event types.Update, api api.API) error {
	message := types.OngoingMessage{ChatId: event.Message.Chat.ID, Text: "hello command"}

	_, err := api.SendMessage(message)

	return err
}
