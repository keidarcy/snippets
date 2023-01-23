package main

import "fmt"

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Find(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Find(ss, "hello"))
}


func Find[T comparable] (arr []T, x T) bool {
    for _, item := range arr {
        if item == x {
            return true
        }
    }
    return false
}
