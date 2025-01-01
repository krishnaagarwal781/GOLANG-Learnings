package main

import "fmt"

func main() {
	const name string = "Krishna Agarwal"

	fmt.Println("My name is", name)

	// const grouping

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Print("Server running on ", host, " On port ", port)
}
