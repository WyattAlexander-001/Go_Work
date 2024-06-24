package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix() // Truly random number based on the current time
	//var seconds int64 = 5 //This works too, but ask why I can't just use this instead

	rand.NewSource(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Welcome to the High/Low game! The time is!:" + time.Now().String())
	fmt.Println("Please think of a number between 1 and 100.")

	fmt.Println("DEBUG, Seconds:", seconds)
	fmt.Println("DEBUG, Target:", target)

}
