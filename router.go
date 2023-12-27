package tgbotapi

type Router struct {
	Name     string
	handlers []handler
}

func NewRouter(name string) *Router {
	return &Router{
		Name: name,
	}
}

func (r *Router) HandleCommand(command string, handlerFunc HandlerFunc) {
	r.handlers = append(r.handlers, handler{
		filter: filter{
			Type:  CommandUpdateType,
			Value: command,
		},
		executable: handlerFunc,
	})
}

func (r *Router) HandleMessage(text string, handlerFunc HandlerFunc) {
	r.handlers = append(r.handlers, handler{
		filter: filter{
			Type:  MessageUpdateType,
			Value: text,
		},
		executable: handlerFunc,
	})
}

func (r *Router) HandleCallback(callbackData string, handlerFunc HandlerFunc) {
	r.handlers = append(r.handlers, handler{
		filter: filter{
			Type:  CallbackQueryUpdateType,
			Value: callbackData,
		},
		executable: handlerFunc,
	})
}

func (r *Router) HandleCustomEvent(f filter, handlerFunc HandlerFunc) {
	r.handlers = append(r.handlers, handler{
		filter:     f,
		executable: handlerFunc,
	})
}
