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
			userID := update.Message.From.ID
			command, err := botDispatcher.GetCommandByName(text)
			currentChainLink, ok := botDispatcher.CurrentChainLink[userID]

			if err != nil && currentChainLink != -1 && !ok {
				continue
			} else if ok && currentChainLink > -1 {
				currentChain := botDispatcher.Chains[botDispatcher.CurrentChain[userID]]

				currentChain.Handlers[botDispatcher.CurrentChainLink[userID]](update)

				continue
			}

			command.Handler(update)
		}

		if len(updates) != 0 {
			params.Offset = updates[len(updates)-1].UpdateID + 1
		}
	}
}
