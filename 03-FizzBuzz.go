package main

import (
	"fmt"
)

func main() {

	var output string = ""

	for i := 1; i <= 100; i++ {
		if output == "" {
			fmt.Println(i)
		}
	}
}
