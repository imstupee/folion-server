package events

type EventBus struct{}

func NewEventBus() *EventBus {
	return &EventBus{}
}
