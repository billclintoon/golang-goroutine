package golanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "bill"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "masukan data"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "only in"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)
	go func() {

		channel <- "bill"
		channel <- "2"
	}()
	go func() {

		fmt.Println(<-channel)
		fmt.Println("Selesai")
		time.Sleep(2 * time.Second)
	}()
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Perulangan ke " + strconv.Itoa(i))
		}
		close(channel)
	}()
	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}

func TestSelectionChannel(t *testing.T) {
	channel1 := make(chan string)
	chnanel2 := make(chan string)

	defer close(channel1)
	defer close(chnanel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(chnanel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1", data)
			counter++
		case data := <-chnanel2:
			fmt.Println("Data dari chnanel2", data)
			counter++
		}
		if counter ==2 {
			break
		}
	}
}

func TestDefaultSelectionChannel(t *testing.T) {
	channel1 := make(chan string)
	chnanel2 := make(chan string)

	defer close(channel1)
	defer close(chnanel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(chnanel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1", data)
			counter++
		case data := <-chnanel2:
			fmt.Println("Data dari chnanel2", data)
			counter++
			default :
			fmt.Println("Menunggu Data")
		}
		if counter ==2 {
			break
		}
	}
}
