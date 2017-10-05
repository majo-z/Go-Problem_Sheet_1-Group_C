//Description: Newton’s method for calculating square roots
//Newton’s method is to approximate the square root of a number x by picking a starting point z and then repeating the following operation.
//zNext = z - ((z*z - x) / (2 * z))
//Author: Marian Ziacik
//Date: 4-oct 2017
//Adapted from Ian McLoughlin

package main

import (
	"fmt"
	//"log" //used for error log message
	"math"
)

func zNext(z float64, x float64) float64 { //z float64, x float64...returns float64

	return z - ((z*z - x) / (2 * z))
}

func main() {

	//the number to find square root of
	//x := 64.0
	var x float64

	//the initial guess
	//z := float64(1) //or...z := 1.0
	var z float64
	var startingGuess float64
	i := 0 //guess counter

	//scan user input for square root
	fmt.Printf("Please enter the number to find square root of: ")
	fmt.Scanf("%f\n", &x)

	//scan user input for initial guess
	fmt.Printf("Please enter your guess: ")
	fmt.Scanf("%f\n", &startingGuess)
	//fmt.Scan(&startingGuess)
	fmt.Println()

	//Another way (Adapted from: https://stackoverflow.com/questions/3751429/reading-an-integer-from-standard-input/3751456)
	/* if _, err := fmt.Scan(&startingGuess); err != nil {
		log.Print("  Scan failed, due to ", err)
		return
	} */

	//Iterate until the next is the same as current guess
	//This is repeated while the values of zNext and z are different.
	//After each iteration the value of z is replaced with that of zNext
	for z = startingGuess; z != zNext(z, x); z = zNext(z, x) {

		i++ //current guess counter

		//print the guess on each iteration
		fmt.Printf("Current guess %d: %13f\n", i, z)

	}
	//finally, z is a good appriximation of th esquare root of x
	fmt.Printf("The square root of %f is %f\n", x, z)

	//print out the math.Sqrt value
	fmt.Printf("math.Sqrt gives the value of square root %f as %f\n", x, math.Sqrt(x))

	//Total number of guesses
	if z == startingGuess {
		fmt.Println("Congratulations, you've got it righ on first attempt!")
	} else {
		fmt.Printf("Total number of guesses: %d\n", i)
	}

}
