package main

import "fmt"

func main() {
	basic()
	arithmeticCalc()
	stringMani()
}

func arithmeticCalc() {
    a := 10
    b := 5
    sum := a + b
    difference := a - b
    product := a * b
    quotient := a / b

    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
    fmt.Println("Product:", product)
    fmt.Println("Quotient:", quotient)
}