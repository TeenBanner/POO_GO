package main

import "fmt"

// una interfaz es un conjunto de firma de metodos

// por buenas practicas se fomenta el uso de interfaces de un solo metodo
// y que su nomre empieze por el nombre del metodo con el sufigo er

type Greeter interface {
	Greet()
}

// las interfaces proporcionan una manera de espicificar el comportamiento esperado de un objeto sin detallar sono se logra ese comportamiento
// cualquier metodo Greet estara satisfaciendo la interfaz Greeter

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s", p.Name)
}

type Text string

func (t Text) Greet() {
	fmt.Printf("hola soy un %v", t)
}

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor: %v, Tipo %T\n", g, g)
	}
}

func main() {
	var g Greeter = Person{Name: "alejandoro"}
	g.Greet()                      // llamada al metodo greet definido en la interfaz
	var g2 Greeter = Text("Daisy") // hacemos referencia al tipo definido
	g2.Greet()
	// llamar a los metodos por la funcion GreetAll()

	p := Person{Name: "Felipe"}
	var t Text = "daisy"
	GreetAll(p, t)
}
