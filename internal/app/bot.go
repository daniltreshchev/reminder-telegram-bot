package app

import (
	"fmt"
	"practice-telegram-bot/pkg/api"
	"practice-telegram-bot/pkg/dispatcher"
	"practice-telegram-bot/pkg/types"
)

func Run(botApi api.API, botDispatcher *dispatcher.Dispatcher, params types.UpdateRequestParams) {
	for {
		updates, err := botApi.GetUpdates(params)

		if err != nil {
			fmt.Println(err)
		}

		for _, update := range updates {
			text := update.Message.Text

			command, err := botDispatcher.GetCommandByName(text)

			if err != nil && botDispatcher.CurrentChainLink == -1 {
				continue
			} else if botDispatcher.CurrentChainLink > -1 {
				botDispatcher.CurrentChain.Handlers[botDispatcher.CurrentChainLink](update)
				continue
			}

			command.Handler(update)
		}

		if len(updates) != 0 {
			params.Offset = updates[len(updates)-1].UpdateID + 1
		}
	}
}
