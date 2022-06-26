package perusahaan

import "fmt"

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
	// return 0 // TODO: replace this
	var total int
	for _, subordinate := range c.Subordinate {
		total += subordinate.TotalDivisonSalary()
	}
	total += c.GetSalary()
	fmt.Println("Total divison CTO salary:", total)
	return total
}
