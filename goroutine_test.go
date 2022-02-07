package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("Ups")

	time.Sleep(1 * time.Second)

}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
