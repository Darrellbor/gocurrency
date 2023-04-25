package gocurrency

import (
	"fmt"
	"sync"
	"time"
)

func SimpleWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment WaitGroup counter
		go func(num int) {
			defer wg.Done() // decrement WaitGroup counter when done
			fmt.Printf("goroutine %d\n", num)
		}(i)
		time.Sleep(time.Duration(1 * time.Second))
	}

	wg.Wait() // blocks until WaitGroup counter is zero
	fmt.Println("All goroutines have finished executing.")
}
