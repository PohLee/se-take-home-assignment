package bot

type BotStatusEnum string

const (
	BotStatusIdle       BotStatusEnum = "IDLE"
	BotStatusProcessing BotStatusEnum = "PROCESSING"
	BotStatusOffline    BotStatusEnum = "OFFLINE"
	BotStatusFaulted    BotStatusEnum = "FAULTED"
)

type Bot struct {
	ID             string
	Status         BotStatusEnum
	CurrentOrderID *int
}
