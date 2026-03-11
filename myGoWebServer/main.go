package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", cwd)

	log.Println("In Main App")
}
