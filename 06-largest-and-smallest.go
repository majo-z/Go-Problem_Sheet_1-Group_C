//Description: Program that returns the largest and smallest elements in a array
//Author: Marian Ziacik
//Date: 6-oct 2017
//The Go Programming Language Documents:
// https://tour.golang.org/moretypes/6
package main

import "fmt"

func main() {
	numbers := []int{23, 43, 54, 23, 43, 54, 65, 76, -10, 32, -343}
	fmt.Println(largestSmallest(numbers))
}

//function takes array of integers and returns smallest and largest number of the array
func largestSmallest(numbers []int) (int, int) { //enter array
	//Initialize smallest and largest to 0
	smallest := numbers[0]
	largest := numbers[0]

	for _, num := range numbers { //don't need index (_)
		if num > largest {
			largest = num
		}
		if num < smallest {
			smallest = num
		}
	}
	return smallest, largest
} //largestSmallest
