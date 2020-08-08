package main

import "time"

// ENUM

type Working int

const (
	Parttime Working = iota // パートタイム
	Fulltime             // フルタイム
	Intern               // インターン
)

func (w Working) String() string {
	switch w {
	case Parttime:
		return "パートタイム"
	case Fulltime:
		return "フルタイム"
	case Intern:
		return "インターン"
	}
	return ""
}

// Employee

type Employee struct{
	WorkingType Working
}

func (e *Employee) CalculatePay() int {
	_ = e.regularHours()
	return 0
}

func (e *Employee) ReportHours() time.Duration {
	_ = e.regularHours()
	return time.Duration(0)
}

func (e *Employee) Save(text string) error {
	return nil
}

func (e *Employee) regularHours() time.Duration {
	switch e.WorkingType  {
	case Parttime, Intern:
		return 5 * time.Hour
	}
	return 8* time.Hour
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
		Employee: &Employee{
			WorkingType: Fulltime,
		},
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