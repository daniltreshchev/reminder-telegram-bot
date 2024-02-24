package chains

import (
	"practice-telegram-bot/pkg/api"
	"practice-telegram-bot/pkg/dispatcher"
	"practice-telegram-bot/pkg/types"
)

func AddChainFirstPhrase(event types.Update, api api.API, dispatcher *dispatcher.Dispatcher) error {
	message := types.OngoingMessage{ChatId: event.Message.Chat.ID, Text: "add 1 phase"}

	_, err := api.SendMessage(message)

	dispatcher.NextChainStep(event.Message.From)

	return err
}

func AddChainSecondPhrase(event types.Update, api api.API, dispatcher *dispatcher.Dispatcher) error {
	message := types.OngoingMessage{ChatId: event.Message.Chat.ID, Text: "add 2 phase"}

	_, err := api.SendMessage(message)

	dispatcher.NextChainStep(event.Message.From)

	return err
}

func AddChainThirdPhrase(event types.Update, api api.API, dispatcher *dispatcher.Dispatcher) error {
	message := types.OngoingMessage{ChatId: event.Message.Chat.ID, Text: "add 3 phase"}

	_, err := api.SendMessage(message)

	dispatcher.ClearCurrentChain(event.Message.From)

	return err
}
