package order

import "time"

type OrderStatusEnum string

const (
	OrderStatusPending    OrderStatusEnum = "PENDING"
	OrderStatusProcessing OrderStatusEnum = "PROCESSING"
	OrderStatusComplete   OrderStatusEnum = "COMPLETE"
)

type OrderTypeEnum string

const (
	OrderTypeNormal OrderTypeEnum = "Normal"
	OrderTypeVIP    OrderTypeEnum = "VIP"
	OrderTypeUrgent OrderTypeEnum = "Urgent"
)

const (
	OrderPriorityNormal int = 10
	OrderPriorityVIP    int = 20
	OrderPriorityUrgent int = 99
)

var PriorityMap = map[OrderTypeEnum]int{
	OrderTypeNormal: OrderPriorityNormal,
	OrderTypeVIP:    OrderPriorityVIP,
	OrderTypeUrgent: OrderPriorityUrgent,
}

type Order struct {
	ID          int
	Type        OrderTypeEnum
	Status      OrderStatusEnum
	Priority    int
	CreatedAt   time.Time
	ProcessedAt *time.Time
	CompletedAt *time.Time
}
