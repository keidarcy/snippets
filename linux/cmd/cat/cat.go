package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// data, err := io.ReadAll(os.Stdin)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if len(data) == 0 {
	// 	fmt.Println("No input received, exiting...")
	// 	return
	// }
	// fmt.Print(string(data))

	paths := os.Args[1:]

	if len(paths) == 0 {
		fmt.Println("Usage: cat file1 [file2 file3...]")
	}

	var ln int
	for _, path := range paths {
		// cat(path)
		ln += count(path)
	}

	println(ln)
	os.Exit(0)

}

func cat(path string) {
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

	fmt.Print(string(contents))
}

func count(path string) (res int) {
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
	return
}
