package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	ready         bool
	workIntervals int
)

func main() {
	// getReady()
	getReadyWithCond()
	broadcastReady()
}

func broadcastReady() {
	bc := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)
	standByForReady(func() {
		fmt.Println("Worker 1 is ready.")
		wg.Done()
	}, bc)
	standByForReady(func() {
		fmt.Println("Worker 2 is ready.")
		wg.Done()
	}, bc)
	standByForReady(func() {
		fmt.Println("Worker 3 is ready.")
		wg.Done()
	}, bc)
	bc.Broadcast()
	wg.Wait()
	fmt.Println("All workers are ready.")
}

func standByForReady(fn func(), bc *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		bc.L.Lock()
		defer bc.L.Unlock()
		bc.Wait()
		fn()
	}()
	wg.Wait()
}

func getReadyWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)

	cond.L.Lock()
	for !ready {
		workIntervals++
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Printf("We are now ready! After %d work intervals.\n", workIntervals)
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()
}

func getReady() {
	go gettingReady()

	for !ready {
		// time.Sleep(5 * time.Second)
		workIntervals++
	}
	fmt.Printf("We are now ready! After %d work intervals.\n", workIntervals)
}

func gettingReady() {
	sleep()
	ready = true
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
