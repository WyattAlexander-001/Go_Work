package main

import "fmt"

func Hello() string {
	return "Hello, world" // Will pass the test "Hello, world
	// return "Hello, world!!!!" // Will fail
}

func HelloName(name string) string {
	return "Hello, " + name
}

func IdealHello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "English" {
		return "Hello, " + name
	} else if language == "Spanish" {
		return "Hola, " + name
	}
	return "Hi, " + name
}

func main() {
	fmt.Println(Hello())
	fmt.Println("Yay it works!")
	fmt.Println(HelloName("Wyatt"))
	fmt.Println(IdealHello("Wyatt", "Spanish"))
	fmt.Println(IdealHello("", "Spanish"))
	fmt.Println(IdealHello("", ""))
}
