//Description: Program calculates the sum of the digits of the factorial of a number 100
//Author: Marian Ziacik
//Date: 5-oct 2017

//The Go Programming Language Documents:
//MulRange: //https://golang.org/pkg/math/big/#Int.MulRange
//strings: https://golang.org/pkg/strings/#Split
//strconv: https://golang.org/pkg/strconv/

package main

import (
	"fmt"
	"math/big"
	"strconv" //package for converting strings (used for ints here)
	"strings" //package for handling strings, used for strings.Split
)

func main() {

	num := big.NewInt(100) // *Int
	//fmt.Println(factorial(num))

	num.MulRange(1, 100)                                         //same as 100 x 99 x 98 ... x 1
	fmt.Printf("Factorial of %d is: %d\n", big.NewInt(100), num) //print factorial of 100

	//split to get every character in a slice ["1", "2", "3" ...]
	//var digitsStr []string = strings.Split(num.String(), "")
	digitsStr := strings.Split(num.String(), "") //shorthand

	total := 0 //for sum up individual digits
	for _, value := range digitsStr {
		//error can be ignored, because the input is always valid
		intValue, _ := strconv.Atoi(value) // Parse value to integer ex. "2" : 2
		total += intValue
	}

	fmt.Printf("Sum up of individual numbers of factorial %v is: %d\n", big.NewInt(100), total)

	/* //example of strconv.Atoi()
	i, _ := strconv.Atoi("42")
	fmt.Println(i) */
} //main

/*
//Another way
//Adapted from https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion
func factorial(n *big.Int) *big.Int {

    b := big.NewInt(0)
    c := big.NewInt(1)
    result := new(big.Int)

    if n.Cmp(b) == -1 {
        result = big.NewInt(1)

    } else if n.Cmp(b) == 0 {
        result = big.NewInt(1)
    } else {
        // return n * factorial(n - 1);
        //fmt.Println("n = ", n)

        result.Set(n)
        result.Mul(result, factorial(n.Sub(n, c)))
    }
    return result
}
*/
