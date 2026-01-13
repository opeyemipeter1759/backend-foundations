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


func main ()  {
	fmt.Println("Hello, Go!")
	x:=16.0
	fmt.Println("The square root of", x, "is", math.Sqrt(x))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("1 + 1 =", add(1, 1))
	fmt.Println(name("John", "Doe"))
	a, b := swap("first", "second")
	fmt.Println("Swapped:", a, b)
}
