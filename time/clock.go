package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("The time is", now)
	year := now.Year()
	fmt.Println("The year is", year)
}
