package main

import "fmt"

// los tipos hace referencia en los tipos declarados de go que es hacer nuevos tipos a partir de tipos ya existentes con la condicion que sus metodos tiene que ser declarados en el paquete en el que son usados
// es decir que si queremos agregarles metodos tendremos que declararlos primero y esta es la razon por la cual no podemos agregarle metodos a tipos como strign, bool, etc..
type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// declaraciones de alias
type myAlias = course

// definicion de tipo
type newCourse course

// definiciones de tipo con base a  tipos preedeclarados

type newBool bool

// hacer esto nos permitiria agregarle un metodo aun tipo de dato
func (c newBool) String() string {
	if c {
		return "VERDADERO"
	}

	return "FALSE"
}

func main() {
	// newCourse tiene acceso a los campos de la estructura course pero no a sus metodos
	c := newCourse{name: "go"}
	// en la declaracion de un tipo se heredan los campos pero no los metodos
	// c.Print no funciona porque es de tipo newCourse
	// imprime el tipo de dato
	fmt.Printf("El tipo es: %T\n", c)

	var b newBool = true

	fmt.Println(b.String())
}
