package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}

	for i := 0; i < 100; i++ {
		go func() {
			once.Do(OnlyOnce)
		}()
	}
	fmt.Println("Counter:", counter)
}
