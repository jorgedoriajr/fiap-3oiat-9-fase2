package domainevent

import (
	"context"
	"reflect"
)

type EventHandler[T any] interface {
	Handle(ctx context.Context, event Event[T]) error
}

func GetEventType[T any](handler EventHandler[T]) reflect.Type {
	handlerType := reflect.TypeOf(handler)
	if handlerType.Kind() != reflect.Ptr || handlerType.Elem().Kind() != reflect.Struct {
		return nil
	}
	handleMethod, ok := handlerType.MethodByName("Handle")
	if !ok {
		return nil
	}
	return handleMethod.Type.In(1)
}
