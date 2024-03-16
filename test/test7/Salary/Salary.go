package Salary

import "fmt"
import "java2go/test/utils"

type Employee struct {
	name    string
	address string
	number  int
}

func (rcvr *Employee) ComputePay() (float64) {
	fmt.Println("Inside Employee computePay")
	return 0.0
}
func (rcvr *Employee) GetAddress() (string) {
	return rcvr.address
}
func (rcvr *Employee) GetName() (string) {
	return rcvr.name
}
func (rcvr *Employee) GetNumber() (int) {
	return rcvr.number
}
func (rcvr *Employee) MailCheck() {
	fmt.Println(fmt.Sprintf("%v%v%v%v", "Mailing a check to ", rcvr.name, " ", rcvr.address))
}
func NewEmployeeStringStringint(name string, address string, number int) (rcvr *Employee) {
	rcvr = &Employee{}
	fmt.Println("Constructing an Employee")
	rcvr.name = name
	rcvr.address = address
	rcvr.number = number
	return
}
func (rcvr *Employee) SetAddressString(newAddress string) {
	rcvr.GetName()
	rcvr.name = "22"
	rcvr.GetName()
	rcvr.address = newAddress
}
func (rcvr *Employee) ToString() (string) {
	return fmt.Sprintf("%v%v%v%v%v", rcvr.name, " ", rcvr.address, " ", rcvr.number)
}

type Salary struct {
	*Employee
	salary float64
}

func (rcvr *Salary) ComputePay() (float64) {
	fmt.Println(fmt.Sprintf("%v%v", "Computing salary pay for ", rcvr.GetName()))
	return rcvr.salary / 52
}
func (rcvr *Salary) GetSalary() (float64) {
	return rcvr.salary
}
func (rcvr *Salary) MailCheck() {
	fmt.Println("Within mailCheck of Salary class ")
	fmt.Println(fmt.Sprintf("%v%v%v%v", "Mailing check to ", rcvr.GetName(), " with salary ", rcvr.salary))
}
func NewSalaryStringStringintdouble(name string, address string, number int, salary float64) (rcvr *Salary) {
	rcvr = &Salary{}
	rcvr.Employee = NewEmployeeStringStringint(name, address, number)
	rcvr.SetSalarydouble(salary)
	return
}
func (rcvr *Salary) SetSalarydouble(newSalary float64) {
	if newSalary >= 0.0 {
		rcvr.salary = newSalary
	}
}
