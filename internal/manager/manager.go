package manager

import (
	"context"
	"fmt"
	"sync"

	"github.com/feedme/order-controller/internal/bot"
	"github.com/feedme/order-controller/internal/event"
	"github.com/feedme/order-controller/internal/order"
	"github.com/feedme/order-controller/internal/utils"
)

// SystemManager orchestrates the order queue and bot pool, handling job assignment
// and tracking simulation statistics.
type SystemManager struct {
	OrderQueue  *order.Queue
	BotPool     *bot.Pool
	EventBus    *event.EventBus
	cancelFuncs map[string]context.CancelFunc
	mu          sync.Mutex
	wg          sync.WaitGroup
}

// NewSystemManager initializes and returns a new SystemManager with an empty queue and pool.
func NewSystemManager() *SystemManager {
	eb := event.NewEventBus()
	order.Bus = eb // Link the order manager to the bus

	return &SystemManager{
		OrderQueue:  order.NewQueue(),
		BotPool:     bot.NewPool(),
		EventBus:    eb,
		cancelFuncs: make(map[string]context.CancelFunc),
	}
}

// AddOrder creates a new order of the specified type and adds it to the system queue.
func (m *SystemManager) AddOrder(orderType order.OrderTypeEnum) {
	order.AddOrder(m.OrderQueue, orderType)
}

// AddBot creates a new bot, adds it to the pool, and starts its processing loop.
// Returns the ID of the newly created bot.
func (m *SystemManager) AddBot() string {
	b := m.BotPool.AddBot()
	utils.Log("Bot #%s added into pool - Status: ACTIVE", b.ID)

	// Start the bot worker loop
	ctx, cancel := context.WithCancel(context.Background())
	m.mu.Lock()
	m.cancelFuncs[b.ID] = cancel
	m.mu.Unlock()

	m.wg.Add(1)
	go m.botLoop(ctx, b)
	return b.ID
}

// RemoveBot stops and removes a bot from the system. If id is empty, the last
// added bot is removed.
func (m *SystemManager) RemoveBot(id string) {
	b := m.BotPool.RemoveBot(id)
	if b == nil {
		utils.Log("No bots available to remove")
		return
	}

	m.mu.Lock()
	if cancel, ok := m.cancelFuncs[b.ID]; ok {
		cancel()
		delete(m.cancelFuncs, b.ID)
	}
	m.mu.Unlock()
	utils.Log("Bot #%s removed from pool", b.ID)
}

// botLoop is the main worker loop for a bot. It waits for order availability
// signals or context cancellation. It is event-reactive, eliminating the need
// for periodic polling.
func (m *SystemManager) botLoop(ctx context.Context, b *bot.Bot) {
	defer m.wg.Done()

	for {
		// Check for cancellation BEFORE picking up work to avoid processing
		// orders after the bot has been decommissioned.
		select {
		case <-ctx.Done():
			return
		default:
		}

		// First, try to pop any existing orders immediately.
		ord := m.OrderQueue.Pop()
		if ord != nil {
			// Notify the system that an order has been assigned.
			m.EventBus.Publish(event.Event{
				Type: event.OrderAssigned,
				Data: ord,
			})
			m.processAndEmit(ctx, b, ord)
			continue
		}

		// No orders available, wait for notification from the queue or cancellation.
		select {
		case <-ctx.Done():
			return
		case <-m.OrderQueue.Notify:
			// New order might be available, loop back to Pop.
			continue
		}
	}
}

// processAndEmit handles the actual bot processing of an order and publishes
// completion or cancellation events to the EventBus.
func (m *SystemManager) processAndEmit(ctx context.Context, b *bot.Bot, ord *order.Order) {
	completed := b.ProcessOrder(ctx, ord, nil)
	if completed {
		m.EventBus.Publish(event.Event{
			Type: event.OrderCompleted,
			Data: ord,
		})
	} else {
		// Order was cancelled, put it back to the front of the queue
		ord.Status = order.OrderStatusPending
		m.OrderQueue.PushFront(ord)
		m.EventBus.Publish(event.Event{
			Type: event.OrderCancelled,
			Data: ord,
		})
	}
}

// Wait blocks until all active bot loops have finished.
func (m *SystemManager) Wait() {
	m.wg.Wait()
}

// GetSummary compiles and returns a formatted string of the current simulation statistics.
func (m *SystemManager) GetSummary() string {
	total := order.GetTotalCount()
	vip := order.GetCountByType(order.OrderTypeVIP)
	normal := order.GetCountByType(order.OrderTypeNormal)
	completed := order.GetCompletedCount()
	activeBots := m.BotPool.GetActiveBotsCount()
	pending := m.OrderQueue.Len()

	return fmt.Sprintf("\nFinal Status:\n- Total Orders Processed: %d (%d VIP, %d Normal)\n- Orders Completed: %d\n- Active Bots: %d\n- Pending Orders: %d",
		total, vip, normal, completed, activeBots, pending)
}
