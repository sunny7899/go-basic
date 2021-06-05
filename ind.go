// main- a package is a way to group functions, and it's made up of all the files in the same directory
package main

// fmt package contains functions for formatting text,
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello!")
	fmt.Println(quote.Go())
}
