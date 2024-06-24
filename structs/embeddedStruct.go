package main

import "fmt"

type car struct {
	make  string
	model string
	year  int
}

type truck struct {
	car
	bedSize   int
	fourWheel bool
}

func main() {
	myTruck := truck{
		car: car{
			make:  "Ford",
			model: "F150",
			year:  2019,
		},
		bedSize:   8,
		fourWheel: true,
	}

	fmt.Println(myTruck)
}
