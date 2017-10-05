//Description: Write a function to reverse a string in Go.
//Author: Marian Ziacik
//Date: 5-oct 2017
//Adapted from https://coderwall.com/p/fw6miw/reverse-text-in-golang
package main

import "fmt"

func main() {
	var myString string

	//ask and read user input
	fmt.Printf("Please enter your string: ")
	fmt.Scanf("%s", &myString)
	//print the entered string, print the reversed string
	fmt.Printf("Entered string: %s\n", myString)
	fmt.Printf("Reversed string: %s\n", ReverseString(myString))
}

//string reverse function
func ReverseString(str string) string {
	var reverse string
	//get the length of the string
	for i := len(str) - 1; i >= 0; i-- {
		//add char from last position to first
		reverse += string(str[i])
	}
	return reverse
}
