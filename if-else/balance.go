package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// Conditional calculation of the final cost with discount for premium users
	if isPremiumUser {
		finalCost = bulkMessageCost * (1 - discountRate)
	} else {
		finalCost = bulkMessageCost
	}

	// Check if the account has sufficient balance
	if accountBalance >= finalCost {
		accountBalance -= finalCost // Deduct the final cost from the account balance
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	msg := fmt.Sprintf("Account balance: $%.2f", accountBalance)
	fmt.Println(msg)
}
