package order

import (
	"testing"
	"time"
)

func TestOrderPriority(t *testing.T) {
	q := NewQueue()

	// Add Normal Order
	o1 := &Order{ID: 1, Type: OrderTypeNormal, Priority: OrderPriorityNormal, CreatedAt: time.Now()}
	q.Push(o1)

	// Add VIP Order later
	time.Sleep(10 * time.Millisecond)
	o2 := &Order{ID: 2, Type: OrderTypeVIP, Priority: OrderPriorityVIP, CreatedAt: time.Now()}
	q.Push(o2)

	// VIP should come out first despite being added later
	first := q.Pop()
	if first.ID != 2 {
		t.Errorf("Expected VIP order (ID 2) to be first, got ID %d", first.ID)
	}

	second := q.Pop()
	if second.ID != 1 {
		t.Errorf("Expected Normal order (ID 1) to be second, got ID %d", second.ID)
	}
}

func TestSamePriorityFIFO(t *testing.T) {
	q := NewQueue()

	o1 := &Order{ID: 1, Type: OrderTypeNormal, Priority: OrderPriorityNormal, CreatedAt: time.Now()}
	q.Push(o1)

	time.Sleep(10 * time.Millisecond)
	o2 := &Order{ID: 2, Type: OrderTypeNormal, Priority: OrderPriorityNormal, CreatedAt: time.Now()}
	q.Push(o2)

	// First added should come out first
	first := q.Pop()
	if first.ID != 1 {
		t.Errorf("Expected ID 1 to be first, got ID %d", first.ID)
	}
}
func TestOrderStats(t *testing.T) {
	q := NewQueue()

	// Add a few orders
	AddOrder(q, OrderTypeNormal)
	AddOrder(q, OrderTypeNormal)
	AddOrder(q, OrderTypeVIP)

	if GetTotalCount() < 3 {
		t.Errorf("Expected total count at least 3, got %d", GetTotalCount())
	}

	if GetCountByType(OrderTypeNormal) < 2 {
		t.Errorf("Expected normal count at least 2, got %d", GetCountByType(OrderTypeNormal))
	}

	if GetCountByType(OrderTypeVIP) < 1 {
		t.Errorf("Expected VIP count at least 1, got %d", GetCountByType(OrderTypeVIP))
	}

	// Pop one and mark complete (status update)
	ord := q.Pop()
	ord.Status = OrderStatusComplete

	if GetCompletedCount() < 1 {
		t.Errorf("Expected completed count at least 1, got %d", GetCompletedCount())
	}
}

func TestIdenticalTimestampTieBreaker(t *testing.T) {
	q := NewQueue()
	now := time.Now()

	// Three orders created with exact same timestamp
	o1 := &Order{ID: 1001, Type: OrderTypeNormal, Priority: OrderPriorityNormal, CreatedAt: now}
	o2 := &Order{ID: 1002, Type: OrderTypeNormal, Priority: OrderPriorityNormal, CreatedAt: now}
	o3 := &Order{ID: 1003, Type: OrderTypeNormal, Priority: OrderPriorityNormal, CreatedAt: now}

	// Push in mixed order
	q.Push(o3)
	q.Push(o1)
	q.Push(o2)

	// Should come out in ID order: 1001, 1002, 1003
	results := []int{1001, 1002, 1003}
	for _, expectedID := range results {
		got := q.Pop()
		if got.ID != expectedID {
			t.Errorf("Expected ID %d, got %d", expectedID, got.ID)
		}
	}
}

func TestQueuePause(t *testing.T) {
	q := NewQueue()
	q.Push(&Order{ID: 1, Type: OrderTypeNormal, Priority: OrderPriorityNormal, CreatedAt: time.Now()})

	q.SetPaused(true)
	if q.Pop() != nil {
		t.Error("Pop should return nil when queue is paused")
	}

	q.SetPaused(false)
	if q.Pop() == nil {
		t.Error("Pop should return order when queue is unpaused")
	}
}
