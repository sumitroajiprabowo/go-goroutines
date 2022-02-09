package go_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0        // 64-bit integer
	wg := sync.WaitGroup{} // WaitGroup is a struct

	for i := 0; i < 1000; i++ { // Loop 1000 times
		go func() { // Create a goroutine
			wg.Add(1)                  // Add 1 to WaitGroup
			for j := 0; j < 100; j++ { // Loop 100 times
				atomic.AddInt64(&x, 1) // Add 1 to x
			}
			wg.Done() // Done with goroutine
		}()
	}
	wg.Wait()                  // Wait for all goroutines to finish
	fmt.Println("Counter:", x) // output Counter: 100000
}

func TestAtomicAgain(t *testing.T) {

	var ops uint64 // 64-bit integer

	var wg sync.WaitGroup // WaitGroup is a struct

	for i := 0; i < 50; i++ { // Loop 50 times
		wg.Add(1) // Add 1 to WaitGroup

		go func() { // Create a goroutine
			for c := 0; c < 1000; c++ { // Loop 1000 times

				atomic.AddUint64(&ops, 1) // Add 1 to ops
			}
			wg.Done() // Done with goroutine
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("Total:", ops) // output ops: 50000
}

func TestAnotherAtomic(t *testing.T) {

	// Assigning values to the int32
	var (
		i int32 = 97
		j int32 = 48
		k int32 = 34754567
		l int32 = -355363
	)

	// Assigning constant
	// values to int32
	const (
		x int32 = 4
		y int32 = 2
	)

	// Calling AddInt32 method
	// with its parameters
	res_1 := atomic.AddInt32(&i, y)
	res_2 := atomic.AddInt32(&j, y-1)
	res_3 := atomic.AddInt32(&k, x-1)
	res_4 := atomic.AddInt32(&l, x)

	// Displays the output after adding
	// addr and delta atomically
	fmt.Println(res_1)
	fmt.Println(res_2)
	fmt.Println(res_3)
	fmt.Println(res_4)
}
