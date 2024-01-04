// Implemetar receptores de puntero
package main

import "fmt"

type Storage interface {
	Get() string
	Set(string)
}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name: "adrian"}
}

// por buenas practicas si hay un metodo que usa un puntero todos deben usarlo tambien para evitar problemas con el scope
func (p *Person) Get() string {
	return p.name
}

// la struct persona tiene disponible el set por ser de tipo puntero y get porque el valor de referencia es de tipo persona
func (p *Person) Set(name string) {
	p.name = name
}

func Exec(s Storage, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}
func main() {
	// p hace referencia a Persona
	p := NewPerson("Alejandro")
	p.Set("alexys")
	fmt.Println(p.Get())
	// tenemos que pasar un puntero para que haga referencia a la estructura persona y que p tengo el metodo set
	Exec(p, "Alvaro")
	// si no pasamos un puntero no podremos acceder a set porque no estan en el mismo scope
}
