// GO program for math operators

package main

import (
	"fmt"
	"log"
)

func main() {

	// read input from key board

	var n1, n2 int
	//var n int

	fmt.Println("Program for math operators")

	again := 1

	for again != 0 {
		fmt.Print(`

		Math operation
		----------------

		1. add
		2. subtract
		3. multiply
		4. divide

		0. exit

		`)

		fmt.Print("Choose your option: ")
		_, err := fmt.Scanln(&again)

		if err != nil {
			fmt.Println("\nGot an error")
			log.Fatal(err)

		}

		if again == 0 {
			break
		}

		fmt.Print("\nEnter first integer number:  ")
		_, err1 := fmt.Scanln(&n1)

		if err1 != nil {
			fmt.Println("\nGot an error")
			log.Fatal(err)

		}
		fmt.Print("Enter second integer number: ")
		_, err2 := fmt.Scanln(&n2)

		if err2 != nil {
			fmt.Println("\nGot an error")
			log.Fatal(err)

		}
		fmt.Println()

		switch again {
		case 1:
			fmt.Printf("Addition of %d and %d is: %v", n1, n2, n1+n2)

		case 2:
			fmt.Printf("Difference between %d and %d is: %v", n1, n2, n1-n2)

		case 3:
			fmt.Printf("Product of %d and %d is: %v", n1, n2, n1*n2)

		case 4:
			fmt.Printf("Division of %d by %d is: %v", n1, n2, n1/n2)

		default:
			fmt.Printf("Invalid option exiting...")
			again = 0
		}

		if err != nil || err2 != nil || err1 != nil {
			fmt.Println("\nGot an error")
			log.Fatal(err)

		}

	}

}
