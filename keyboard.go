package tgbotapi

type KeyboardSettings struct {
	InlineKeyboard *InlineKeyboardMarkup `json:"inline_keyboard,omitempty"`
	ReplyKeyboard  *ReplyKeyboardMarkup  `json:"reply_keyboard,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	Url                          string        `json:"url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo   `json:"web_app,omitempty"`
	LoginUrl                     *LoginUrl     `json:"login_url,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  string        `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

type ReplyKeyboardMarkup struct {
	Keyboard     [][]*KeyboardButton `json:"keyboard"`
	IsPersistent bool                `json:"is_persistent,omitempty"`
	Resize       bool                `json:"resize_keyboard,omitempty"`
	OneTime      bool                `json:"one_time_keyboard,omitempty"`
	Placeholder  string              `json:"input_field_placeholder,omitempty"`
}

type KeyboardButton struct {
	Text            string                     `json:"text"`
	RequestUser     *KeyboardButtonRequestUser `json:"request_user,omitempty"`
	RequestChat     *KeyboardButtonRequestChat `json:"request_chat,omitempty"`
	RequestContact  bool                       `json:"request_contact,omitempty"`
	RequestLocation bool                       `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType    `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo                `json:"web_app,omitempty"`
}

type KeyboardButtonRequestUser struct {
}

type KeyboardButtonRequestChat struct {
}

type KeyboardButtonPollType struct {
}

type LoginUrl struct {
}

type CallbackGame struct {
}

func NewReplyMarkup(oneTime bool) *KeyboardSettings {
	return NewReplyMarkupFromButtons(nil, oneTime)
}

func NewReplyMarkupFromButtons(buttons [][]*KeyboardButton, oneTime bool) *KeyboardSettings {
	return &KeyboardSettings{
		ReplyKeyboard: &ReplyKeyboardMarkup{
			Keyboard: buttons,
			OneTime:  oneTime,
		},
	}
}

func NewReplyKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{
		Text: text,
	}
}

func (rk *ReplyKeyboardMarkup) AddRow(buttons ...*KeyboardButton) {
	rk.Keyboard = append(rk.Keyboard, buttons)
}
