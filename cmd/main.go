package main

import (
	"strings"
	"time"

	"github.com/feedme/order-controller/internal/manager"
	"github.com/feedme/order-controller/internal/order"
	"github.com/feedme/order-controller/internal/utils"
)

func main() {
	utils.LogRaw("McDonald's Order Controller - Starting Simulation")
	utils.LogRaw(strings.Repeat(" ", 5))

	sm := manager.NewSystemManager()

	// Scenario: Add some orders
	sm.AddOrder(order.OrderTypeNormal) // Order 1
	sm.AddOrder(order.OrderTypeVIP)    // Order 2 (Should be processed before Order 1)
	sm.AddOrder(order.OrderTypeNormal) // Order 3

	// Initial Bots
	sm.AddBot()
	sm.AddBot()

	// Wait a bit for bots to pick up
	time.Sleep(2 * time.Second)

	// Add more orders
	sm.AddOrder(order.OrderTypeVIP)    // Order 4
	sm.AddOrder(order.OrderTypeNormal) // Order 5

	// Increase Bots
	sm.AddBot() // Total 3 bots

	// Decrease Bot (should stop one processing order)
	time.Sleep(3 * time.Second)
	sm.RemoveBot("")

	time.Sleep(3 * time.Second)
	sm.AddBot()

	// Let simulation run for a while to complete some orders
	// 10s per order, with 2-3 bots it should take around 20-30s for 5 orders
	time.Sleep(25 * time.Second)

	utils.LogRaw(strings.Repeat(" ", 5))
	utils.LogRaw(sm.GetSummary())
}
