package botApi

import (
	"bytes"
	"encoding/json"
)

type User struct {
	ID    int  `json:"id"`
	IsBot bool `json:"is_bot"`
}

type APIResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Message struct {
	ID       int    `json:"message_id"`
	From     User   `json:"from"`
	Chat     Chat   `json:"chat"`
	Date     int    `json:"date"`
	Text     string `json:"text"`
	Entities []struct {
		Offset int    `json:"offset"`
		Length int    `json:"length"`
		Type   string `json:"type"`
	} `json:"entities"`
}

type OngoingMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

func (message OngoingMessage) params() (*bytes.Buffer, error) {
	byteParams, err := json.Marshal(message)

	return bytes.NewBuffer(byteParams), err

}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type UpdateRequestParams struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	// Timeout int `json:"timeout"`
}

func (params UpdateRequestParams) params() (*bytes.Buffer, error) {
	byteParams, err := json.Marshal(params)

	return bytes.NewBuffer(byteParams), err
}
