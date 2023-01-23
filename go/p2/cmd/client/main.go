package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"example.com/p2/internal/client"
)

const URL = "http://localhost:8080"

func main() {
	add := flag.Bool("add", false, "Add activity")
	get := flag.Bool("get", false, "Get activity")
	list := flag.Bool("list", false, "Get list of activities")

	flag.Parse()
	activityClient := &client.ActivityClient{URL: URL}

	switch {
	case *get:
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, `--get only accept integer`)
			os.Exit(1)
		}
		a, err := activityClient.Retrieve(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, `http error: `, err.Error())
			os.Exit(1)
		}
		fmt.Printf("id: %v, activity: %v\n", id, *a)

	case *list:
		if len(os.Args) != 2 {
			fmt.Fprintln(os.Stderr, `Usage: --list`)
			os.Exit(1)
		}
		list, err := activityClient.List()

		if err != nil {
			log.Fatal(err)
		}

		for id, a := range *list {
			fmt.Printf("id: %v, activity: %v\n", id, a)
		}

	case *add:
		if len(os.Args) != 3 {
			fmt.Fprintln(os.Stderr, `Usage: --add "message"`)
			os.Exit(1)
		}
		a := client.Activity{
			Time:        time.Now(),
			Description: os.Args[2],
		}
		id, err := activityClient.Insert(a)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("id: %v, activity: %v\n", id, a)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
