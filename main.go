package main

import "fmt"

func main() {
	StartLogger()
	fmt.Println("Application started, listening on port 4000")
	StartListener()
}
