package main

import "fmt"

func main() {
	var myString string

	fmt.Printf("Please enter your string:\n")
	fmt.Scanf("%s", myString)
	fmt.Println(myString)
	fmt.Println(ReverseString(myString))

}

//function for string reverse
func ReverseString(str string) string {
	var reverse string

	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse
}
