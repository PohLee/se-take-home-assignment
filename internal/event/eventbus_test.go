package event

import (
	"sync"
	"testing"
	"time"
)

func TestNewEventBus(t *testing.T) {
	eb := NewEventBus()
	if eb == nil {
		t.Fatal("NewEventBus returned nil")
	}
	if eb.subscribers == nil {
		t.Error("subscribers map not initialized")
	}
}

func TestSubscribe(t *testing.T) {
	eb := NewEventBus()
	ch := eb.Subscribe(OrderCreated)

	eb.mu.RLock()
	defer eb.mu.RUnlock()

	subs, ok := eb.subscribers[OrderCreated]
	if !ok {
		t.Error("OrderCreated type not found in subscribers")
	}
	if len(subs) != 1 {
		t.Errorf("expected 1 subscriber, got %d", len(subs))
	}
	if subs[0] != ch {
		t.Error("subscriber channel mismatch")
	}
}

func TestPublish(t *testing.T) {
	eb := NewEventBus()
	ch1 := eb.Subscribe(OrderCreated)
	ch2 := eb.Subscribe(OrderCreated)
	ch3 := eb.Subscribe(OrderPending)

	data := "test order"
	eb.Publish(Event{Type: OrderCreated, Data: data})

	// Check delivery to OrderCreated subscribers
	select {
	case ev := <-ch1:
		if ev.Data != data {
			t.Errorf("ch1: expected %v, got %v", data, ev.Data)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("ch1: timeout waiting for event")
	}

	select {
	case ev := <-ch2:
		if ev.Data != data {
			t.Errorf("ch2: expected %v, got %v", data, ev.Data)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("ch2: timeout waiting for event")
	}

	// Check OrderPending subscriber received nothing
	select {
	case ev := <-ch3:
		t.Errorf("ch3: received unexpected event %v", ev)
	case <-time.After(10 * time.Millisecond):
		// Expected timeout
	}
}

func TestPublishNonBlocking(t *testing.T) {
	eb := NewEventBus()
	// Channel has buffer of 10 (hardcoded in Subscribe)
	ch := eb.Subscribe(OrderCreated)

	// Fill the buffer
	for i := 0; i < 10; i++ {
		eb.Publish(Event{Type: OrderCreated, Data: i})
	}

	// This 11th call should be skipped because buffer is full (non-blocking)
	eb.Publish(Event{Type: OrderCreated, Data: "skipped"})

	// Drain the buffer and verify the first 10
	for i := 0; i < 10; i++ {
		ev := <-ch
		if ev.Data != i {
			t.Errorf("expected %d, got %v", i, ev.Data)
		}
	}

	// Verify "skipped" was not delivered
	select {
	case ev := <-ch:
		if ev.Data == "skipped" {
			t.Error("should have skipped the overflow event")
		}
	case <-time.After(10 * time.Millisecond):
		// Expected timeout
	}
}

func TestUnsubscribe(t *testing.T) {
	eb := NewEventBus()
	ch := eb.Subscribe(OrderCreated)

	eb.Unsubscribe(OrderCreated, ch)

	eb.mu.RLock()
	subs := eb.subscribers[OrderCreated]
	eb.mu.RUnlock()

	if len(subs) != 0 {
		t.Errorf("expected 0 subscribers, got %d", len(subs))
	}

	// Verify channel is closed
	_, ok := <-ch
	if ok {
		t.Error("channel should be closed after Unsubscribe")
	}

	// Verify multiple unsubscribes don't panic
	eb.Unsubscribe(OrderCreated, ch)
}

func TestEventBusConcurrency(t *testing.T) {
	eb := NewEventBus()
	const numGoroutines = 50
	const numEvents = 100
	var wg sync.WaitGroup

	// Multiple subscribers
	chs := make([]chan Event, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		chs[i] = eb.Subscribe(OrderCreated)
	}

	// Subscribers draining channels
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			count := 0
			for ev := range chs[idx] {
				if ev.Type == OrderCreated {
					count++
				}
			}
		}(i)
	}

	// Multiple publishers
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numEvents; j++ {
				eb.Publish(Event{Type: OrderCreated, Data: j})
			}
		}()
	}

	// Unsubscribe some while publishing
	for i := 0; i < numGoroutines/2; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			eb.Unsubscribe(OrderCreated, chs[idx*2])
		}(i)
	}

	// Wait for publishers to finish
	time.Sleep(100 * time.Millisecond)

	// Unsubscribe/close remaining channels to stop subscribers
	for i := 0; i < numGoroutines; i++ {
		eb.Unsubscribe(OrderCreated, chs[i])
	}

	wg.Wait()
}
