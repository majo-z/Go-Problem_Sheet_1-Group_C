//Description: Program prints the current time and date to the console
//Author: Marian Ziacik
//Date: 21-sep 2017
//Adapted from https://tour.golang.org/welcome/4

package main

import (
	"fmt"
	"time" // The time package provides the current time and data
)

func main() {
	fmt.Println("The time and date are: ", time.Now())

	//Formatting & parsing options
	//Adapted from https://gobyexample.com/time-formatting-parsing

	//substituting printLine for fmt.Println
	//now the printLine can be used intead of fmt.Println
	printLine := fmt.Println

	//substituting t for time.Now()
	t := time.Now()
	printLine(t.Format("The time and date are: " + time.RFC3339))
	printLine(t.Format("Date: Mon Jan _2 2006, Time: 15:04:05 "))

}
