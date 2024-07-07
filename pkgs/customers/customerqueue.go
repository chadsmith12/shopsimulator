package customers

import (
	"sync"

	"github.com/chadsmith12/coffeeshop/pkgs/queue"
)

type CustomerQueue struct {
    rwMutex sync.RWMutex
    line *queue.Queue[Customer]
}

func Start() *CustomerQueue {
    line := queue.Init[Customer]()

    return &CustomerQueue{
        rwMutex: sync.RWMutex{},
        line: line,
    }
}

func (q *CustomerQueue) Add(customer Customer) {
    q.rwMutex.Lock()
    defer q.rwMutex.Unlock()
    q.line.Enqueue(customer) 
}

func (q *CustomerQueue) Remove() (Customer, bool) {
    q.rwMutex.Lock()
    defer q.rwMutex.Unlock()
    return q.line.Deque()
}

func (q *CustomerQueue) Len() int {
    q.rwMutex.RLock()
    defer q.rwMutex.RUnlock()
    return q.line.Len()
}
