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

			if err != nil && botDispatcher.CurrentChainLink[userID] == 0 {
				continue
			} else if len(botDispatcher.CurrentChainLink) > 0 && botDispatcher.CurrentChainLink[userID] > 0 {
				// botDispatcher.Chains

				currentChainName = botDispatcher.CurrentChain[userID]

				botDispatcher.Chains[currentChainName][0].Handlers
				continue
			}

			command.Handler(update)
		}

		if len(updates) != 0 {
			params.Offset = updates[len(updates)-1].UpdateID + 1
		}
	}
}
