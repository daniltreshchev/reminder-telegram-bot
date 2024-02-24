package main

import (
	"fmt"
	"practice-telegram-bot/internal/app"
	"practice-telegram-bot/internal/app/handlers"
	"practice-telegram-bot/internal/config"
	"practice-telegram-bot/pkg/api"
	"practice-telegram-bot/pkg/dispatcher"
	"practice-telegram-bot/pkg/types"
)

func main() {
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

	app.Run(botApi, botDispatcher, params)
}
