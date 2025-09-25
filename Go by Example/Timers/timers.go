package main

import (
	"fmt"
	"time"
)

func main() {
	firstTimer := time.NewTimer(2 * time.Second)

	<-firstTimer.C
	fmt.Println("Timer 1 fired")

	secondTimer := time.NewTimer(1 * time.Second)

	go func() {
		<-secondTimer.C
		fmt.Println("Timer 2 fired")
	}()

	stopSecondTimer := secondTimer.Stop()
	if stopSecondTimer {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
