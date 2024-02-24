package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"practice-telegram-bot/pkg/types"
)

func parseApiResponse(resp *http.Response) (types.APIResponse, error) {
	var apiResponse types.APIResponse

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return apiResponse, err
	}

	err = json.Unmarshal(body, &apiResponse)

	return apiResponse, err

}

func (api API) makeApiUrl(method string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", api.Token, method)
}

func (api API) MakeGetApiCall(telegramMethod string) (types.APIResponse, error) {
	var apiResponse types.APIResponse
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(
		api.makeApiUrl(telegramMethod),
	)

	if err != nil {
		return apiResponse, err
	}

	return parseApiResponse(resp)
}

func (api API) MakePostApiCall(telegramMethod string, params *bytes.Buffer) (types.APIResponse, error) {
	var apiResponse types.APIResponse
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Post(
		api.makeApiUrl(telegramMethod),
		"application/json",
		params,
	)

	if err != nil {
		return apiResponse, err
	}

	return parseApiResponse(resp)
}
