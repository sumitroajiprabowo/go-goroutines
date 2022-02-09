package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}
	addToMap := func(value int) {
		defer group.Done()

		group.Add(1)
		data.Store(value, value)
	}

	for i := 0; i < 100; i++ {
		go addToMap(i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

}

func TestSimpleMap(t *testing.T) {
	data := sync.Map{}
	data.Store("key", "value")
	data.Store("key2", "value2")
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

func TestAnotherMap(t *testing.T) {
	wg := sync.WaitGroup{}
	m := sync.Map{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(j int) {
			m.Store(j, fmt.Sprintf("test %v", j))
			// fmt.Println(m.Load(j))
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Done.")

	for i := 0; i < 5; i++ {
		t, _ := m.Load(i)
		fmt.Println("for loop: ", t)
	}

	m.Range(func(k, v interface{}) bool {
		fmt.Println(v)
		return true
	})
}
