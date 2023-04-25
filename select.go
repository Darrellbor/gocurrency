package gocurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func SelectRace() {
	// Create two channels for the cars
	ferrari, lamborghini := make(chan string), make(chan string)

	// Start two Goroutines that add cars to the channels
	go addCarWithChan("Ferrari", ferrari)
	go addCarWithChan("Lamborghini", lamborghini)

	// Start a Goroutine that simulates the race
	go startRaceWithSelect(ferrari, lamborghini)

	// Wait for the race to finish
	time.Sleep(5 * time.Second)

	fmt.Println("Race over!")
}

func addCarWithChan(name string, c chan<- string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		c <- name
		fmt.Println(name, "added to the race!")
	}
	close(c)
}

func startRaceWithSelect(ferrari <-chan string, lamborghini <-chan string) {
	for {
		select {
		case car, ok := <-ferrari:
			if !ok {
				fmt.Println("Ferrari channel closed!")
				ferrari = nil
				break
			}
			fmt.Println(car, "is racing...")

		case car, ok := <-lamborghini:
			if !ok {
				fmt.Println("Lamborghini channel closed!")
				lamborghini = nil
				break
			}
			fmt.Println(car, "is racing...")
		}

		if ferrari == nil && lamborghini == nil {
			break
		}
		time.Sleep(time.Second)
	}

	fmt.Println("All cars have finished the race!")
}
