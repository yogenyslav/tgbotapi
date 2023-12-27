# tgbotapi

A **zero-dependency** Telegram Bot API wrapper in Golang inspired by **[aiogram](https://github.com/aiogram/aiogram)** and **[fiber](https://github.com/gofiber/fiber).**\
**tgbotapi** will help you to write simple Telegram bots (_support of more complex features is comming soon..._) using convenient abstractions.\
Feel the **power** of routing and **easy** custom filtering for Telegram events matching any of your scenarios!

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
    b := tgbotapi.NewBotAPI("token", true /* debugMode */)
    dp := tgbotapi.NewDispatcher(b)

    router := tgbotapi.NewRouter("main")

    // Handle '/start' command
    router.HandleCommand("start", func(c *tgbotapi.Ctx) error {
        err := c.AnswerMessage(fmt.Sprintf("Hello, %s!", c.Message().From.FirstName), nil /* replyMarkup */)
        if err != nil {
            panic(err)
        }

        // Reply keyboard example
        markup := tgbotapi.NewReplyKeyboardMarkup(true /* oneTime */)
        markup.AddRow(tgbotapi.NewReplyKeyboardButton("Button1"), tgbotapi.NewReplyKeyboardButton("Button2"))
        _ = markup

        // Inline keyboard example
        kb := tgbotapi.NewInlineKeyboardMarkup()
        for i := 0; i < 4; i++ {
            kb.AddRow(tgbotapi.NewInlineKeyboardButton(fmt.Sprintf("Button %d", i+1), "testCallbackData"))
        }

        return c.AnswerMessage("This bot is written in Golang!", kb)
    })

    // Handle event with custom filter
    emptyMessageFilter := tgbotapi.NewFilter(tgbotapi.MessageUpdateType, "")
    router.HandleCustomEvent(emptyMessageFilter, func(c *tgbotapi.Ctx) error {
        return c.AnswerMessage(c.Message().Text, nil)
    })

    // Handle inline keyboard presses
    router.HandleCallback("testCallbackData", func(c *tgbotapi.Ctx) error {
        return c.AnswerMessage("Got callbackData "+c.Update().CallbackQuery.Data, nil)
    })

    dp.AddRouter(router)
    dp.StartPolling(60) // Blocks main goroutine
}
```

## Features

- ### Messages

  - [X] Handle messages with specific filter
  - [X] Answer with messages
  - [X] Reply keyboard support
  - [X] Inline keyboard support

- ### Filters

  - [ ] Add 'startsWith' support for filter value
  - [ ] Multiple values matching

- ### Callback Queries

  - [ ] Handle queries with specific data

- ### States

  - [ ] In-memory state storaging
  - [ ] External state storages support
