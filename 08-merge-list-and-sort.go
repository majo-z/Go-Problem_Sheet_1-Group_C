//Description: Program merges two sorted lists into a new sorted list, e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].
//Varible type: floats
//Author: Marian Ziacik
//Date: 6-oct 2017

//Go Programming Language Documents:
//Slices: usage and internals https://blog.golang.org/go-slices-usage-and-internals
//Appending to a slice https://tour.golang.org/moretypes/15
//Sort slices https://golang.org/pkg/sort/
package main

import (
	"fmt"
	"sort"
)

//import "bufio"
//import "os"

func main() {
	//Declare slices
	slice1 := []float64{3.2, 6.3, 54.3}
	slice2 := []float64{1.2, 4.5, 23.0}
	//print slices
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)

	fmt.Println()

	slice3 := make([]float64, 0) //create new empty slice
	//slice3 = append(slice3, slice1...)//append slice1
	//slice3 = append(slice3, slice2...)//append slice2
	slice3 = append(slice3, append(slice1, slice2...)...) //shorthand for appending slice1 and slice2...append can take any number of values
	fmt.Println("Slice3", slice3)                         //print the slice3(slice1 and slice2 are appended)

	sort.Float64s(slice3)                //sorts in-place
	fmt.Println("Slice3 Sorted", slice3) //print the sorted slice

	fmt.Println()
	//======================================================================

	//==================== MORE EXAMPLES ===================================
	//var slice4 []float64 = make([]float64, 5) //create slice of 5 elements, initialized to 0
	//slice4 := make([]float64, 5) //shorthand
	slice4 := make([]float64, len(slice1)+len(slice2)) //create slice of length of slice1 and slice2
	fmt.Println(slice4)

	/*
	   scanner := bufio.NewScanner(os.Stdin)
	   fmt.Println("Enter slice 1 numbers.")
	   scanner.Scan()
	   slice1String := scanner.Text()
	   fmt.Println("Enter slice 2 numbers")
	   scanner.Scan()
	   slice2String := scanner.Text()
	   fmt.Println(slice1String)
	   fmt.Println(slice2String)
	*/
}
