package main

import (
	"fmt"
	"math"
	"strings"
)

// Project 1: Temperature Conversion
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// Project 2: Reverse a String
func reverseString(s string) string {
	// Convert string to slice of runes to handle Unicode properly
	runes := []rune(s)
	
	// Reverse the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	return string(runes)
}

// Project 3: Calculate Area of a Circle
func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	// Separator for output
	separator := strings.Repeat("=", 50)
	
	// Project 1: Temperature Conversion
	fmt.Println(separator)
	fmt.Println("Project 1: Temperature Conversion")
	fmt.Println(separator)
	celsius := 25.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)
	
	celsius = 0.0
	fahrenheit = celsiusToFahrenheit(celsius)
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)
	
	celsius = 100.0
	fahrenheit = celsiusToFahrenheit(celsius)
	fmt.Printf("%.1f°C = %.1f°F\n\n", celsius, fahrenheit)
	
	// Project 2: Reverse a String
	fmt.Println(separator)
	fmt.Println("Project 2: Reverse a String")
	fmt.Println(separator)
	original := "Hello, World!"
	reversed := reverseString(original)
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Reversed: %s\n", reversed)
	
	original = "Go Programming"
	reversed = reverseString(original)
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Reversed: %s\n\n", reversed)
	
	// Project 3: Calculate Area of a Circle
	fmt.Println(separator)
	fmt.Println("Project 3: Calculate Area of a Circle")
	fmt.Println(separator)
	radius := 5.0
	area := circleArea(radius)
	fmt.Printf("Radius: %.1f\n", radius)
	fmt.Printf("Area: %.2f square units\n", area)
	
	radius = 10.0
	area = circleArea(radius)
	fmt.Printf("Radius: %.1f\n", radius)
	fmt.Printf("Area: %.2f square units\n", area)
	
	radius = 7.5
	area = circleArea(radius)
	fmt.Printf("Radius: %.1f\n", radius)
	fmt.Printf("Area: %.2f square units\n", area)
}