package main

import "fmt"

func main() {
	myCar := struct { //Type unique to this instance, not common
		brand string
		model string
		year  int
	}{
		brand: "Ford",
		model: "Fiesta",
		year:  2005,
	}

	fmt.Println(myCar)

	type car struct {
		make    string
		model   string
		doors   int
		mileage int
		// wheel is a field containing an anonymous struct
		wheel struct {
			radius   int
			material string
		}
	}

	myCar2 := car{
		make:    "Ford",
		model:   "Fiesta",
		doors:   4,
		mileage: 23000,
		wheel: struct {
			radius   int
			material string
		}{
			radius:   18,
			material: "rubber",
		},
	}

	fmt.Println(myCar2)

}
