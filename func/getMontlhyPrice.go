package main

import "fmt"

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		return 100.00
	} else if tier == "premium" {
		return 150.00
	} else if tier == "enterprise" {
		return 500.00
	} else {
		return 0
	}
}

func main() {
	// Get the monthly price for the "premium" tier
	price := getMonthlyPrice("premium")
	fmt.Println("Plan cost:$", price)
}
