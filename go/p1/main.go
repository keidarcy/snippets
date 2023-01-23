package main

import "example.com/at/internal/server"

func main() {
	println("Starting")
	s := server.NewHTTPServer(":8080")
	s.ListenAndServe()
}
