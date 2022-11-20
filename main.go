package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// create a secret number
	secret := getRandomNumber()
	fmt.Println(secret)
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11 //get a random number from zero to ten
}
