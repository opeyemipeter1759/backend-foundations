package main

import (
	"fmt"
	"math"
	"math/rand"
)
func add( x int, y int)  int {
	return x + y
}

func name( firstName , lastName string)  string {
	return  "I am " + firstName + " and my surname is " + lastName
}

func swap( x , y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func basic ()  {
	fmt.Println("Hello, Go!")
	x:=16.0
	fmt.Println("The square root of", x, "is", math.Sqrt(x))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("1 + 1 =", add(1, 1))
	fmt.Println(name("John", "Doe"))
	a, b := swap("first", "second")
	fmt.Println("Swapped:", a, b)
	fmt.Println(split(1000))

	var hello  string = "Hello, World!"
	var c, python, java  bool  = true, false, true
	k:= 42

	fmt.Println(hello, c, python, java, k)
}


func stringMani() {
	name := "Alice"
	age := 30
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
	pi := 3.14159
	fmt.Printf("The value of pi is approximately %.2f\n", pi) // Limiting to 2 decimal place
}
