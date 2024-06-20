package main

import "fmt"

func main() { //I shouldn't have two mains in the same package
	var Cat string = "Meow"
	var Dog string = "Woof"
	fmt.Println("Cat is: ", Cat)
	fmt.Println("Dog is: ", Dog)

	fmt.Println("Swapping Cat and Dog")
	var temp string = "blank"

	temp = Cat // Temp is cheap trick to swap values
	Cat = Dog
	Dog = temp
	fmt.Println("Cat is: ", Cat)
	fmt.Println("Dog is: ", Dog)

	var Tiger string = "Roar"
	var Wolf string = "Howl"
	fmt.Println("Tiger is: ", Tiger)
	fmt.Println("Wolf is: ", Wolf)

	fmt.Println("Swapping Tiger and Wolf")
	Tiger, Wolf = Wolf, Tiger // This is the correct way to swap values
	fmt.Println("Tiger is: ", Tiger)
	fmt.Println("Wolf is: ", Wolf)

}
