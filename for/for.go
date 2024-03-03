package main

import "fmt"

func main() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("=====^^^UGLY VERSION^^==============-")

	for j := 0; j < 3; j++ {
		fmt.Println(j, "// Well that is the imdex")
	}

	fmt.Println("=======^^Standard Version^^============-")

	for i := range 3 {
		fmt.Println("INDEX: ", i)
	}
	fmt.Println("========^^ Another version but shorter syntax ^^===========-")

	for { //infinite loop
		fmt.Println("loop")
		break
		fmt.Println("after the break I am unreachable") //this line will never be executed
	}
	fmt.Println("========^^Just a break, useful under weird corcumstance^^===========-")

	for n := range 6 { //range returns both the index and the value
		if n%2 == 0 {
			fmt.Println(n, "is even because")
			fmt.Println(n, "mod 2 is", n%2)
			continue
		}
		fmt.Println(n)
	}
}
