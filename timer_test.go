package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Second * 5) // 5 seconds
	fmt.Println(time.Now())                 // Prints: 2016-11-10 23:00:00.000000000 +0000 UTC

	time := <-timer.C // Wait for the timer to expire.

	fmt.Println(time)        // Prints: 2016-11-10 23:00:05.000000000 +0000 UTC
	fmt.Println("Completed") // Prints: Completed
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(time.Second * 5)
	fmt.Println(time.Now()) // Prints: 2016-11-10 23:00:00.000000000 +0000 UTC

	time := <-channel
	fmt.Println(time)        // Prints: 2016-11-10 23:00:05.000000000 +0000 UTC
	fmt.Println("Completed") // Prints: Completed
}

func TestAfterFuction(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)

	time.AfterFunc(time.Second*5, func() {
		fmt.Println(time.Now())
		wg.Done()
	})

	fmt.Println("Started")
	fmt.Println(time.Now())

	wg.Wait()
}
