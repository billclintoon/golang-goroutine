package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorl() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorl()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestGoroutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
