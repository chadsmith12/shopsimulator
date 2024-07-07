package main

import (
	"fmt"
	"time"

	"github.com/chadsmith12/coffeeshop/pkgs/coffeeshop"
	"github.com/chadsmith12/coffeeshop/pkgs/customers"
	"github.com/chadsmith12/coffeeshop/pkgs/simulation"
)

func main() {
    simulatedDay := simulation.NewSimulatedDay(24 * time.Second)
    coffeeShop := coffeeshop.NewBuilder(simulatedDay).Build()

    coffeeShop.Open()
    go func() {
        customerTimer := time.NewTicker(simulatedDay.Minute()) 
        defer customerTimer.Stop()
        customerId := 1
        for {
            select {
                case <-customerTimer.C:
                    coffeeShop.AcceptCustomer(customers.New(customerId))
                    customerId++ 
                case <-coffeeShop.Closed():
                    fmt.Println("We see the shop is closing...")
                    return
            }
        }
    }()
   
    coffeeShop.Wait()
    fmt.Printf("A total of %d customers were in line while the shop was open.\n", coffeeShop.CustomersProcessed())
}
