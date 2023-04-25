package main

import (
	"fmt"

	"github.com/Darrellbor/gocurrency"
)

func main() {
	fmt.Println("------------------Build Race-------------------------")
	gocurrency.BuildRace()
	fmt.Println("-----------------Simple Channel--------------------------")
	gocurrency.SimpleChannel()
	fmt.Println("------------------Buffered Channel-------------------------")
	gocurrency.BufferedChannel()
	fmt.Println("----------------Select Race---------------------------")
	gocurrency.SelectRace()
	fmt.Println("----------------Simple Wait Group---------------------------")
	gocurrency.SimpleWaitGroup()
	fmt.Println("----------------Mutexes---------------------------")
	gocurrency.Mutexes()
}
