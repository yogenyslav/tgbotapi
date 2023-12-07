package tgbotapi

type handler struct {
	filter     Filter
	executable HandlerFunc
}

type HandlerFunc = func(*Ctx) error

func (h *handler) processUpdate(c *Ctx) error {
	if c.Update().Type == h.filter.Type {
		return h.executable(c)
	} else {
		return HandlerMismatch{c.Update().Type, h.filter.Type}
	}
}
