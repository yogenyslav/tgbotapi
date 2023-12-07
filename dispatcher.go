package tgbotapi

import (
	"context"
	"log"
	"time"
)

type Dispatcher struct {
	b       *Bot
	routers []*Router
	tasks   chan *Ctx
}

func NewDispatcher(b *Bot) *Dispatcher {
	return &Dispatcher{
		b: b,
	}
}

// now supports only text messages
func (d *Dispatcher) handleUpdate(c *Ctx) {
	if c.Message() != nil {
		if c.Message().IsCommand() {
			c.Update().Type = CommandUpdateType
		} else {
			c.Update().Type = MessageUpdateType
		}
	}

	for _, r := range d.routers {
		for _, h := range r.handlers {
			err := h.processUpdate(c)
			if err == nil {
				if d.b.debugMode {
					log.Printf("DEBUG: Handler %s executed update %d of type %d", r.Name, c.Update().Id, c.Update().Type)
				}
				break
			}
			if d.b.debugMode {
				log.Printf("%+v", err)
			}
		}
	}
}

func (d *Dispatcher) AddRouter(r *Router) {
	d.routers = append(d.routers, r)
}

// timeout in seconds
func (d *Dispatcher) StartPolling(timeout int) {
	retryTimeout := 3 * time.Second
	limit := 1

	go func() {
		ctx := context.Background()
		var offset int64 = 0
		for {
			apiResult, err := d.b.getUpdates(offset, limit, retryTimeout)
			if err != nil {
				d.b.updates <- NewCtx(ctx, d.b, apiResult.Ok, GetUpdatesFailed{err.Error(), retryTimeout})
				time.Sleep(retryTimeout)
				continue
			}

			if len(apiResult.Result) == 0 {
				continue
			}

			u := apiResult.Result[0]

			if offset == 0 {
				offset = u.Id + 1
			} else {
				offset++
			}

			d.b.updates <- NewCtx(ctx, d.b, true, nil).Store("update", u)
		}
	}()

	for updateCtx := range d.b.updates {
		if d.b.debugMode {
			if updateCtx.Error() != nil {
				log.Printf("%+v", updateCtx.Error())
			} else {
				log.Printf("DEBUG: update: %+v\n", updateCtx.Update())
			}
		}

		if updateCtx.Error() == nil {
			go d.handleUpdate(updateCtx)
		}
	}
}
