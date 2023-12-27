package tgbotapi

type KeyboardMarkup interface {
	Buttons() [][]string
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

func (kb *InlineKeyboardMarkup) Buttons() [][]string {
	inlineKeyboard := make([][]string, 0, len(kb.InlineKeyboard))
	for _, row := range kb.InlineKeyboard {
		buff := make([]string, 0, len(row))
		for _, button := range row {
			buff = append(buff, button.Text)
		}
		inlineKeyboard = append(inlineKeyboard, buff)
	}
	return inlineKeyboard
}

func (kb *InlineKeyboardMarkup) AddRow(buttons ...*InlineKeyboardButton) {
	kb.InlineKeyboard = append(kb.InlineKeyboard, buttons)
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
	ReplyKeyboard [][]*ReplyKeyboardButton `json:"keyboard"`
	IsPersistent  bool                     `json:"is_persistent,omitempty"`
	Resize        bool                     `json:"resize_keyboard,omitempty"`
	OneTime       bool                     `json:"one_time_keyboard,omitempty"`
	Placeholder   string                   `json:"input_field_placeholder,omitempty"`
}

func (kb *ReplyKeyboardMarkup) Buttons() [][]string {
	replyKeyboard := make([][]string, 0, len(kb.ReplyKeyboard))
	for _, row := range kb.ReplyKeyboard {
		buff := make([]string, 0, len(row))
		for _, button := range row {
			buff = append(buff, button.Text)
		}
		replyKeyboard = append(replyKeyboard, buff)
	}
	return replyKeyboard
}

func (kb *ReplyKeyboardMarkup) AddRow(buttons ...*ReplyKeyboardButton) {
	kb.ReplyKeyboard = append(kb.ReplyKeyboard, buttons)
}

type ReplyKeyboardButton struct {
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

func NewReplyKeyboardMarkup(oneTime bool) *ReplyKeyboardMarkup {
	return NewReplyKeyboardMarkupFromButtons(nil, oneTime)
}

func NewReplyKeyboardMarkupFromButtons(buttons [][]*ReplyKeyboardButton, oneTime bool) *ReplyKeyboardMarkup {
	return &ReplyKeyboardMarkup{
		ReplyKeyboard: buttons,
		OneTime:       oneTime,
	}
}

func NewReplyKeyboardButton(text string) *ReplyKeyboardButton {
	return &ReplyKeyboardButton{
		Text: text,
	}
}

func NewInlineKeyboardMarkup() *InlineKeyboardMarkup {
	return NewInlineKeyboardMarkupFromButtons(nil)
}

func NewInlineKeyboardMarkupFromButtons(buttons [][]*InlineKeyboardButton) *InlineKeyboardMarkup {
	return &InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}

func NewInlineKeyboardButton(text, callbackData string) *InlineKeyboardButton {
	return &InlineKeyboardButton{
		Text:         text,
		CallbackData: callbackData,
	}
}
