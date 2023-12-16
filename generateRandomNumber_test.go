package main

import (
	"testing"
	"unicode"
)

// TestGenerateRandNum tests if GenerateRandNum function returns a string of digits with the correct length
func TestGenerateRandNum(t *testing.T) {
	length := 10 // Example length for testing
	randStr := GenerateRandNum(length)

	// Check the length of the returned string
	if len(randStr) != length {
		t.Errorf("Expected string of length %d, got %d", length, len(randStr))
	}

	// Check that all characters in the string are digits
	for _, char := range randStr {
		if !unicode.IsDigit(char) {
			t.Errorf("Expected only digits in string, found non-digit: %c", char)
		}
	}
}