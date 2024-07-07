package worker

import (
	"sync"
	"time"

	"github.com/chadsmith12/coffeeshop/pkgs/customers"
	"github.com/chadsmith12/coffeeshop/pkgs/simulation"
)

type Worker struct {
    id int
    customersProcessed []customers.Customer
    simulatedDay simulation.SimulatedDay
}

func NewWorker(id int, simulatedDay simulation.SimulatedDay) *Worker {
   return &Worker{
   	id:                 id,
   	customersProcessed: make([]customers.Customer, 0, 10),
    simulatedDay: simulatedDay,
   } 
}

func (w *Worker) ProcessedCustomers() int {
    return len(w.customersProcessed)
}

func (w *Worker) ProcessCustomer(customer customers.Customer) {
    time.Sleep(w.simulatedDay.Minute())

    w.customersProcessed = append(w.customersProcessed, customer)
}

func (w *Worker) Work(wg *sync.WaitGroup, newCustomerChan <-chan struct{}, closedChan <-chan struct{}, line *customers.CustomerQueue) {
    for {
        select {
        case <-newCustomerChan:
            customer, _ := line.Remove()
            w.ProcessCustomer(customer)
            wg.Done()
        case <-closedChan: 
            for line.Len() > 0 {
                customer, _ := line.Remove()
                w.ProcessCustomer(customer)
                wg.Done()
            }
            return
        }
    }
}
