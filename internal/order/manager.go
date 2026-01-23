package order

import (
	"sync"
	"time"

	"github.com/feedme/order-controller/internal/utils"
)

var (
	allOrders   []*Order
	lastOrderID = 1000
	idMu        sync.Mutex
)

// AddOrder creates a new order with a unique ID and correct priority, then adds it to the queue.
func AddOrder(q *Queue, orderType OrderTypeEnum) *Order {
	idMu.Lock()
	defer idMu.Unlock()

	lastOrderID++
	orderID := lastOrderID

	newOrder := &Order{
		ID:        orderID,
		Type:      orderType,
		Status:    OrderStatusPending,
		Priority:  PriorityMap[orderType],
		CreatedAt: time.Now(),
	}

	allOrders = append(allOrders, newOrder)

	utils.Log("Order •%d (Priority: %d - %s) Created - Status: PENDING", newOrder.ID, newOrder.Priority, newOrder.Type)
	q.Push(newOrder)
	return newOrder
}

// GetTotalCount returns the total number of orders created.
func GetTotalCount() int {
	idMu.Lock()
	defer idMu.Unlock()
	return len(allOrders)
}

// GetCountByType returns the number of orders of a specific type.
func GetCountByType(orderType OrderTypeEnum) int {
	idMu.Lock()
	defer idMu.Unlock()

	count := 0
	for _, o := range allOrders {
		if o.Type == orderType {
			count++
		}
	}
	return count
}

// GetCompletedCount returns the number of orders with StatusComplete.
func GetCompletedCount() int {
	idMu.Lock()
	defer idMu.Unlock()

	count := 0
	for _, o := range allOrders {
		if o.Status == OrderStatusComplete {
			count++
		}
	}
	return count
}
