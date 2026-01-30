package bot

import "testing"

func TestForEach(t *testing.T) {
	p := NewPool()
	p.AddBot(BotTypeFast) // Idle
	p.AddBot(BotTypeSlow) // Idle
	b3 := p.AddBot(BotTypeFast) 
	b3.Status = BotStatusOffline

	// Test case 1: Iterate all bots (empty status)
	countAll := 0
	p.ForEach("", func(b *Bot) {
		countAll++
	})
	if countAll != 3 {
		t.Errorf("Expected to iterate 3 bots, visited %d", countAll)
	}

	// Test case 2: Iterate only Idle bots
	countIdle := 0
	p.ForEach(BotStatusIdle, func(b *Bot) {
		countIdle++
		if b.Status != BotStatusIdle {
			t.Errorf("Expected bot status Idle, got %s", b.Status)
		}
	})
	if countIdle != 2 {
		t.Errorf("Expected to iterate 2 Idle bots, visited %d", countIdle)
	}

	// Test case 3: Iterate only Offline bots
	countOffline := 0
	p.ForEach(BotStatusOffline, func(b *Bot) {
		countOffline++
		if b.Status != BotStatusOffline {
			t.Errorf("Expected bot status Offline, got %s", b.Status)
		}
	})
	if countOffline != 1 {
		t.Errorf("Expected to iterate 1 Offline bot, visited %d", countOffline)
	}
}
