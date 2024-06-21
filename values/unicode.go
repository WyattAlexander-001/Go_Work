package main

import "fmt"

func main() {
	fmt.Println('A')
	fmt.Println('B')
	fmt.Println('A' + 'B') // 131 because 65 + 66
	//Print unicode character using unicode code point
	//A is 65 in unicode
	fmt.Println(string(65))
	//B is 66 in unicode
	fmt.Println(string(66))
	fmt.Println(string(65) + string(66))
	fmt.Println(string(1174)) // Җ
}
