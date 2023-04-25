package gocurrency

import (
	"fmt"
	"time"
)

func BuildRace() {
	car1 := "Ferrari"
	car2 := "Lamborghini"

	// Create a Goroutine for each car
	go race(car1)
	go race(car2)

	// Wait for the race to finish
	time.Sleep(5 * time.Second)

	fmt.Println("Race over!")
}

func race(car string) {
	for i := 0; i < 5; i++ {
		fmt.Println(car, "is racing...")
		time.Sleep(1 * time.Second)
	}
}
