package main

import "fmt"

func getPoint() (x int, y int) {
	return 3, 4 // x = 3, y = 4 they are named return values not parameters
}

// ignore y value

func main() {
	x, _ := getPoint()
	fmt.Println(x) // 3

	_, y := getPoint()
	fmt.Println(y) // 4

	fmt.Println(getPoint()) // 3, 4

}
