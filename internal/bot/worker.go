package bot

import (
	"context"
	"time"

	"github.com/feedme/order-controller/internal/order"
	"github.com/feedme/order-controller/internal/utils"
)

// ProcessOrder handles the simulation of order processing (10 seconds per order).
// It transitions the bot and order states through PROCESSING and COMPLETE/OFFLINE.
// It returns true if the order was completed, and false if it was cancelled
// by a context signal (e.g., bot shutdown).
func (b *Bot) ProcessOrder(ctx context.Context, ord *order.Order, onComplete func(*order.Order)) bool {
	b.Status = BotStatusProcessing
	b.CurrentOrderID = &ord.ID
	ord.Status = order.OrderStatusProcessing
	now := time.Now()
	ord.ProcessedAt = &now

	utils.Log("Bot #%s picked up Order •%d - Status: PROCESSING", b.ID, ord.ID)

	// Determine processing duration from model map
	duration, ok := ProcessingTimeMap[b.Type]
	if !ok {
		// Fallback to default if type not found (should not happen with proper initialization)
		duration = 10 * time.Second
	}

	// Simulate processing
	select {
	case <-time.After(duration):
		ord.Status = order.OrderStatusComplete
		doneAt := time.Now()
		ord.CompletedAt = &doneAt
		b.Status = BotStatusIdle
		b.CurrentOrderID = nil
		utils.Log("Bot #%s completed Order •%d - Status: COMPLETE (Processing time: %v)", b.ID, ord.ID, duration)
		if onComplete != nil {
			onComplete(ord)
		}
		return true
	case <-ctx.Done():
		// Bot was removed or system stopped
		b.Status = BotStatusOffline
		b.CurrentOrderID = nil
		utils.Log("Bot #%s cancelled Order •%d - Status: CANCELLED", b.ID, ord.ID)
		return false
	}
}
