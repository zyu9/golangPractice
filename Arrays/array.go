package main

import (
	"fmt"
)

func main() {
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of Students: %v\n", len(students))
}
