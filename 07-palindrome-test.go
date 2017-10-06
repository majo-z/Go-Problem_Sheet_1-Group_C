package main

import (
	"fmt"
	"strings"
)

func main() {
	pal1 := "lEvel"
	pal2 := "nooN"
	pal3 := "Hello"
	pal4 := "World"

	fmt.Println(TestPalindrome(pal1))
	fmt.Println(TestPalindrome(pal2))
	fmt.Println(TestPalindrome(pal3))
	fmt.Println(TestPalindrome(pal4))
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
