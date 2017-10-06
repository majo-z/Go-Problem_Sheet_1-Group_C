//Description: Guessing game program, where the user must guess a randomly generated number.
//Author: Marian Ziacik
//Date: 6-oct 2017
//Inspired by http://golangcookbook.blogspot.ie/2012/12/guess-number-game-v2.html
//Inspired by https://github.com/cwchentw/number-guess-golang/blob/master/main.go

package main

import (
	"fmt"
	"math/rand" //for generating random numbers
	"time"
)

const min, max = 1, 5 //range of numbers

func main() {
	//each new game, seed makes sure that different number is generated based on time shift
	rand.Seed(time.Now().Unix())

	myrand := rand.Intn(max-min) + min //generates random number
	var attempts int
	var guess int

	fmt.Println("Welcome to Guess My Number Game!")
	for guess != myrand {

		//read user input
		fmt.Printf("Take a guess between %d and %d:\n", min, max)
		fmt.Scan(&guess)

		attempts++ //count inputs
		if guess < min || guess > max {
			fmt.Println("Out of range, try again")
			attempts-- //don't count input when out of range
		} else if guess > myrand {
			fmt.Println("Too high")
		} else if guess < myrand {
			fmt.Println("Too low")
		}
	} //for
	//print number of attempts
	fmt.Printf("Right Answer! You guessed it in %v attempts", attempts)
} //main
