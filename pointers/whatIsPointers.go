package main

import "fmt"

func main() {
	var x int = 50
	fmt.Println(" X is: ")
	fmt.Println(x) // 50

	// A pointer is a variable that stores the memory address of another variable.
	var y *int = &x // y is a pointer to an int
	*y = 100        // x is now 100
	fmt.Println(" X is: ")
	fmt.Println(x)

	fmt.Println(" Y is: ")
	fmt.Println(y)  // memory address of x
	fmt.Println(*y) // 100 because y is a pointer to x

	// Pointers save memory because they allow multiple functions to share and modify the same data.
	// We don't need to pass the data to each function, we can pass the pointer
	// No need to copy the data, just pass the pointer
}
