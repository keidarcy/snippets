package main

import (
	"fmt"
)

func main() {
	// channel := make(chan string, 1)
	// go func() {
	// 	channel <- "From inside the goroutine"
	// }()
	// channel <- "First message"
	// fmt.Println(<-channel)

	// channel := make(chan string, 1)
	// channel <- "First message"
	// fmt.Println(<-channel)
	// channel <- "Second message"
	// fmt.Println(<-channel)

	channel := make(chan string, 2)
	channel <- "First message"
	channel <- "Second message"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
