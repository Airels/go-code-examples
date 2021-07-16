package main

import (
	"fmt"
	"os"
)

func main() {
	var username string = os.Getenv("USER");

	if username == "" {
		username = os.Getenv("USERNAME")
	}

	fmt.Println("Hello, world!")
	fmt.Println(username, ", let's be friends!")
}
