package coffeeshop

import (
	"fmt"
	"sync"
	"time"

	"github.com/chadsmith12/coffeeshop/pkgs/customers"
	"github.com/chadsmith12/coffeeshop/pkgs/simulation"
	"github.com/chadsmith12/coffeeshop/pkgs/worker"
)

type CoffeShop struct {
    hoursOpened time.Duration
    dayLength simulation.SimulatedDay
    customerQueue *customers.CustomerQueue
    closedCh chan struct{}
    newCustomerCh chan struct{}
    isOpen bool
    customersWaitGroup sync.WaitGroup
    baristas []*worker.Worker
}

func New(hoursOpened time.Duration, dayLength simulation.SimulatedDay) *CoffeShop {
    baristas := []*worker.Worker {
        worker.NewWorker(1, dayLength),
    }
    return &CoffeShop{
    	hoursOpened:        hoursOpened,
        dayLength:          dayLength,
    	customerQueue:      customers.Start(),
    	closedCh:           make(chan struct{}),
        newCustomerCh:      make(chan struct{}),
        isOpen:             false,
        customersWaitGroup: sync.WaitGroup{},
        baristas: baristas,
    }
}

// Opens the coffee shop and starts having it accept customers and orders for the hours it is opened
func (cs *CoffeShop) Open() {
    go func() {
        cs.isOpen = true
        shopTimer := time.NewTicker(cs.dayLength.Hour())
        defer shopTimer.Stop()
        timeOpened := 0
        for timeOpened <= int(cs.hoursOpened.Seconds()) {
            <-shopTimer.C
            fmt.Println("Been an hour/second")
            timeOpened++
        }
        cs.Close()
    }()

    for _, barista := range cs.baristas {
        go barista.Work(&cs.customersWaitGroup, cs.newCustomerCh, cs.closedCh, cs.customerQueue)
    }
}


// This will close the coffee shop and not have it accept anymore customers inside it's queue
func (cs *CoffeShop) Close() {
    cs.isOpen = false
    close(cs.closedCh)
}

// Returns a channel to read when the coffee shop has been closed.
func (cs *CoffeShop) Closed() <- chan struct{} {
    return cs.closedCh
}

// Wait will wait until the shop has closed and all customers have been processed.
func (cs *CoffeShop) Wait() {
    <-cs.Closed()
    cs.customersWaitGroup.Wait()
}

// Adds a customer to the queue/line in the coffee shop to get an order
func (cs *CoffeShop) AcceptCustomer(customer customers.Customer) {
    select {
    case <-cs.closedCh: break
    default: 
        cs.customerQueue.Add(customer)
        cs.customersWaitGroup.Add(1)
        go cs.notify() 
    }
}

// Tells you the number of customers that were processed
func (cs *CoffeShop) CustomersProcessed() int {
    sum := 0
    for _, barista := range cs.baristas {
       sum += barista.ProcessedCustomers()
    }
    
    return sum
}

// notify is used to notify the cashier/barista that there is a new customer to process
func (cs *CoffeShop) notify() {
    cs.newCustomerCh <- struct{}{}
}

// this will process all the remaining customers that are in line
func (cs *CoffeShop) processRemainingCustomers(barista *worker.Worker) {
    for cs.customerQueue.Len() > 0 {
        cs.processNextCustomer(barista)   
    }
}

func (cs *CoffeShop) processNextCustomer(barista *worker.Worker) {
    customer, _ := cs.customerQueue.Remove()
    barista.ProcessCustomer(customer)
    cs.customersWaitGroup.Done()
}
