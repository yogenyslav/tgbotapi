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
		filter: Filter{
			Type:  CommandUpdateType,
			Value: command,
		},
		executable: handlerFunc,
	})
}

func (r *Router) HandleMessage(text string, handlerFunc HandlerFunc) {
	r.handlers = append(r.handlers, handler{
		filter: Filter{
			Type:  MessageUpdateType,
			Value: text,
		},
		executable: handlerFunc,
	})
}

func (r *Router) HandleCallback(callbackData any, handlerFunc HandlerFunc) {
	r.handlers = append(r.handlers, handler{
		filter: Filter{
			Type:  MessageUpdateType,
			Value: callbackData,
		},
		executable: handlerFunc,
	})
}
