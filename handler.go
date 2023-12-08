package tgbotapi

type handler struct {
	filter     Filter
	executable HandlerFunc
}

type HandlerFunc = func(*Ctx) error

func (h *handler) processUpdate(c *Ctx) error {
	if c.Update().Type == h.filter.Type {
		matched := false

		switch c.Update().Type {
		case CommandUpdateType:
			matched = c.Message().Command() == h.filter.Value
		case MessageUpdateType:
			if h.filter.Value == "" {
				matched = true
			} else {
				matched = c.Message().Text == h.filter.Value
			}
		case CallbackUpdateType:
			matched = c.Update().CallbackQuery.Data == h.filter.Value
		}
		if matched {
			return h.executable(c)
		} else {
			return InvalidType{c.Update().Message.Command(), h.filter.Value}
		}
	} else {
		return HandlerMismatch{c.Update().Type, h.filter.Type}
	}
}
