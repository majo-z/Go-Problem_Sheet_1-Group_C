//Current time in Go
//Adapted from https://tour.golang.org/welcome/4
//Author: Marian Ziacik
//Date: 21-sep 2017

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("The time is", time.Now())

	//Formatting options
	//Adapted from https://gobyexample.com/time-formatting-parsing
	printLine := fmt.Println

	t := time.Now()
	printLine(t.Format(time.RFC3339))

}