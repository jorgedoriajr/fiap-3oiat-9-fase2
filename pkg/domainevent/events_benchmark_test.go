package domainevent

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkEventDispatcher_RegisterFuzzy(b *testing.B) {
	dispatcher := NewEventDispatcher[TestEventType]()
	rand.Seed(time.Now().UnixNano())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Simula diferentes manipuladores sendo registrados
		handler := &TestEventHandler{ID: rand.Intn(1000)}
		event := &TestEvent[TestEventType]{Name: "BenchmarkEvent", Payload: "TestPayload"}

		dispatcher.Register(event.GetName(), handler)
	}
}
