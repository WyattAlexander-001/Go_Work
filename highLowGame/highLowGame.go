package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Print("Make a guess dude, you're on guess #: ", guesses+1)
		fmt.Println("")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Your guess is too low!")
		} else if guess > target {
			fmt.Println("Your guess is too high!")
		} else {
			fmt.Println("You got it!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess the number. It was:", target)
	}
}
