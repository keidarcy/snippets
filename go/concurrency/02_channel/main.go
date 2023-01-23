package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	helloChan := make(chan bool)

	go talk("John", helloChan)
	// helloChan <- false
	<-helloChan
}

func talk(student string, helloChan chan bool) {
	fmt.Printf("Hello, %s\n", student)
	time.Sleep(time.Second)
	helloChan <- true
}
