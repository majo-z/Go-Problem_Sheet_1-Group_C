package main

import (
	"fmt"
)

func main() {

	outputFizz := "Fizz"
	outputBuzz := "Buzz"

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(outputFizz)
		} else if i%5 == 0 {
			fmt.Println(outputBuzz)
		} else {
			fmt.Println(i)
		}
	} //for
} //main
