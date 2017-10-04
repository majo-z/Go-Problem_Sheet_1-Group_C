package main

import (
	"fmt"
)

func main() {

	outputNumbers := ""
	//outputFizz := "Fizz"
	//outputBuzz := "Buzz"

	for i := 1; i <= 100; i++ {
		if outputNumbers == "" {
			fmt.Println(i)
		}
	}
}
