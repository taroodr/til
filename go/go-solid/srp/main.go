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

type EmployeeData struct {
	WorkingType Working
}

type EmployeeFacade struct {
	saver EmployeeSaver
	reporter HourReporter
	calculator PayCalculator
}

func NewEmployeeFacade() *EmployeeFacade {
	return &EmployeeFacade{
		saver: EmployeeSaver{},
		reporter: HourReporter{},
		calculator: PayCalculator{},
	}
}

func (e *EmployeeFacade) Save(data EmployeeData) error {
	return e.saver.saveEmployee(data)
}

func (e *EmployeeFacade) CalculatePay(data EmployeeData) int {
	return e.calculator.calculatePay(data)
}

func (e *EmployeeFacade) ReportHours(data EmployeeData) time.Duration {
	return e.reporter.reportHours(data)
}

type EmployeeSaver struct {}

func (e *EmployeeSaver) saveEmployee(data EmployeeData) error {
	return nil
}

type PayCalculator struct {}

func (e *PayCalculator) calculatePay(data EmployeeData) int {
	return 0
}

type HourReporter struct {}

func (e *HourReporter) reportHours(data EmployeeData) time.Duration {
	return time.Duration(0)
}

// Actors

type CTO struct {
	EmployeeFacade *EmployeeFacade
}

func (a *CTO) Exec(data EmployeeData)error {
	return a.EmployeeFacade.Save(data)
}

type COO struct {
	EmployeeFacade *EmployeeFacade
}

func (a *COO) Exec(data EmployeeData) error {
	_ = a.EmployeeFacade.ReportHours(data)
	return nil
}

type CFO struct {
	EmployeeFacade *EmployeeFacade
}

func (a *CFO) Exec(data EmployeeData) error {
	_ = a.EmployeeFacade.CalculatePay(data)
	return nil
}

func main() {
	EmployeeFacade := NewEmployeeFacade()
	cto := CTO{
		EmployeeFacade: EmployeeFacade,
	}
	cto.Exec(EmployeeData{WorkingType: Fulltime})

	coo := COO{
		EmployeeFacade: EmployeeFacade,
	}
	coo.Exec(EmployeeData{WorkingType: Parttime})

	cfo := CFO{
		EmployeeFacade: EmployeeFacade,
	}
	cfo.Exec(EmployeeData{WorkingType: Intern})
}