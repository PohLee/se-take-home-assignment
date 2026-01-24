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

	sm.AddBot()

	sm.AddOrder(order.OrderTypeNormal)
	sm.AddOrder(order.OrderTypeNormal)
	sm.AddOrder(order.OrderTypeVIP)
	sm.AddOrder(order.OrderTypeNormal)

	sm.AddBot()

	sm.AddOrder(order.OrderTypeVIP)
	sm.AddOrder(order.OrderTypeNormal)

	sm.AddBot()

	// Decrease Bot (should stop one processing order)
	time.Sleep(3 * time.Second)
	sm.RemoveBot("")

	time.Sleep(3 * time.Second)
	sm.AddBot()

	time.Sleep(35 * time.Second)

	utils.LogRaw(strings.Repeat(" ", 5))
	utils.LogRaw(strings.Repeat("=", 50))
	utils.LogRaw(sm.GetSummary())
}
