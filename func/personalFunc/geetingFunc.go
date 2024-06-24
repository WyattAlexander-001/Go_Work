package main

import (
	"fmt"
)

func Greet() {
	fmt.Println("Hello!")
}

// Go does not allow for function overloading:
// https://stackoverflow.com/questions/6986944/does-the-go-language-have-function-meth
// func Greet(name string) {
// 	fmt.Println("Hello", name)
// }

func GreetByName(name string) {
	fmt.Println("Hello", name)
}

func main() {

}
