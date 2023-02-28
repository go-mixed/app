package providers

import (
	"gopkg.in/go-mixed/framework.v1/contracts/event"
	eventfacade "gopkg.in/go-mixed/framework.v1/facades/event"
)

type EventServiceProvider struct {
}

func (receiver *EventServiceProvider) Register() {
	eventfacade.Register(receiver.listen())
}

func (receiver *EventServiceProvider) Boot() {

}

func (receiver *EventServiceProvider) listen() map[event.Event][]event.Listener {
	return map[event.Event][]event.Listener{}
}
