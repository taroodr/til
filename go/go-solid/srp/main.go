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

type EmployeeSaver struct {}

func (e *EmployeeSaver) SaveEmployee(data EmployeeData) error {
	return nil
}

type PayCalculator struct {}

func (e *PayCalculator) CalculatePay(data EmployeeData) int {
	return 0
}

type HourReporter struct {}

func (e *HourReporter) ReportHours(data EmployeeData) time.Duration {
	return time.Duration(0)
}

// Actors

type CTO struct {
	EmployeeSaver *EmployeeSaver
}

func (a *CTO) Exec(data EmployeeData)error {
	return a.EmployeeSaver.SaveEmployee(data)
}

type COO struct {
	HourReporter *HourReporter
}

func (a *COO) Exec(data EmployeeData) error {
	_ = a.HourReporter.ReportHours(data)
	return nil
}

type CFO struct {
	PayCalculator *PayCalculator
}

func (a *CFO) Exec(data EmployeeData) error {
	_ = a.PayCalculator.CalculatePay(data)
	return nil
}

func main() {
	cto := CTO{
		EmployeeSaver: &EmployeeSaver{},
	}
	cto.Exec(EmployeeData{WorkingType: Fulltime})

	coo := COO{
		HourReporter: &HourReporter{},
	}
	coo.Exec(EmployeeData{WorkingType: Parttime})

	cfo := CFO{
		PayCalculator: &PayCalculator{},
	}
	cfo.Exec(EmployeeData{WorkingType: Intern})
}