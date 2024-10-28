package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var choice int

	// Asking for the first number
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Asking for the second number
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Showcasing the operation choices
	fmt.Println("Select an operation:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Print("Enter choice (1-4): ")
	_, err = fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > 4 {
		fmt.Println("Invalid choice. Please select an option between 1 and 4.")
		return
	}

	// Performing the chosen operation
	switch choice {
	case 1:
		fmt.Printf("Result: %.2f\n", num1+num2)
	case 2:
		fmt.Printf("Result: %.2f\n", num1-num2)
	case 3:
		fmt.Printf("Result: %.2f\n", num1*num2)
	case 4:
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Printf("Result: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Invalid operation selected.")
	}
}
