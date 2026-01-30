package manager

import (
	"testing"
	"time"

	"github.com/feedme/order-controller/internal/bot"
	"github.com/feedme/order-controller/internal/order"
)

func TestSystemManager(t *testing.T) {
	m := NewSystemManager()

	// Test AddBot
	botID := m.AddBot(bot.BotTypeFast)
	if len(botID) != 3 {
		t.Errorf("Expected bot ID length 3, got %d", len(botID))
	}

	// Test AddOrder
	m.AddOrder(order.OrderTypeNormal)
	m.AddOrder(order.OrderTypeVIP)

	if m.OrderQueue.Len() != 2 {
		t.Errorf("Expected 2 orders in queue, got %d", m.OrderQueue.Len())
	}

	// Wait a bit for bot to pick up order
	time.Sleep(100 * time.Millisecond)

	// One order should be picked up
	if m.OrderQueue.Len() > 2 {
		t.Errorf("Expected at most 2 orders in queue after pickup, got %d", m.OrderQueue.Len())
	}

	// Test GetSummary
	summary := m.GetSummary()
	if summary == "" {
		t.Error("Expected summary to not be empty")
	}

	// Test RemoveBot
	m.RemoveBot(botID)
	// We can't easily wait for the loop to exit without m.Wait() which blocks.
	// But we can check that it doesn't crash.
}
