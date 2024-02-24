package app

import (
	"fmt"
	"practice-telegram-bot/internal/app/handlers"
	"practice-telegram-bot/internal/config"
	"practice-telegram-bot/pkg/api"
	"practice-telegram-bot/pkg/dispatcher"
	"practice-telegram-bot/pkg/types"
)

func Run() {
	token, err := config.GetApiToken()

	if err != nil {
		panic(err)
	}

	botApi := api.API{Token: token}
	botDispatcher := dispatcher.Dispatcher{Commands: make(map[string]dispatcher.Command)}

	fmt.Println(botApi.GetMe())

	helpCommand := dispatcher.Command{Name: "help", Handler: func(event types.Update, api api.API) { handlers.Hello(event, api) }}

	botDispatcher.AddCommand(helpCommand)

	params := types.UpdateRequestParams{Limit: 100}
	for {
		updates, err := botApi.GetUpdates(params)

		if err != nil {
			fmt.Println(err)
		}

		for _, update := range updates {
			text := update.Message.Text

			command, err := botDispatcher.GetCommandByName(text)

			if err != nil {
				continue
			}

			go command.Handler(update, botApi)
		}

		if len(updates) != 0 {
			params.Offset = updates[len(updates)-1].UpdateID + 1
		}
	}
}
