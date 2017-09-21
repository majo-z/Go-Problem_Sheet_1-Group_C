//Hello, world! in Go
//Adapted from https://tour.golang.org/welcome/1
//Author: Marian Ziacik
//Date: 19-sep 2017

package main

import "fmt"

func main() {
	fmt.Println("こんにちは、世界！")

//Another way in english
	string_1 := "Hello"
	string_2 := "World"
	fmt.Println(string_1 + ", " + string_2)
}