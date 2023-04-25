package gocurrency

import (
	"fmt"
	"sync"
)

var trackMutex sync.Mutex
var track [3]int

func Mutexes() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(car int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				trackMutex.Lock()
				track[car]++
				fmt.Printf("Car %d is on lap %d\n", car, track[car])
				trackMutex.Unlock()
			}
		}(i)
	}

	wg.Wait()
}
