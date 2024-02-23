package botApi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func parseApiResponse(resp *http.Response) (APIResponse, error) {
	var apiResponse APIResponse

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return apiResponse, err
	}

	err = json.Unmarshal(body, &apiResponse)

	if err != nil {
		return apiResponse, err
	}

}

func (bot BotAPI) makeApiUrl(method string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.Token, method)
}

func (bot BotAPI) MakeGetApiCall(telegramMethod string) (APIResponse, error) {
	var apiResponse APIResponse
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(
		bot.makeApiUrl(telegramMethod),
	)

	if err != nil {
		return apiResponse, err
	}

	return parseApiResponse(resp)
}

func (bot BotAPI) MakePostApiCall(telegramMethod string, params *bytes.Buffer) (APIResponse, error) {
	var apiResponse APIResponse
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Post(
		bot.makeApiUrl(telegramMethod),
		"application/json",
		params,
	)

	if err != nil {
		return apiResponse, err
	}

	return parseApiResponse(resp)
}
