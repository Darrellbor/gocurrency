package gocurrency

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

func (c *Counter) increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

func (c *Counter) getValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func DataAccess() {
	counter := &Counter{}

	for i := 1; i <= 5; i++ {
		go func() {
			counter.increment()
		}()
	}

	time.Sleep(time.Second) // Wait for goroutines to finish
	fmt.Println("Counter value:", counter.getValue())
}
