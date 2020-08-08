package main

import "time"

// Employee

type Employee struct{}

func (e *Employee) CalculatePay() int {
	return 0
}

func (e *Employee) ReportHours() time.Duration {
	return time.Duration(0)
}

func (e *Employee) Save(text string) error {
	return nil
}

// Actors

type CTO struct {
	Employee *Employee
}

func (a *CTO) Exec(text string)error {
	return a.Employee.Save(text)
}

type COO struct {
	Employee *Employee
}

func (a *COO) Exec() error {
	_ = a.Employee.ReportHours()
	return nil
}

type CFO struct {
	Employee *Employee
}

func (a *CFO) Exec() error {
	_ = a.Employee.CalculatePay()
	return nil
}

func main() {
	cto := CTO{
		Employee: &Employee{},
	}
	cto.Exec("any")

	coo := COO{
		Employee: &Employee{},
	}
	coo.Exec()

	cfo := CFO{
		Employee: &Employee{},
	}
	cfo.Exec()
}