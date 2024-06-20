package main

import "fmt"

func main() {
	x := 5
	increment(x)

	fmt.Println(x)            // still prints 5, not 6 because the function does not modify the value of x because it is passed by value
	fmt.Println(increment(x)) // prints 6 because the function returns the incremented value
}

func increment(x int) int {
	x++
	return x
}
