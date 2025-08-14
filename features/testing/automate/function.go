package automate

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary int
}

func (e Employee) Greet() string {
	return fmt.Sprintf("Hello %s", e.Name)
}
