package domainevent

import (
	"errors"
	"log"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher[T any] interface {
	Register(eventName string, handler EventHandler[T]) error
	Dispatch(event Event[T]) error
}

type Dispatcher[T any] struct {
	Handlers map[string][]EventHandler[T]
}

func (d *Dispatcher[T]) Dispatch(event Event[T]) <-chan error {
	errChan := make(chan error, len(d.Handlers[event.GetName()]))
	var wg sync.WaitGroup

	for _, h := range d.Handlers[event.GetName()] {
		wg.Add(1)
		go func(handler EventHandler[T]) {
			defer wg.Done()
			if err := handler.Handle(event); err != nil {
				errChan <- err
			}
		}(h.(EventHandler[T]))
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	return errChan
}

func (d *Dispatcher[T]) registerHandler(handler EventHandler[T]) error {
	eventType := GetEventType(handler)
	if eventType == nil {
		return errors.New("ErrInvalidHandlerType")
	}
	d.Handlers[eventType.String()] = append(d.Handlers[eventType.String()], handler)
	return nil
}

func (d *Dispatcher[T]) RegisterHandlers(handlers ...EventHandler[T]) error {
	var err error
	for _, h := range handlers {
		err = d.registerHandler(h)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dispatcher[T]) Register(eventName string, handlers ...EventHandler[T]) {
	if _, ok := d.Handlers[eventName]; !ok {
		d.Handlers[eventName] = make([]EventHandler[T], 0)
	}

	for _, h := range handlers {

		if containsHandler(d.Handlers[eventName], h) {
			log.Printf("Handler %v already registered for event %v", h, eventName)
			continue
		}
		d.Handlers[eventName] = append(d.Handlers[eventName], h)
	}
}

func containsHandler[T any](handlers []EventHandler[T], handler EventHandler[T]) bool {
	for _, h := range handlers {
		if h == handler {
			return true
		}
	}
	return false
}

func NewEventDispatcher[T any]() *Dispatcher[T] {
	return &Dispatcher[T]{Handlers: make(map[string][]EventHandler[T])}
}
