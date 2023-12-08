package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers() []string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Create a slice to store the random numbers as strings
	randomNumbers := make([]string, 5)

	// Generate 5 random numbers and store them in the slice
	for i := 0; i < 5; i++ {
		randomNumbers[i] = fmt.Sprintf("%d", rand.Intn(10))
	}

	return randomNumbers
}

func main() {
	// Call the function to generate random numbers
	randomNumbers := generateRandomNumbers()

	// Print the generated random numbers
	fmt.Println("Generated Random Numbers:", randomNumbers)
}
