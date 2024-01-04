package main

import "fmt"

// estructura persona
type Person struct {
	Name string
	Age  uint
}

// metodo constructor
func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

// metodo para saludar
func (n Person) greet() {
	println("Hello")
}

// embebemos la estructura Person en empleado
type Employee struct {
	// al realizar los campos embebidos estamos asignando los campos y metodos del tipo interno al externo
	Person // de Person asignamos a Empleado
	Salary float64
}

// contructor de empleado
func NewEmployee(name string, age uint, salary float64) Employee {
	return Employee{NewPerson(name, age), salary}
}

// calcular nomina

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {
	e := NewEmployee("alfredo", 34, 3000)
	e.Payroll()
	// atraves de empleado accedemos a persona
	fmt.Println(e.Name)
	e.greet()
	fmt.Println(e.Age)
}
