package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var isGood bool

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			if foundGood() {
				once.Do(markGood)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	checkGood()
}

func checkGood() {
	if isGood {
		fmt.Println("It's good")
	} else {
		fmt.Println("It's not good")
	}
}

func markGood() {
	fmt.Println("marking good")
	isGood = true
}

func foundGood() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
