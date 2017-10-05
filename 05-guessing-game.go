//Inspired by http://golangcookbook.blogspot.ie/2012/12/guess-number-game-v2.html
//Inspired by https://github.com/cwchentw/number-guess-golang/blob/master/main.go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	const min, max = 1, 5

	myrand := rand.Intn(max-min) + min
	var attempts int
	var guess int
	var isCorrect bool

	fmt.Println("Welcome to Guess My Number Game!")
	for guess != myrand {

		fmt.Printf("Take a guess between %d and %d:\n", min, max)
		fmt.Scan(&guess)

		attempts++
		if guess < min || guess > max {
			fmt.Println("Out of range, try again")
			attempts--
		} else if guess > myrand {
			fmt.Println("Too high")
		} else if guess < myrand {
			fmt.Println("Too low")
		} else {
			isCorrect = true
			break
		}
	} //for
	if isCorrect {
		fmt.Printf("Right Answer! You guessed it in %v tries", attempts)
	}
} //main
