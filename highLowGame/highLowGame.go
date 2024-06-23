package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the High/Low game! The time is!:" + time.Now().String())
	fmt.Println("Please think of a number between 1 and 100.")
	target := rand.Intn(100) + 1
	fmt.Println("Target:", target)

}
