# tgbotapi

**A Telegram Bot API wrapper in Golang inspired by aiogram and fiber.**

## Installation
```bash
go get github.com/yogenyslav/tgbotapi
```

## Quick Start
```go
package main

import (
	"fmt"

	"github.com/yogenyslav/tgbotapi"
)

func main() {
	debugMode := true

	b := tgbotapi.NewBotAPI("token", debugMode)
	dp := tgbotapi.NewDispatcher(b)

	router := tgbotapi.NewRouter("main")

	oneTime := true
	markup := tgbotapi.NewReplyMarkup(oneTime)
	markup.ReplyKeyboard.AddRow(tgbotapi.NewReplyKeyboardButton("Button 1"), tgbotapi.NewReplyKeyboardButton("Button 2"))

	router.HandleCommand("start", func(c *tgbotapi.Ctx) error {
		err := c.AnswerMessage(fmt.Sprintf("Hello, %s!", c.Message().From.FirstName), nil /* replyMarkup */)
		if err != nil {
			panic(err)
		}
		return c.AnswerMessage("This bot is written in Golang!", markup)
	})

	// filter value "" matches every message
	router.HandleMessage("", func(c *tgbotapi.Ctx) error {
		return c.AnswerMessage(c.Message().Text, nil)
	})

	dp.AddRouter(router)
	dp.StartPolling(60)
}
```