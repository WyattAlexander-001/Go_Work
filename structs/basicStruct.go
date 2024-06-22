package main

import "fmt"

func main() {
	type person struct {
		name       string
		age        int
		gender     string
		eyeColor   string
		weightLb   float32
		heightInch float32
		email      string
		phone      string
		address    string
	}

	var p1 person
	p1.name = "John"
	p1.age = 25
	p1.address = "123 Main St"
	p1.email = "email@gmail.com"
	p1.eyeColor = "blue"
	p1.phone = "123-456-7890"
	p1.weightLb = 160
	p1.heightInch = 72
	p1.gender = "male"

	fmt.Println(p1)
	fmt.Println("Hi your name is", p1.name, "and you are", p1.age, "years old.")

	var p2 person = person{} // longhand
	p3 := person{}           // shorthand
	fmt.Println(p2)
	fmt.Println(p3)
}
