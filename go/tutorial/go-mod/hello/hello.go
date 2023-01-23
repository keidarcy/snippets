package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Mary", "Jane", "Jack"}

	messages, err := greetings.Hellos(names)

	// message, err := greetings.Hello("Gla")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
