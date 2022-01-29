package belajar_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func TestHelloWorldGoRoutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number ",number)
}

func TestDisplayNumberGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(2 * time.Second)
}

func TestChannelGoroutine(t *testing.T) {
	channel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Daewu Hus Bintara Putra"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(3 * time.Second)
	close(channel)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke "+strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

func GiveMeResponse(channel chan string) {
	time.Sleep(10 * time.Microsecond)
	channel <- "Daewu Gus BP"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}