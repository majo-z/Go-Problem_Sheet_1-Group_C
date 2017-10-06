package main

import "fmt"

func main() {
	numbers := []int{23, 43, 54, 23, 43, 54, 65, 76, -10, 32, -343}
	fmt.Println(largestSmallest(numbers))
}

//function takes array of integers and returns smallest and largest number of the array
func largestSmallest(numbers []int) (int, int) { //enter array
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
}
