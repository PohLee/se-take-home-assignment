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

	// Simulate 10 seconds of processing
	select {
	case <-time.After(10 * time.Second):
		ord.Status = order.OrderStatusComplete
		doneAt := time.Now()
		ord.CompletedAt = &doneAt
		b.Status = BotStatusIdle
		b.CurrentOrderID = nil
		utils.Log("Bot #%s completed Order •%d - Status: COMPLETE (Processing time: 10s)", b.ID, ord.ID)
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
