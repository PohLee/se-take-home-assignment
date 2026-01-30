package bot

import (
	"context"
	"testing"
	"time"

	"github.com/feedme/order-controller/internal/order"
)

func TestPool(t *testing.T) {
	p := NewPool()

	// Test AddBot
	b1 := p.AddBot(BotTypeFast)
	if len(b1.ID) != 3 {
		t.Errorf("Expected bot ID length 3, got %d", len(b1.ID))
	}
	if p.GetActiveBotsCount() != 1 {
		t.Errorf("Expected 1 active bot, got %d", p.GetActiveBotsCount())
	}

	b2 := p.AddBot(BotTypeSlow)
	if b2.ID == "" {
		t.Error("Expected b2 ID to not be empty")
	}
	if p.GetActiveBotsCount() != 2 {
		t.Errorf("Expected 2 active bots, got %d", p.GetActiveBotsCount())
	}

	// Test RemoveBot by ID
	p.RemoveBot(b1.ID)
	if p.GetActiveBotsCount() != 1 {
		t.Errorf("Expected 1 active bot after removal, got %d", p.GetActiveBotsCount())
	}

	// Test RemoveBot default (last)
	p.RemoveBot("")
	if p.GetActiveBotsCount() != 0 {
		t.Errorf("Expected 0 active bots after second removal, got %d", p.GetActiveBotsCount())
	}
}

func TestProcessOrder(t *testing.T) {
	b := &Bot{ID: "testbot", Status: BotStatusIdle}
	ord := &order.Order{ID: 100, Priority: 10, Type: order.OrderTypeNormal}

	ctx := context.Background()

	// Test successful processing (mocking behavior with context cancellation)
	cancelCtx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	completed := b.ProcessOrder(cancelCtx, ord, nil)
	if completed {
		t.Error("Expected order to be cancelled, but it completed")
	}
	if b.Status != BotStatusOffline {
		t.Errorf("Expected bot status Offline after cancellation, got %v", b.Status)
	}
}
