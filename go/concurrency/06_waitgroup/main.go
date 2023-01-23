package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go talk("Hello", &wg)
// 	wg.Wait()
// 	fmt.Println("World")
// }

func talk(message string, wg *sync.WaitGroup) {
	fmt.Println(message)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	students := []string{"A", "B", "C", "D", "E"}
	wg.Add(len(students))
	for _, student := range students {
		go talk(student, &wg)
	}
	wg.Wait()
	fmt.Println("Completed")
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	wg.Done()
// 	wg.Done()
// 	wg.Wait()
// }
