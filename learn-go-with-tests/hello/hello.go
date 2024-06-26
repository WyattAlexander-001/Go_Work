package main

import "fmt"

func Hello() string {
	return "Hello, world" // Will pass the test "Hello, world
	// return "Hello, world!!!!" // Will fail
}

func main() {
	fmt.Println(Hello())
	fmt.Println("Yay it works!")
}
