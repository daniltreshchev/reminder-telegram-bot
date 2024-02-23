package botApi

import (
	"encoding/json"
)

type BotAPI struct {
	Token string
}

func (bot BotAPI) GetUpdates(getUpdatesParams UpdateRequestParams) ([]Update, error) {
	var updates []Update

	params, err := getUpdatesParams.params()
	if err != nil {
		return updates, err
	}

	apiResponse, err := bot.MakePostApiCall("getUpdates", params)

	if err != nil {
		return updates, err
	}

	err = json.Unmarshal(apiResponse.Result, &updates)
	return updates, err
}

func (bot BotAPI) GetMe() (User, error) {
	var user User

	apiResponse, err := bot.MakeGetApiCall("getMe")

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(apiResponse.Result, &user)

	return user, err
}

func (bot BotAPI) SendMessage(newMessage OngoingMessage) (Message, error) {
	var message Message
	params, err := newMessage.params()

	if err != nil {
		return message, err
	}

	apiResponse, err := bot.MakePostApiCall("sendMessage", params)

	if err != nil {
		return message, err
	}

	err = json.Unmarshal(apiResponse.Result, &message)
	return message, err
}
