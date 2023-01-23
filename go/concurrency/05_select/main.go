package main

import (
	"fmt"
	"time"
)

func main() {
	chan1, chan2 := make(chan string), make(chan string)
	go talk(chan1, "Hello")
	go talk(chan2, "World")

	select {
	case message := <-chan1:
		fmt.Println(message)
	case message := <-chan2:
		fmt.Println(message)
		// default:
		// 	fmt.Println("No message received")
	}
	// fmt.Println(<-chan1, <-chan2)
	roughlyFair()
}

func talk(cha chan string, message string) {
	time.Sleep(time.Second * 2)
	cha <- message
}

func roughlyFair() {
	chan1 := make(chan interface{})
	close(chan1)
	chan2 := make(chan interface{})
	close(chan2)

	var count1, count2 int
	for i := 0; i < 1000; i++ {
		select {
		case <-chan1:
			count1++
		case <-chan2:
			count2++
		}
	}

	fmt.Printf("count1: %d, count2: %d", count1, count2)
}
