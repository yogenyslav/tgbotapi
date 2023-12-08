package tgbotapi

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageId string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

type InlineQuery struct {
}

type ChosenInlineResult struct {
}

type ShippingQuery struct {
}

type PreCheckoutQuery struct {
}
