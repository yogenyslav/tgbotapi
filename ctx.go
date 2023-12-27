package tgbotapi

import (
	"context"
)

type ApiResultUpdate struct {
	Ok          bool      `json:"ok"`
	Description string    `json:"description,omitempty"`
	Result      []*Update `json:"result"`
}

type ApiResult struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
	Result      any    `json:"result"`
}

type Ctx struct {
	context context.Context
	bot     *Bot
}

func NewCtx(ctx context.Context, bot *Bot, ok bool, err error) *Ctx {
	ctx = context.WithValue(ctx, "ok", ok)
	ctx = context.WithValue(ctx, "error", err)

	return &Ctx{
		context: ctx,
		bot:     bot,
	}
}

func (c *Ctx) setUpdateType(t UpdateType) {
	c.Update().Type = t
}

func (c *Ctx) AnswerMessage(text string, replyMarkup KeyboardMarkup) error {
	var chatId int64

	if c.Message() != nil {
		chatId = c.Message().From.Id
	} else if c.CallbackQuery() != nil {
		chatId = c.CallbackQuery().From.Id
	}

	options := SendMessageOptions{
		ChatId:      chatId,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}

	return c.AnswerMessageWithOptions(options)
}

func (c *Ctx) AnswerMessageWithOptions(options SendMessageOptions) error {
	apiResult, err := c.Bot().callApi("POST", "sendMessage", options)
	if err != nil {
		c.Store("error", ApiError{err.Error()})
		return err
	}

	if !apiResult.Ok {
		err = ApiError{apiResult.Description}
		c.Store("error", err)
		return err
	}

	return nil
}

func (c *Ctx) Store(key any, value any) *Ctx {
	c.context = context.WithValue(c.context, key, value)

	return c
}

func (c *Ctx) Context() context.Context {
	return c.context
}

func (c *Ctx) Bot() *Bot {
	return c.bot
}

func (c *Ctx) IsOk() bool {
	isOk, ok := c.context.Value("ok").(bool)
	if !ok {
		return false
	}
	return isOk
}

func (c *Ctx) Error() error {
	err, ok := c.context.Value("error").(error)
	if !ok {
		return nil
	}
	return err
}

func (c *Ctx) Update() *Update {
	update, ok := c.context.Value("update").(*Update)
	if !ok {
		return nil
	}
	return update
}

func (c *Ctx) Message() *Message {
	return c.Update().Message
}

func (c *Ctx) CallbackQuery() *CallbackQuery {
	return c.Update().CallbackQuery
}
