package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func power(a, b int) int {
	return a ^ b
}

func mod(a, b int) int {
	return a % b
}

func increment(a int) int {
	return a + 1
}

func decrement(a int) int {
	return a - 1
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(subtract(2, 3))
	fmt.Println(multiply(2, 3))
	fmt.Println(divide(2, 3))
	fmt.Println(power(2, 3))
	fmt.Println(mod(2, 3))
	fmt.Println(increment(2))
	fmt.Println(decrement(2))
	fmt.Println(abs(-8888))

}
