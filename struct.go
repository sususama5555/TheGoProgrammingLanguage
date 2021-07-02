package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) //"Pointy-haired boss"
	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason
}

func EmployeeByID(id int) *Employee {
	return &Employee{
		1,
		"2",
		"2",
		time.Now(),
		"23",
		1,
		2,
	}
}
