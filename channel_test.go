package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	c := make(chan int)

	// close channel using defer
	// defer close(c)

	go func() {
		time.Sleep(2 * time.Second)
		c <- 1
		fmt.Println("Done sender data to channel")
	}()

	fmt.Println(<-c)

	close(c)
}

func GiveMeResponse(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Hello"
}

func TestChannelAsParameter(t *testing.T) {
	ch1 := make(chan string)
	defer close(ch1)

	go GiveMeResponse(ch1)

	fmt.Println(<-ch1)

}

// Send only channel
func OnlyIn(ch2 chan<- string) {
	time.Sleep(2 * time.Second)
	ch2 <- "Terimakasih"
}

// Receive only channel
func OnlyOut(ch2 <-chan string) {
	fmt.Println(<-ch2)
}

func TestInOutChannel(t *testing.T) {
	ch2 := make(chan string)
	defer close(ch2)

	go OnlyIn(ch2)
	go OnlyOut(ch2)

	time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	ch3 := make(chan string, 3)
	defer close(ch3)

	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "Data 1"
		ch3 <- "Data 2"
		ch3 <- "Data 3"
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-ch3)
		}
	}()

	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	ch4 := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			ch4 <- fmt.Sprintf("Menerima Data %d", i)
		}
		close(ch4)
	}()

	for i := range ch4 {
		fmt.Println(i)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	defer close(ch1)
	defer close(ch2)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	// select {
	// case data := <-ch1:
	// 	fmt.Println("Data from ch1", data)
	// case data := <-ch2:
	// 	fmt.Println("Data from ch2", data)
	// }

	counter := 0

	for {
		select {
		case data := <-ch1:
			fmt.Println("Data from ch1", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data from ch2", data)
			counter++
		default:
			fmt.Println("Waiting for data !!!")
		}
		if counter == 2 {
			break
		}
	}
}
