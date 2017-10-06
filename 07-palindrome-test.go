//Description: Program tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.
//Author: Marian Ziacik
//Date: 6-oct 2017

// Goolang Documentation: strings
//https://golang.org/src/strings/example_test.go
package main

import (
	"fmt"
	"strings" //package for handling strings, in this case used for strings.ToLower
)

func main() {
	fmt.Printf("Strings are case insensitive\n")
	fmt.Printf("========== HARD CODED TESTS ==========\n")
	pal1 := "lEvel"
	pal2 := "nooN"
	pal3 := "Hello"
	pal4 := "World"
	pal5 := "radar"

	fmt.Printf("The string %s is palindrome: %v\n", pal1, TestPalindrome(pal1))
	fmt.Printf("The string %s is palindrome: %v\n", pal2, TestPalindrome(pal2))
	fmt.Printf("The string %s is palindrome: %v\n", pal3, TestPalindrome(pal3))
	fmt.Printf("The string %s is palindrome: %v\n", pal4, TestPalindrome(pal4))
	fmt.Printf("The string %s is palindrome: %v\n", pal5, TestPalindrome(pal5))

	fmt.Println()

	fmt.Printf("========== USER INPUT TEST ==========\n")
	var userInput string
	fmt.Printf("Please enter your string:\n")
	fmt.Scanf("%s%\n", &userInput)
	fmt.Printf("The string %s is palindrome: %v\n", userInput, TestPalindrome(userInput))
}

//same funtion is used in reverse string example(interchangeable)
func ReverseString(str string) string {
	var reverse string
	//get the length of the string
	for i := len(str) - 1; i >= 0; i-- {
		//add char from last position to first
		reverse += string(str[i])
	}
	return reverse
}

//TestPalindrome function returns boolean true if palindrome, false if not
func TestPalindrome(testString string) bool {
	testString = strings.ToLower(testString)       //ignore case
	return ReverseString(testString) == testString //compare strings
}
