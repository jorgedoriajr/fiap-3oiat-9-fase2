package domainevent

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestEventType string

type TestEvent[T any] struct {
	Name    string
	Payload T
}

func (t *TestEvent[T]) GetName() string {
	return t.Name
}

func (t *TestEvent[T]) GetDateTime() time.Time {
	return time.Now()
}

func (t *TestEvent[T]) GetPayload() T {
	return t.Payload
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event Event[TestEventType]) error {
	fmt.Println(event.GetPayload())
	return nil
}

type EventDispatcherTestSuite struct {
	suite.Suite
	eventDispatcher EventDispatcher[TestEventType]
	dispatcher      *Dispatcher[TestEventType]
	handler         EventHandler[TestEventType]
	handler2        EventHandler[TestEventType]
	handler3        EventHandler[TestEventType]
	event           Event[TestEventType]
	event2          Event[TestEventType]
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.dispatcher = NewEventDispatcher[TestEventType]()
	suite.handler = &TestEventHandler{ID: 1}
	suite.handler2 = &TestEventHandler{ID: 2}
	suite.handler3 = &TestEventHandler{ID: 3}
	suite.event = &TestEvent[TestEventType]{Name: "Test !", Payload: "Test1"}
	suite.event2 = &TestEvent[TestEventType]{Name: "Test 2", Payload: "Test2"}
}

func TestEventDispatcherSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (suite *EventDispatcherTestSuite) TestRegisterSingleHandler() {
	suite.dispatcher.Register(suite.event.GetName(), suite.handler)

	// Verifica se o manipulador foi registrado corretamente
	suite.assertHandlersCount(suite.event.GetName(), 1)
	suite.assertHandlerRegistered(suite.event.GetName(), suite.handler)
}

func (suite *EventDispatcherTestSuite) TestRegisterMultipleHandlers() {
	// Registrar vários handles para o mesmo evento
	handlers := []EventHandler[TestEventType]{suite.handler, suite.handler2, suite.handler3}
	for _, handler := range handlers {
		suite.dispatcher.Register(suite.event.GetName(), handler)
	}

	// Verifica se todos os manipuladores foram registrados corretamente
	suite.assertHandlersCount(suite.event.GetName(), 3)
	suite.assertHandlerRegistered(suite.event.GetName(), suite.handler)
	suite.assertHandlerRegistered(suite.event.GetName(), suite.handler2)
	suite.assertHandlerRegistered(suite.event.GetName(), suite.handler3)
}

func (suite *EventDispatcherTestSuite) TestRegisterHandlerForDifferentEvent() {
	suite.dispatcher.Register(suite.event2.GetName(), suite.handler)

	// Verifica se o handler foi registrado corretamente para o evento2
	suite.assertHandlersCount(suite.event2.GetName(), 1)
	suite.assertHandlerRegistered(suite.event2.GetName(), suite.handler)
}

func (suite *EventDispatcherTestSuite) assertHandlersCount(eventName string, expectedCount int) {
	suite.Len(suite.dispatcher.Handlers[eventName], expectedCount)
}

func (suite *EventDispatcherTestSuite) assertHandlerRegistered(eventName string, handler EventHandler[TestEventType]) {
	suite.Contains(suite.dispatcher.Handlers[eventName], handler)
}
