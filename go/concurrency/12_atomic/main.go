package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Print("hello", "world")
	fmt.Println("hello", "world")
	fmt.Fprint(os.Stdout, "hello")
	var sum int64
	fmt.Println(sum)
	atomic.AddInt64(&sum, 1)
	fmt.Println(sum)

	var mu sync.Mutex
	mu.Lock()
	sum++
	mu.Unlock()
	fmt.Println(sum)

	var diffSum int64
	fmt.Println(atomic.LoadInt64(&diffSum))
	atomic.StoreInt64(&diffSum, 1)
	fmt.Println(diffSum)

	var v atomic.Value
	list := A{"Johnny"}
	v.Store(list)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		w := v.Load().(A)
		w.name = "NOT Johnny"
		v.Store(w)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(v.Load().(A).name)
}

type A struct {
	name string
}
