package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func main() {
	//_main()
	readWrite()
}

func readWrite() {
	go read()
	go read()
	go read()
	go write()

	time.Sleep(time.Second * 5)
	fmt.Println("DONE")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("read locking")
	time.Sleep(time.Second * 1)
	fmt.Println("read unlocking")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("write locking")
	time.Sleep(time.Second * 1)
	fmt.Println("write unlocking")
}

func _main() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}

	time.Sleep(time.Second * 5)
	fmt.Println("hello"+"world: ", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}
