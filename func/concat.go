package main

import "fmt"

func concat(s1 string, s2 string) string { // Returns a string
	return s1 + s2
}

func test(s1, s2 string) { // Void function
	fmt.Println(concat(s1, s2))
}

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")

	fmt.Println("Concatenation is the best!")
	fmt.Println(concat("I love", " Go"))
}
