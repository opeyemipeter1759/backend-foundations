package main

import "fmt"

func varTest() {
	var realAge int
	var studentName string // Declare a string variable for student's name
    var studentAge int    // Declare an integer variable for student's age
    var studentGPA float64 // Declare a float64 variable for student's GPA
    var isEnrolled bool   // Declare a boolean variable to check enrollment status

    studentName = "Alice Smith"
    studentAge = 20
    studentGPA = 3.85
    isEnrolled = true

    age := 30
    name := "John Doe"
    age, newName := 40, "Jane Doe" // Valid: newName is a new variable
	newAge, newName := 50, "Peter"
	realAge = 90
    fmt.Println(age, newName, name, newAge,realAge)
    fmt.Println("Student Name:", studentName)
    fmt.Println("Student Age:", studentAge)
    fmt.Println("Student GPA:", studentGPA)
    fmt.Println("Is Enrolled:", isEnrolled)
	declaredOps()
}


  func declaredOps() {
    courseName := "Go Basics" // Declare and initialize a string variable
    courseDuration := 4        // Declare and initialize an integer variable (weeks)
    isCompleted := false       // Declare and initialize a boolean variable
    fmt.Println("Course Name:", courseName)
    fmt.Println("Course Duration:", courseDuration, "weeks")
    fmt.Println("Is Completed:", isCompleted)
}