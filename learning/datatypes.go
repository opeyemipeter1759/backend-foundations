package main

import "fmt"

func varTest() {
	var realAge int
    age := 30
    name := "John Doe"
    age, newName := 40, "Jane Doe" // Valid: newName is a new variable
	newAge, newName := 50, "Peter"
	realAge = 90
    fmt.Println(age, newName, name, newAge,realAge)

    // This will cause a compilation error because both variables are already declared
}