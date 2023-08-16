package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}
var dilbert Employee

func EmployeeByID(id int) *Employee { /* ... */ }

func main() {
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason
}
