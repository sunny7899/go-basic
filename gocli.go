// program for cli input, converts give list of decimal to hex number

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Program to read command line arguements..\n")
	args := os.Args

	fmt.Println("Total Arguments: ", len(os.Args)-1)

	fmt.Printf("Program name: %v\n\n", os.Args[0])

	for i := 1; i < len(args); i++ {
		fmt.Printf("Argument: %v: %v\n", i, args[i])

	}
}
