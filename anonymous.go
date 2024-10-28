// anonymous functions

package main

import (
	"fmt"
	"time"
)

func main() {

	// anonymous function, the function has no name can be called by the variable assigned
	var greet = func() {
		fmt.Println("Greetings from anonymous function..")
	}

	// another anonymous function
	var hello = func() {
		c := time.Now()
		fmt.Printf("Hello Today is %v-%v-%v\n", c.Month(), c.Day(), c.Year())
	}
	// function call
	greet()
	hello()

}
