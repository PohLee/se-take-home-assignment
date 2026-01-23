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
	return pq[i].CreatedAt.Before(pq[j].CreatedAt)
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
type Queue struct {
	pq PriorityQueue
	mu sync.Mutex
}

// NewQueue initializes and returns a new empty order priority Queue.
func NewQueue() *Queue {
	q := &Queue{
		pq: make(PriorityQueue, 0),
	}
	heap.Init(&q.pq)
	return q
}

// Push adds a new order to the priority queue.
func (q *Queue) Push(order *Order) {
	q.mu.Lock()
	defer q.mu.Unlock()
	heap.Push(&q.pq, order)
}

// Pop removes and returns the highest-priority order from the queue.
// Returns nil if the queue is empty.
func (q *Queue) Pop() *Order {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.pq.Len() == 0 {
		return nil
	}
	return heap.Pop(&q.pq).(*Order)
}

// PushFront adds an order back to the queue (e.g., after a bot cancellation).
// In this implementation, it's equivalent to Push since the priority/time logic
// will naturally put it back in the correct position.
func (q *Queue) PushFront(order *Order) {
	q.mu.Lock()
	defer q.mu.Unlock()
	// To push to front in a priority queue, we could just push it normally
	// and if priorities are the same, it will be sorted by time.
	// However, if we want it to be ABSOLUTELY first, we'd need to manipulate time.
	// But according to requirements, it just goes back to PENDING.
	heap.Push(&q.pq, order)
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
