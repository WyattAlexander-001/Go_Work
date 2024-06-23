package main

import (
	"bufio" // For buffer I/O
	"fmt"   // For printing
	"os"    // For reading input
)

func main() {
	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin)
	grade, _ := reader.ReadString('\n')
	fmt.Println(grade + " is a great grade!")
}
