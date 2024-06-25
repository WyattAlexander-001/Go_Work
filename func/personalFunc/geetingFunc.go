package main

import (
	"fmt"
)

func greet() {
	fmt.Println("Hello!")
}

// Go does not allow for function overloading:
// https://stackoverflow.com/questions/6986944/does-the-go-language-have-function-meth
// func Greet(name string) {
// 	fmt.Println("Hello", name)
// }

func greetByName(name string) {
	fmt.Println("Hello", name)
}

func greetByNameXTimes(name string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println("Hello", name)
	}
}

func main() {
	greet()
	greetByName("Gopher")
	greetByNameXTimes("Jimmy", 5)
}
