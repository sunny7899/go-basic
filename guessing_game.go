package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	var playAgain string

	for {
		// Generate a random number between 0 and 9
		number := rand.Intn(10)

		// Prompt the user to guess
		fmt.Println("Guess a number between 0 and 9:")
		var guess int
		fmt.Scan(&guess)

		// Check if the guess is correct
		if guess == number {
			fmt.Println("You Won!")
		} else {
			fmt.Printf("You Lost! The correct number was %d\n", number)
		}

		// Ask if they want to play again
		fmt.Print("Do you want to play again? (y/n): ")
		fmt.Scan(&playAgain)

		if playAgain != "y" && playAgain != "Y" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}
