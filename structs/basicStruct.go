package main

import (
	"fmt"
)

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
		race       string
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
	p1.race = "white"

	fmt.Println(p1)
	fmt.Println("Hi your name is", p1.name, "and you are", p1.age, "years old.")

	var p2 person = person{} // longhand
	p3 := person{}           // shorthand
	fmt.Println(p2)
	fmt.Println(p3)
	p3.name = "Latisha"
	p3.age = 30
	p3.address = "456 Elm St"
	p3.race = "black"
	p3.gender = "female"
	p3.eyeColor = "brown"
	fmt.Println(p3)

	//Car example
	type wheel struct {
		radius   int
		material string
	}

	type car struct {
		make       string
		model      string
		doors      int
		mileage    int
		frontWheel wheel
		backWheel  wheel
	}

	myCar := car{}
	myCar.frontWheel.radius = 5
	fmt.Println(myCar)
}
