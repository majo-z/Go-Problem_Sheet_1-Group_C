package main

import "fmt"
import "sort"

//import "bufio"
//import "os"

func main() {
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

	slice2 := []float64{1.2, 4.5, 23.0}
	slice1 := []float64{3.2, 6.3, 54.3}
	fmt.Println(slice1)
	fmt.Println(slice2)
	// strings.Split(s, " ") // if the floats are entered like "54.34 654.4 343.54"
	// strconv.ParseFloat(s, 64) // to get float64

	// slice3 := make([]float64, len(slice1) + len(slice2))
	var slice3 []float64 = make([]float64, 0)
	fmt.Println(slice3)

	// https://tour.golang.org/moretypes/15
	//slice3 = append(slice3, append(slice1, slice2...)...) // append can take any number of values.
	slice3 = append(slice3, slice1...)
	slice3 = append(slice3, slice2...)

	fmt.Println(slice3)
	// https://golang.org/pkg/sort/
	sort.Float64s(slice3) // sorts in-place
	fmt.Println(slice3)
	//function(1,2,3,4,5,6,7,98, 32,43,54,65,34)
}

/*
func function(numbers ...int){
    for _, num := range numbers{
        fmt.Println(num)
    }
}*/
