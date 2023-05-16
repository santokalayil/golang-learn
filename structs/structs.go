package structs

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

type Rectangle struct {
	width  float64
	height float64
}

// Method to calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method to calculate the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func RunStructExamples() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	// structs with methods
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Println("Area:", rectangle.Area())
}
