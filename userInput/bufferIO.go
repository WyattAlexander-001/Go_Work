package main

import (
	"bufio" // For buffer I/O
	"fmt"   // For printing
	"os"    // For reading input
)

func main() {
	// Create a new reader
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)

	// Read the input from the user
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from input:", err)
		return
	}
	fmt.Println(name)
}
