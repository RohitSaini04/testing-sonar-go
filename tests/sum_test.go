package test

import (
	"testing"
	"variant-inc_dx-demo-go-api/app/sum"
)

func TestSum(t *testing.T) {
	// Define test cases
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},      // Test case 1: 1 + 2 = 3
		{-1, 1, 0},     // Test case 2: -1 + 1 = 0
		{100, -50, 50}, // Test case 3: 100 + (-50) = 50
		{0, 0, 0},      // Test case 4: 0 + 0 = 0
	}

	// Iterate through test cases
	for _, test := range tests {
		// Call the function being tested
		result := sum.Sum(test.a, test.b)
		// Check if the result matches the expected value
		if result != test.expected {
			t.Errorf("sum(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}
