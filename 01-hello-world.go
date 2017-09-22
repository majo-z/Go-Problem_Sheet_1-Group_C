//Description: Program prints “Hello, world!” in Japanese (using Japanese characters) to the screen
//Author: Marian Ziacik
//Date: 19-sep 2017
//Adapted from https://tour.golang.org/welcome/1

// This file belongs to the package main - this means it will become an executable
package main

// The fmt package is used to print to the console, in this case
import "fmt"

func main() {
	fmt.Println("こんにちは、世界！")

	//Another way in english using concatenating
	string_1 := "Hello"
	string_2 := "World!"
	fmt.Println(string_1 + ", " + string_2)

	//Different ways
	var message string //message := "Hello, World"
	message = "Hello, World"

	fmt.Println(message)

	//Printing "Hello, World" using Printf function - string formatting
	fmt.Printf("Hello, World\n")
}
