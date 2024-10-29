package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"
)

// Character sets
const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}<>?"
)

// Function to generate a secure random password
func generatePassword(length int, includeUpper, includeNumbers, includeSpecial bool) (string, error) {
	// Validate length
	if length <= 0 {
		return "", fmt.Errorf("password length must be greater than zero")
	}

	// Build charset and guarantee inclusion of at least one of each selected type
	var charset strings.Builder
	var guaranteedChars []byte

	charset.WriteString(lowerChars) // Lowercase is always included
	guaranteedChars = append(guaranteedChars, randomChar(lowerChars))

	if includeUpper {
		charset.WriteString(upperChars)
		guaranteedChars = append(guaranteedChars, randomChar(upperChars))
	}
	if includeNumbers {
		charset.WriteString(numberChars)
		guaranteedChars = append(guaranteedChars, randomChar(numberChars))
	}
	if includeSpecial {
		charset.WriteString(specialChars)
		guaranteedChars = append(guaranteedChars, randomChar(specialChars))
	}

	// Fill the rest of the password with random characters from charset
	password := make([]byte, length)
	for i := range password {
		password[i] = randomChar(charset.String())
	}

	// Replace random positions with guaranteed character types to meet criteria
	for i, char := range guaranteedChars {
		randomIndex := randomIndex(length)
		password[randomIndex] = char
	}

	// Shuffle the final password for added randomness
	shuffle(password)

	return string(password), nil
}

// Helper function to pick a random character from a string
func randomChar(charset string) byte {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		log.Fatal("Failed to generate random character:", err)
	}
	return charset[num.Int64()]
}

// Helper function to generate a random index within a given range
func randomIndex(max int) int {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Fatal("Failed to generate random index:", err)
	}
	return int(num.Int64())
}

// Shuffle function to randomize the final password
func shuffle(password []byte) {
	for i := range password {
		j := randomIndex(len(password))
		password[i], password[j] = password[j], password[i]
	}
}

func main() {
	var length int
	var includeUpper, includeNumbers, includeSpecial bool

	// Get user input for password configuration
	fmt.Print("Enter password length: ")
	fmt.Scan(&length)
	fmt.Print("Include uppercase letters? (true/false): ")
	fmt.Scan(&includeUpper)
	fmt.Print("Include numbers? (true/false): ")
	fmt.Scan(&includeNumbers)
	fmt.Print("Include special characters? (true/false): ")
	fmt.Scan(&includeSpecial)

	// Generate the password
	password, err := generatePassword(length, includeUpper, includeNumbers, includeSpecial)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated Password:", password)
}



// SAMPLE INPUT OUTPUT AND EXPLANATION

// Sample Input:
// Password length: 12
// Include uppercase letters: true
// Include numbers: true
// Include special characters: true


// Sample Output:
// Generated Password: G9@bC2!qP$e1


// Explanation:
// Input: The desired length and character options for the password.
// Output: A random, secure password containing uppercase letters, numbers, and special characters, based on the specified criteria.
