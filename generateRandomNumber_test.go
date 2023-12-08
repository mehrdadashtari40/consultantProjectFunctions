package main

import (
	"testing"
)

func TestGenerateRandomNumbers(t *testing.T) {
	// Run the function to generate random numbers
	randomNumbers := generateRandomNumbers()

	// Check if the length is 5
	if len(randomNumbers) != 5 {
		t.Errorf("Expected length 5, but got %d", len(randomNumbers))
	}

	// Check if all characters are numbers between 0 and 9
	for _, char := range randomNumbers {
		if char < '0' || char > '9' {
			t.Errorf("Character %s is not a number between 0 and 9", string(char))
		}
	}
}
