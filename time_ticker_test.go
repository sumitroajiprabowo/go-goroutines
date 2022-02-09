package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second * 1)
	for i := 0; i < 5; i++ {
		<-ticker.C // Wait for the ticker to expire.
		fmt.Println(time.Now())
	}

	fmt.Println("Completed Ticker using for loop")

	go func() {
		time.Sleep(10 * time.Second)
		ticker.Stop()
		fmt.Println("Completed Ticker using goroutine")
	}()

	for time := range ticker.C { // Wait for the ticker to expire.
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(time.Second * 1)
	for i := 0; i < 5; i++ {
		<-channel // Wait for the ticker to expire.
		fmt.Println(time.Now())
		fmt.Println("Completed using tick")
	}
}
