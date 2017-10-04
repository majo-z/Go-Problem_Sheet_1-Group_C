//Description: Program prints the numbers from 1 to 100,
//For multiples of three print Fizz instead of the number, and for multiples of five print Buzz.
//For numbers that are multiples of both three and five print FizzBuzz.

//Author: Marian Ziacik
//Date: 4-oct 2017
//Insired by video from https://www.youtube.com/watch?v=QPZ0pIK_wsc

package main

import (
	"fmt"
)

func main() {

	outputFizz := "Fizz"
	outputBuzz := "Buzz"

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(outputFizz + outputBuzz) //multiples of three and five print FizzBuzz
		} else if i%3 == 0 {
			fmt.Println(outputFizz) //multiples of three print Fizz
		} else if i%5 == 0 {
			fmt.Println(outputBuzz) //multiples of five print Buzz
		} else {
			fmt.Println(i) //otherwise print number
		}
	} //for
} //main
