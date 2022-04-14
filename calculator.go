// Package calculator provides a library for
// simple calculations in Go.
package calculator

import "errors"

// Add takes two numbers and returns the
// result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and returns the result
// of multiplying b by a.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers a and b, and returns the result
// of dividing a by b.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
