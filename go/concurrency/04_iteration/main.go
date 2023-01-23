package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	count := 6
	go talk(channel, count)
	for i := 0; i < count; i++ {
		println(i, "before receive")
		value := <-channel
		println(value)
		println(i, "after receive")
		// fmt.Println(<-channel)
	}
}

func talk(cha chan string, count int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		score := rand.Intn(10)
		println(i, "before send")
		cha <- fmt.Sprintf("%d Score", score)
		println(i, "after send")
	}
}

// func main() {
// 	cha := make(chan string)
// 	go talk(cha)
// 	// for message := range cha {
// 	// fmt.Println(message)
// 	// }
// 	for {
// 		message, open := <-cha
// 		if !open {
// 			break
// 		}
// 		fmt.Println(message)
// 	}
// }
//
// func talk(cha chan string) {
// 	rand.Seed(time.Now().UnixNano())
// 	count := 3
// 	for i := 0; i < count; i++ {
// 		score := rand.Intn(10)
// 		cha <- fmt.Sprintf("Score %d", score)
// 	}
// 	close(cha)
// }
