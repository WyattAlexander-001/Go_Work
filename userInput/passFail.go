package main

import (
	"bufio"   // For buffer I/O
	"fmt"     // For printing
	"log"     // For logging
	"os"      // For reading input
	"strconv" // For string conversion
	"strings" // For string manipulation
)

func main() {
	// Create a new reader
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)

	// Read the input from the user
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading from input:", err)
		return
	}
	fmt.Println(name)

	// Convert a string to a float64
	fmt.Print("Enter your grade: ")
	grade, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading from input:", err)
		return
	}

	grade = strings.TrimSpace(grade)
	gradeNum, err := strconv.ParseFloat(grade, 64)
	if err != nil {
		log.Println("Error converting string to number:", err)
		return
	}
	if gradeNum >= 70 {
		fmt.Println("You passed!")
	} else {
		fmt.Println("You failed!")
	}
}
