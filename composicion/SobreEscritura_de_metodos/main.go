package main

// en go no existe la sobreescritura de metodos si no que cuando nosotros "sobre escribimos el metodos" en realidad estamos oculatando para la clase hija el atributo original
import "fmt"

type Human struct {
	Age      uint
	Children uint
}

func NewHuman(age, children uint) Human {
	return Human{age, children}
}

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
func (n Person) Greet() {
	println("Hello")
}

// embebemos la estructura Person en empleado
type Employee struct {
	// al realizar los campos embebidos estamos asignando los campos y metodos del tipo interno al externo
	Person // de Person asignamos a Empleado
	Human  // al embeber mas de una estructura al querer usar el metodo age se podira generar un conflicto si ya existe en las otras 2 estructuras
	Salary float64
}

// nuevo metodo saludar
//
func (e Employee) Greet() {
	fmt.Println("Hello from Employee")
}

// contructor de empleado
func NewEmployee(name string, age, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}
}

// calcular nomina

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {
	e := NewEmployee("alfredo", 34, 2, 3000)
	e.Payroll()
	// atraves de empleado accedemos a persona
	fmt.Println(e.Name)
	e.Greet()
	fmt.Println(e.Human.Age) // el codigo funcionara aunque un campo sea ambiguo pero si nesecitamos usarlo debemos especificar de donde provendra
	// aunque este oculto el metodo original podemos acceder a el
	e.Person.Greet()
}
