package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert a string to a float64
	grade, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println("Error converting string to number:", err)
		return
	}
	fmt.Println(grade)
}
