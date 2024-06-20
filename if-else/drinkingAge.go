package main

import "fmt"

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	//Naked return would be like this but it is not recommended
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func main() {
	age := 20
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("Years until you can vote:", yearsUntilAdult)
	fmt.Println("Years until you can drink:", yearsUntilDrinking)
	fmt.Println("Years until you can rent a car:", yearsUntilCarRental)
}
