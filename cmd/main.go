package main

import (
	"strings"
	"time"

	"github.com/feedme/order-controller/internal/bot"
	"github.com/feedme/order-controller/internal/manager"
	"github.com/feedme/order-controller/internal/order"
	"github.com/feedme/order-controller/internal/utils"
)

func main() {
	utils.LogRaw("McDonald's Order Controller - Starting Simulation")
	utils.LogRaw(strings.Repeat(" ", 5))

	sm := manager.NewSystemManager()

	// Add a Fast Bot (5s processing)
	sm.AddBot(bot.BotTypeFast)

	sm.AddOrder(order.OrderTypeNormal)
	sm.AddOrder(order.OrderTypeNormal)
	sm.AddOrder(order.OrderTypeVIP)
	sm.AddOrder(order.OrderTypeNormal)

	// Add a Slow Bot (10s processing)
	sm.AddBot(bot.BotTypeSlow)

	sm.AddOrder(order.OrderTypeVIP)
	sm.AddOrder(order.OrderTypeNormal)

	// Add another Fast Bot
	sm.AddBot(bot.BotTypeFast)

	// Decrease Bot (should stop one processing order)
	time.Sleep(3 * time.Second)
	sm.RemoveBot("")

	time.Sleep(3 * time.Second)
	sm.AddBot(bot.BotTypeSlow)

	time.Sleep(35 * time.Second)

	utils.LogRaw(strings.Repeat(" ", 5))
	utils.LogRaw(strings.Repeat("=", 50))
	utils.LogRaw(sm.GetSummary())
}
