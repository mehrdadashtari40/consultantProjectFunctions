package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandNum(length int) string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Create a slice to store the random numbers as strings
	var randomNumbers []string

	// Generate random digits and store them in the slice
	for i := 0; i < length; i++ {
		randomNumbers = append(randomNumbers, fmt.Sprintf("%d", rand.Intn(10)))
	}

	// Join the strings into a single string
	result := strings.Join(randomNumbers, "")

	return result
}

func main() {
	// Call the function to generate random numbers
	randomNumbers := GenerateRandNum(5)

	// Print the generated random numbers
	fmt.Println("Generated Random Numbers:", randomNumbers)
}
