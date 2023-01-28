package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	stdin := make(chan []byte)
	go func() {

		data, _ := io.ReadAll(os.Stdin)
		stdin <- data
	}()
	select {
	case data := <-stdin:
		if len(os.Args) == 1 {
			fmt.Print(string(data))
		} else {
			cat()
		}
	case <-time.After(time.Millisecond * 10):
		cat()
	}
	count()
}

func cat() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		fmt.Println("Usage: cat file1 [file2 file3...]")
		os.Exit(0)
	}
	var count int
	for _, path := range paths {

		file, err := os.Open(path)

		if err != nil {
			fmt.Printf("Error opening file: %s \n", path)
			return
		}

		defer file.Close()

		contents, err := io.ReadAll(file)

		if err != nil {
			fmt.Printf("Error read file: %s \n", path)
			return
		}

		lines := strings.Split(string(contents), "\n")

		for _, line := range lines {
			fmt.Printf("%d: %s\n", count, line)
			count++
		}
	}
}

func count() {
	var res int
	paths := os.Args[1:]
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return
		}

		bytes, err := io.ReadAll(file)
		if err != nil {
			return
		}
		for _, b := range bytes {
			if b == byte('\n') {
				res++
			}
		}
	}
	fmt.Println()
	fmt.Println("line count: ", res+1)
}
