package go_goroutines

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{}) // create a new map and store some data in it (key, value)
var wg = sync.WaitGroup{}              // create a wait group

func WaitCondition(value int) {
	defer wg.Done() // decrement the wait group counter when the goroutine completes

	cond.L.Lock() // acquire the lock
	cond.Wait()   // wait for the signal

	fmt.Println("Done:", value) // print the value
	cond.L.Unlock()             // release the lock
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ { // create 10 goroutines
		wg.Add(1)           // increment the wait group counter
		go WaitCondition(i) // create a new goroutine
	}

	go func() { // create a new anonymous goroutine
		for i := 0; i < 10; i++ { // create 10 goroutines
			time.Sleep(1 * time.Second) // wait for 1 second
			cond.Signal()               // signal the condition
		}
	}()

	// go func() { // create a new anonymous goroutine
	// 	for i := 0; i < 10; i++ { // create 10 goroutines
	// 		time.Sleep(1 * time.Second) // wait for 1 second
	// 		cond.Broadcast() // broadcast the condition
	// 	}

	// }()

	wg.Wait() // wait for all the goroutines to complete
}

var done = false

// func read(name string, c *sync.Cond) { // read function to read the data from the map
// 	c.L.Lock()  // acquire the lock
// 	for !done { // loop until the done flag is set
// 		c.Wait() // wait for the signal
// 	} // release the lock
// 	log.Println(name, "starts reading") // print the value
// 	c.L.Unlock()                        // release the lock
// }

func read(name string, c *sync.Cond) { // read function to read the data from the map
	defer wg.Done() // decrement the wait group counter when the goroutine completes
	c.L.Lock()      // acquire the lock
	// for !done { // loop until the done flag is set
	// 	c.Wait() // wait for the signal
	// } // release the lock
	c.Wait()                            // wait for the signal
	log.Println(name, "starts reading") // print the value
	c.L.Unlock()                        // release the lock
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing") // print the value
	time.Sleep(1 * time.Second)         // wait for 1 second
	c.L.Lock()                          // acquire the lock
	// done = true                         // set the done flag
	c.L.Unlock()                   // release the lock
	log.Println(name, "wakes all") // print the value
	c.Broadcast()                  // broadcast the condition
}

func TestSyncCond(t *testing.T) {

	// go read("reader1", cond) // create a new goroutine
	// go read("reader2", cond) // create a new goroutine
	// go read("reader3", cond) // create a new goroutine
	// write("writer", cond)    // create a new goroutine

	wg.Add(3) // increment the wait group counter
	for i := 0; i < 3; i++ {
		go read(fmt.Sprintf("reader%d", i), cond) // create a new goroutine
	}
	write("writer", cond) // create a new goroutine

	// time.Sleep(time.Second * 5) // wait for 3 seconds
	wg.Wait() // wait for all the goroutines to complete

	log.Println("All Done") // print the value
}
