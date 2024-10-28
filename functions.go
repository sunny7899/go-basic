// go program with functions

package main

import "fmt"

// function with two variables
func sum1(a int, b int) int {
	return a + b
}

// function with Variadic  list of integers
func sum2(a int, b ...int) int {
	var s int = 0

	for _, i := range b {
		s += i
	}
	return s

}

// function with 3 fixed value integers
func sum3(a int, b [3]int) int {
	var s int = a
	var i int = 0

	for i = 0; i < len(b); i++ {
		s += b[i]
	}
	return s
}

func main() {

	var a int = 10                // single variable
	var b int = 20                // single variable
	var c = []int{2, 3, 12, 1, 6} // Variadic array
	var d = [3]int{2, 3, 12}      // fixed 3 number array

	fmt.Printf("Sum of %v and %v is %v\n", a, b, sum1(a, b))
	fmt.Printf("Sum of %v and %v is %v\n", a, c, sum2(a, c...))
	fmt.Printf("Sum of %v and %v is %v\n", a, d, sum3(a, d))

}
