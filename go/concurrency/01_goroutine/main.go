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

	students := []string{"John", "Paul", "George", "Ringo"}
	for _, student := range students {
		go talk(student)
	}
	time.Sleep(time.Second)
}

func talk(student string) {
	fmt.Printf("Hello, %s\n", student)
	time.Sleep(time.Second)
}
