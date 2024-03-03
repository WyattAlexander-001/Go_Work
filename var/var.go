package main

import "fmt"

func main() {

	var a = "PLEASE DON'T USE a AS A VARIABLE NAME!"
	fmt.Println(a)

	var b, c int = 1, 2 //GO infers the type of initialized variables
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int //GO assigns a default value to uninitialized variables
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
