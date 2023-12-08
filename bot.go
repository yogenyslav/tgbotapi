package tgbotapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Bot struct {
	token     string
	apiUrl    string
	updates   chan *Ctx
	debugMode bool
}

func NewBotAPI(token string, debugMode bool) *Bot {
	return &Bot{
		token:     token,
		apiUrl:    fmt.Sprintf("https://api.telegram.org/bot%s", token),
		updates:   make(chan *Ctx),
		debugMode: debugMode,
	}
}

func (b *Bot) getUpdates(offset int64, limit int, timeout time.Duration) (*ApiResultUpdate, error) {
	response, err := http.Get(fmt.Sprintf("%s/getUpdates?offset=%d&timeout=%d&limit=1", b.apiUrl, offset, timeout))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("getUpdates got status code %d", response.StatusCode)
	}

	var data *ApiResultUpdate
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %+v", err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %+v", err)
	}

	if b.debugMode {
		if len(data.Result) > 0 {
			log.Printf("DEBUG: method: getUpdates, response: %s", string(body))
		}
	}

	if !data.Ok {
		return nil, fmt.Errorf("api returned not ok: %s", data.Description)
	}

	return data, nil
}

func (b *Bot) callApi(requestType, method string, requestBody any) (*ApiResult, error) {
	var (
		response *http.Response
		err      error
	)
	if requestType == "GET" {
		response, err = http.Get(fmt.Sprintf("%s/%s", b.apiUrl, method))
	} else if requestType == "POST" {
		rawRequestBody, err := json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %+v", err)
		}
		response, err = http.Post(fmt.Sprintf("%s/%s", b.apiUrl, method), "application/json", bytes.NewReader(rawRequestBody))
	} else {
		return nil, fmt.Errorf("invalid request type: %s", requestType)
	}
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%s got status code %d", method, response.StatusCode)
	}

	var data *ApiResult
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %+v", err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %+v", err)
	}

	if b.debugMode {
		log.Printf("DEBUG: method: %s, response: %s", method, string(body))
	}

	if !data.Ok {
		return nil, fmt.Errorf("api returned not ok: %s", data.Description)
	}

	return data, nil
}
