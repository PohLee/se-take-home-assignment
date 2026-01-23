package manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/feedme/order-controller/internal/bot"
	"github.com/feedme/order-controller/internal/order"
	"github.com/feedme/order-controller/internal/utils"
)

// SystemManager orchestrates the order queue and bot pool, handling job assignment
// and tracking simulation statistics.
type SystemManager struct {
	OrderQueue  *order.Queue
	BotPool     *bot.Pool
	cancelFuncs map[string]context.CancelFunc
	mu          sync.Mutex
	wg          sync.WaitGroup
}

// NewSystemManager initializes and returns a new SystemManager with an empty queue and pool.
func NewSystemManager() *SystemManager {
	return &SystemManager{
		OrderQueue:  order.NewQueue(),
		BotPool:     bot.NewPool(),
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

func (m *SystemManager) botLoop(ctx context.Context, b *bot.Bot) {
	defer m.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Check for available orders
			ord := m.OrderQueue.Pop()
			if ord == nil {
				// No orders, wait a bit and check again
				time.Sleep(500 * time.Millisecond)
				continue
			}

			// Process the order
			completed := b.ProcessOrder(ctx, ord, nil)
			if !completed {
				// Order was cancelled, put it back to the front of the queue
				ord.Status = order.OrderStatusPending
				m.OrderQueue.PushFront(ord)
			}
		}
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
