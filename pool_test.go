package go_goroutines

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"
)

type Customer struct {
	Name string
}

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	// pool put with pointer
	pool.Put(&Customer{Name: "Danu"})
	pool.Put(&Customer{Name: "Budi"})
	pool.Put(&Customer{Name: "Siti"})

	// pool.Put("Hello")
	// pool.Put("Hi")
	// pool.Put("How are you?")

	// pool.Put(pool.Get())

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(5 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(2 * time.Second)

	fmt.Println("Completed")

}

func TestAnotherPool(t *testing.T) {
	pool := sync.Pool{
		// New creates an object when the pool has nothing available to return.
		// New must return an interface{} to make it flexible. You have to cast
		// your type after getting it.
		New: func() interface{} {
			// Pools often contain things like *bytes.Buffer, which are
			// temporary and re-usable.
			return &bytes.Buffer{}
		},
	}

	// When getting from a Pool, you need to cast
	s := pool.Get().(*bytes.Buffer)
	// We write to the object
	s.Write([]byte("Hello "))
	// Then put it back
	pool.Put(s)

	// Pools can return dirty results

	// Get 'another' buffer
	s = pool.Get().(*bytes.Buffer)
	// Write to it
	s.Write([]byte("World"))
	// At this point, if GC ran, this buffer *might* exist already, in
	// which case it will contain the bytes of the string "dirtyappend"
	fmt.Println(s)
	// So use pools wisely, and clean up after yourself
	s.Reset()
	pool.Put(s)

	// When you clean up, your buffer should be empty
	s = pool.Get().(*bytes.Buffer)
	// Defer your Puts to make sure you don't leak!
	defer pool.Put(s)
	s.Write([]byte("reset!"))
	// This prints "reset!", and not "dirtyappendreset!"
	fmt.Println(s)

}

//Define a person structure with name and age variables
type Person struct {
	Name string
	Age  int
}

//Initialize sync.pool, and the new function is to create the person structure
func initPool() *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a person.")
			return &Person{}
		},
	}
}

//Main function
func TestPoolAgain(t *testing.T) {
	pool := initPool()
	person := pool.Get().(*Person)
	fmt.Println("get person from sync. Pool for the first time:", person)
	person.Name = "Jack"
	person.Age = 23
	pool.Put(person)
	fmt.Println("set object name:", person.Name)
	fmt.Println("set object age:", person.Age)
	fmt.Println("there is an object in the pool, call the get method to get:", pool.Get().(*Person))
	fmt.Println("there is no object in the pool, call the get method again:", pool.Get().(*Person))
}
