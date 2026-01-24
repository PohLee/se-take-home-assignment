// Package event provides a simple in-memory Pub/Sub event bus implementation.
// It is used to decouple system components and facilitate future Kafka integration.
package event

import (
	"sync"
)

// EventType defines the type of event being published in the system.
type EventType string

const (
	// OrderCreated is emitted when a new order is initially submitted.
	OrderCreated EventType = "ORDER_CREATED"
	// OrderPending is emitted when an order is queued and available for bots.
	OrderPending EventType = "ORDER_PENDING"
	// OrderAssigned is emitted when an order is picked up by a bot worker.
	OrderAssigned EventType = "ORDER_ASSIGNED"
	// OrderCompleted is emitted when a bot successfully finishes processing an order.
	OrderCompleted EventType = "ORDER_COMPLETED"
	// OrderCancelled is emitted when an order processing is interrupted (e.g., bot removed).
	OrderCancelled EventType = "ORDER_CANCELLED"
)

// Event represents a system-wide notification containing a type and payload.
type Event struct {
	Type EventType
	Data interface{}
}

// EventBus handles the subscription and broadcasting of events.
type EventBus struct {
	subscribers map[EventType][]chan Event
	mu          sync.RWMutex
}

// NewEventBus initializes and returns a new thread-safe EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventType][]chan Event),
	}
}

// Subscribe returns a channel that will receive events of the specified type.
// The channel is buffered to prevent producers from blocking on slow consumers.
func (eb *EventBus) Subscribe(eventType EventType) chan Event {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	ch := make(chan Event, 10) // Buffered channel to prevent blocking producers
	eb.subscribers[eventType] = append(eb.subscribers[eventType], ch)
	return ch
}

// Publish broadcasts an event to all active subscribers of the event type.
// If a subscriber's channel is full, the event is skipped for that subscriber
// to prevent system-wide stalls (non-blocking).
func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if subs, ok := eb.subscribers[event.Type]; ok {
		for _, ch := range subs {
			// Non-blocking publish to avoid slow consumers stalling the system
			select {
			case ch <- event:
			default:
				// If channel is full, we skip it or log (depending on requirements)
			}
		}
	}
}

// Unsubscribe removes a channel from the subscriber list.
func (eb *EventBus) Unsubscribe(eventType EventType, ch chan Event) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if subs, ok := eb.subscribers[eventType]; ok {
		for i, sub := range subs {
			if sub == ch {
				close(ch)
				eb.subscribers[eventType] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
	}
}
