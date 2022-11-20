package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// create a secret number
	secret := getRandomNumber()
	//fmt.Println(secret)

	for matching := false; !matching;{
		// get a user input
		guess := getUserInput()
		//fmt.Println(secret, guess)

		// make a comparison (secret vs guess)
		matching = isMatching(secret, guess)
		fmt.Println(matching)
	}
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11 //get a random number from zero to ten
}

func getUserInput() int {
	fmt.Println("Please enter your guess: ")
	var input int
	_, err := fmt.Scan(&input) // _number of bytes scan and err: potential error
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed:", input)
	}

	return input
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your guess is too big")
		return false
	} else if guess < secret {
		fmt.Println("Your guess is too small")
		return false
	} else {
		fmt.Println("YOU GOT IT!")
		return true
	}
}
