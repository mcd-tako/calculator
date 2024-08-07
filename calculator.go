// Package calculator does simple calculations
// Author: Tiago Macedo Cesar
// Date: Aug 3 2024, 08:40h
package calculator

import "errors"

// Add takes two numbers and return the result
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and return the result
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and return the result
func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}
