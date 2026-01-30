package bot

import (
	"sync"

	"github.com/feedme/order-controller/internal/utils"
)

// Pool manages a collection of bot workers and provides thread-safe operations
// for adding, removing, and counting active bots.
type Pool struct {
	bots []*Bot
	mu   sync.Mutex
}

// NewPool initializes and returns a new empty bot Pool.
func NewPool() *Pool {
	return &Pool{
		bots: make([]*Bot, 0),
	}
}

// AddBot creates a new bot with a random 6-character ID, initializes its status
// to Idle, and returns the newly created Bot.
func (p *Pool) AddBot(botType BotTypeEnum) *Bot {
	p.mu.Lock()
	defer p.mu.Unlock()

	newID := utils.GenerateRandomID(3)
	newBot := &Bot{
		ID:     newID,
		Status: BotStatusIdle,
		Type:   botType,
	}
	p.bots = append(p.bots, newBot)
	return newBot
}

// RemoveBot looks for a bot with the specified ID and removes it from the pool.
// If an empty string is provided, it removes the last bot added.
// Returns the removed Bot or nil if no matching bot is found or the pool is empty.
func (p *Pool) RemoveBot(id string) *Bot {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.bots) == 0 {
		return nil
	}

	var targetBot *Bot
	var targetIndex int = -1

	if id != "" {
		for i, b := range p.bots {
			if b.ID == id {
				targetBot = b
				targetIndex = i
				break
			}
		}
	} else {
		// Default to last bot
		targetIndex = len(p.bots) - 1
		targetBot = p.bots[targetIndex]
	}

	if targetBot == nil {
		return nil
	}

	// Remove from slice
	p.bots = append(p.bots[:targetIndex], p.bots[targetIndex+1:]...)

	// Set status to Offline to signal stoppage
	targetBot.Status = BotStatusOffline
	return targetBot
}

// GetActiveBotsCount returns the number of bots currently in the pool
// that are not marked as Offline.
func (p *Pool) GetActiveBotsCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	count := 0
	for _, b := range p.bots {
		if b.Status != BotStatusOffline {
			count++
		}
	}
	return count
}

// ForEach executes a function for every bot in the pool (thread-safe).
// If status is provided, it only iterates over bots with that status.
func (p *Pool) ForEach(status BotStatusEnum, fn func(*Bot)) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, b := range p.bots {
		if status == "" || b.Status == status {
			fn(b)
		}
	}
}
