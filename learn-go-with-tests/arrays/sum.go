package main

import (
	"fmt"
)

func main() {
	faveNumbers := [5]int{6, 7, 777, 21, 18}
	fmt.Println(faveNumbers)

	hatedNumbers := [5]int{12, 2}
	fmt.Println(hatedNumbers) //Non-filled will default to 0
}
