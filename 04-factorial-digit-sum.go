//Adapted from https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion
package main

import "fmt"
import "math/big"

func main() {

	num := big.NewInt(100)
	fmt.Println(factorial(num))

	/* result := big.NewInt(0)
	temp := big.NewInt(0)

	for i := num; i.Cmp(num) > 0; {
		temp = num % big.NewInt(10)
		result += temp * temp
		number /= 10

	} */

} //main

//another way
/* func factorial(n *big.Int) (result *big.Int) {

	b := big.NewInt(0)
	c := big.NewInt(1)

	if n.Cmp(b) == -1 {
		result = big.NewInt(1)

	} else if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else {
		// return n * factorial(n - 1);
		//fmt.Println("n = ", n)
		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, c)))
	}
	return
} //func factorial */

func factorial(n *big.Int) *big.Int {

	b := big.NewInt(0)
	c := big.NewInt(1)
	result := new(big.Int)

	if n.Cmp(b) == -1 {
		result = big.NewInt(1)

	} else if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else {
		// return n * factorial(n - 1);
		//fmt.Println("n = ", n)

		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, c)))
	}
	return result
}
