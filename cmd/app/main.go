package main

import (
	"fmt"
	"practice-telegram-bot/internal/app"
	"practice-telegram-bot/internal/app/chains"
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
	botDispatcher := dispatcher.NewRedisDispatcher("redis://localhost:6379")

	fmt.Println(botApi.GetMe())

	helpCommand := dispatcher.Command{Name: "help", Handler: func(event types.Update) { handlers.Hello(event, botApi, &botDispatcher) }}
	addCommand := dispatcher.Command{Name: "add", Handler: func(event types.Update) { handlers.Add(event, botApi, &botDispatcher) }}

	addChain := dispatcher.Chain{
		Name:              "addChain",
		StartChainCommand: addCommand,
		Handlers: []func(event types.Update){
			func(event types.Update) { chains.AddChainFirstPhrase(event, botApi, &botDispatcher) },
			func(event types.Update) { chains.AddChainSecondPhrase(event, botApi, &botDispatcher) },
			func(event types.Update) { chains.AddChainThirdPhrase(event, botApi, &botDispatcher) },
		},
	}

	botDispatcher.AddChains([]dispatcher.Chain{addChain})
	botDispatcher.AddCommands([]dispatcher.Command{helpCommand, addCommand})

	params := types.UpdateRequestParams{Limit: 100}

	app.Run(botApi, &botDispatcher, params)
}
