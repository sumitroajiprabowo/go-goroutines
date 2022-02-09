package go_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func TestExample(t *testing.T) {
	runtime.GOMAXPROCS(2)

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}
