package main

import (
	"fmt"
	"strings"
)

func main() {
	//Broken String
	var brokenString string = "G# r#cks!"
	fmt.Println(brokenString)

	//Replace
	var fixedString string = strings.Replace(brokenString, "#", "o", -1) // -1 replaces all instances
	fmt.Println(fixedString)

	//Another example
	var anotherBrokenString string = "RJses are red, viJlets are blTe, GJphers are awesJme, and sJ are yJT!"
	fmt.Println(anotherBrokenString)
	//Fix U and J
	var anotherFixedString string = strings.Replace(anotherBrokenString, "J", "o", -1)
	anotherFixedString = strings.Replace(anotherFixedString, "T", "U", -1)
	fmt.Println(anotherFixedString)
}
