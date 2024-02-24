package handlers

import (
	"practice-telegram-bot/pkg/api"
	"practice-telegram-bot/pkg/dispatcher"
	"practice-telegram-bot/pkg/types"
)

func Add(event types.Update, api api.API, dispatcher *dispatcher.Dispatcher) error {
	ongoingMessage := types.OngoingMessage{ChatId: event.Message.Chat.ID, Text: "hello command"}
	message := event.Message.Text

	_, err := api.SendMessage(ongoingMessage)

	if err != nil {
		return err
	}

	command, _ := dispatcher.GetCommandByName(message)

	dispatcher.StartChain(command)

	return nil
}
