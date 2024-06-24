package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("You thought I was gonna say 4 didn't you?")

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("Lift off!")
}
