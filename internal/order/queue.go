package order

import (
	"container/heap"
	"sync"
)

// PriorityQueue implements heap.Interface and holds Orders.
type PriorityQueue []*Order

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Higher priority comes first
	if pq[i].Priority != pq[j].Priority {
		return pq[i].Priority > pq[j].Priority
	}
	// For same priority, earlier CreatedAt comes first
	if !pq[i].CreatedAt.Equal(pq[j].CreatedAt) {
		return pq[i].CreatedAt.Before(pq[j].CreatedAt)
	}
	// Tie breaker: use ID to ensure strict FIFO if timestamps are identical
	return pq[i].ID < pq[j].ID
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Order)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Queue manages the order priority queue with thread safety.
// NOTE: This is an in-memory implementation for demonstration and prototype purposes.
// In a production environment, this should be replaced by a robust task queue system
// such as Asynq or Machinery, backed by a persistent broker like Redis or RabbitMQ
// to ensure task persistence, scalability, and better reliability.
type Queue struct {
	pq PriorityQueue
	mu sync.Mutex
	// Notify is used as an internal signaling mechanism to wake up bot workers
	// immediately when an order is available, minimizing idle polling.
	Notify chan struct{}
	paused bool //  allows a manager to "freeze" bots from picking up orders
}

// NewQueue initializes and returns a new empty order priority Queue.
func NewQueue() *Queue {
	q := &Queue{
		pq:     make(PriorityQueue, 0),
		Notify: make(chan struct{}, 100), // Buffered to prevent blocking producers
	}
	heap.Init(&q.pq)
	return q
}

func (q *Queue) SetPaused(paused bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.paused = paused
	// If we just unpaused, signal bots to check the queue again
	if !paused {
		select {
		case q.Notify <- struct{}{}:
		default:
		}
	}
}

// Push adds a new order to the priority queue.
func (q *Queue) Push(order *Order) {
	q.mu.Lock()
	heap.Push(&q.pq, order)
	q.mu.Unlock()

	// Signal that a new order is available
	select {
	case q.Notify <- struct{}{}:
	default:
		// Channel full, signal already pending
	}
}

// Pop removes and returns the highest-priority order from the queue.
// Returns nil if the queue is empty.
func (q *Queue) Pop() *Order {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.paused || q.pq.Len() == 0 {
		return nil
	}
	return heap.Pop(&q.pq).(*Order)
}

// PushFront adds an order back to the queue (e.g., after a bot cancellation).
func (q *Queue) PushFront(order *Order) {
	q.mu.Lock()
	heap.Push(&q.pq, order)
	q.mu.Unlock()

	// Signal that a new order is available
	select {
	case q.Notify <- struct{}{}:
	default:
	}
}

// Peek returns the highest-priority order without removing it from the queue.
func (q *Queue) Peek() *Order {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.pq.Len() == 0 {
		return nil
	}
	return q.pq[0]
}

// Len returns the number of orders currently in the queue.
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.pq.Len()
}
