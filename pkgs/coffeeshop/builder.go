package coffeeshop

import (
	"sync"
	"time"

	"github.com/chadsmith12/coffeeshop/pkgs/customers"
	"github.com/chadsmith12/coffeeshop/pkgs/simulation"
	"github.com/chadsmith12/coffeeshop/pkgs/worker"
)

// A CoffeeShopBuilder allows you to build out a coffee shop with a number of workers fluently.
// Allows for an easy way to configure and build out coffee shops for testing out configurations.
type CoffeeShopBuilder struct {
    hoursOpened time.Duration
    dayLength simulation.SimulatedDay
    baristas []*worker.Worker
}

// Starts to build out a new CoffeeShop that uses the passed in the simulated day.
// The default CoffeShopBuidler sets the coffee shop to be opened for 8 hours.
// Also by default starts with at least 1 worker.
func NewBuilder(dayLength simulation.SimulatedDay) *CoffeeShopBuilder {
    barista := worker.NewWorker(1, dayLength)
    return &CoffeeShopBuilder{
    	hoursOpened: dayLength.Hour() * 8,
    	dayLength:   dayLength,
    	baristas:    []*worker.Worker{ barista },
    }
}

// Sets how long, in hours, the coffee shop is opened.
// This "hour" will be based on the simulated day length that is being used for the CoffeeShop
func (b *CoffeeShopBuilder) OpenedForHours(hoursOpened int) *CoffeeShopBuilder {
    b.hoursOpened = b.dayLength.Hour() * time.Duration(hoursOpened)

    return b
}

// Adds a new Barista to the shop.
func (b *CoffeeShopBuilder) AddBarista() *CoffeeShopBuilder {
    nextId := len(b.baristas) + 1
    b.baristas = append(b.baristas, worker.NewWorker(nextId, b.dayLength))

    return b
}

// Buuilds out a new CoffeeShop with the configuration supplied.
func (b *CoffeeShopBuilder) Build() *CoffeShop {
    return &CoffeShop{
    	hoursOpened:        b.hoursOpened,
    	dayLength:          b.dayLength,
    	line:               customers.Start(),
    	customerOrderCh:    make(chan customers.Customer),
    	closedCh:           make(chan struct{}),
    	newCustomerCh:      make(chan struct{}),
    	isOpen:             false,
    	customersWaitGroup: sync.WaitGroup{},
    	baristas:           b.baristas,
    }
}
