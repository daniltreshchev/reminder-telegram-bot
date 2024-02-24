package api

import (
	"encoding/json"
	"practice-telegram-bot/pkg/types"
)

type API struct {
	Token string
}

func (api API) GetUpdates(getUpdatesParams types.UpdateRequestParams) ([]types.Update, error) {
	var updates []types.Update

	params, err := getUpdatesParams.Params()
	if err != nil {
		return updates, err
	}

	apiResponse, err := api.MakePostApiCall("getUpdates", params)

	if err != nil {
		return updates, err
	}

	err = json.Unmarshal(apiResponse.Result, &updates)
	return updates, err
}

func (api API) GetMe() (types.User, error) {
	var user types.User

	apiResponse, err := api.MakeGetApiCall("getMe")

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(apiResponse.Result, &user)

	return user, err
}

func (api API) SendMessage(newMessage types.OngoingMessage) (types.Message, error) {
	var message types.Message
	params, err := newMessage.Params()

	if err != nil {
		return message, err
	}

	apiResponse, err := api.MakePostApiCall("sendMessage", params)

	if err != nil {
		return message, err
	}

	err = json.Unmarshal(apiResponse.Result, &message)
	return message, err
}
