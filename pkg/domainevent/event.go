package domainevent

import "time"

type Event[T any] interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() T
}

type EventModel[T any] struct {
	Name     string
	DateTime time.Time
	Payload  T
}

func (e EventModel[T]) GetName() string {
	return e.Name
}

func (e EventModel[T]) GetDateTime() time.Time {
	return time.Now()
}

func (e EventModel[T]) GetPayload() T {
	return e.Payload
}
