package bot

import "time"

type BotStatusEnum string

const (
	BotStatusIdle       BotStatusEnum = "IDLE"
	BotStatusProcessing BotStatusEnum = "PROCESSING"
	BotStatusOffline    BotStatusEnum = "OFFLINE"
	BotStatusFaulted    BotStatusEnum = "FAULTED"
)

type BotTypeEnum string

const (
	BotTypeFast BotTypeEnum = "FAST"
	BotTypeSlow BotTypeEnum = "SLOW"
)

var ProcessingTimeMap = map[BotTypeEnum]time.Duration{
	BotTypeFast: 5 * time.Second,
	BotTypeSlow: 10 * time.Second,
}

type Bot struct {
	ID             string
	Status         BotStatusEnum
	Type           BotTypeEnum
	CurrentOrderID *int
	
	// Business context consideration
	// Bot Model
	// Bot capabilities [Burger, Fench Fries, Fried Chicken]
}
