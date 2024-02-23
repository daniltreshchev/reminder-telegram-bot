package app

import (
	"fmt"
	"practice-telegram-bot/internal/config"
	"practice-telegram-bot/internal/services"
	"practice-telegram-bot/internal/services/handlers"
	"practice-telegram-bot/pkg/botApi"
)

func Run() {
	token, err := config.GetApiToken()

	if err != nil {
		panic(err)
	}

	bot := botApi.BotAPI{Token: token}
	dispatcher := services.Dispatcher{Commands: make(map[string]services.Command)}

	fmt.Println(bot.GetMe())

	helpCommand := services.Command{Name: "help", Handler: func(event botApi.Update, api botApi.BotAPI) { handlers.Hello(event, api) }}

	dispatcher.AddCommand(helpCommand)

	params := botApi.GetUpdatesParams{Limit: 100}
	for {
		updates, err := bot.GetUpdates(params)

		if err != nil {
			fmt.Println(err)
		}

		for _, update := range updates {
			text := update.Message.Text

			command, err := dispatcher.GetCommandByName(text)

			if err != nil {
				continue
			}

			go command.Handler(update, bot)
		}

		if len(updates) != 0 {
			params.Offset = updates[len(updates)-1].UpdateID + 1
		}
	}
}
