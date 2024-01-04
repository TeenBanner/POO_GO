package main

import (
	"fmt"
)

// la composicion es una alternativa al a herencia esta se basa en una relacion de uno a muchos
// Definición de una estructura addNames que contiene dos campos para nombres y apellidos.
type addNames struct {
	Name       string
	SecondName string
}

// Definición de una estructura User que utiliza composición anidada de addNames.
// Un usuario tiene dos instancias de addNames (Name y SecondName) y un campo para la contraseña.
type User struct {
	Name       addNames
	SecondName addNames
	Password   int
}

// Método AddName de la estructura User que permite agregar un nombre al campo Name.
func (c *User) AddName(name string) {
	c.Name.Name = name
}

// Método AddSecondName de la estructura User que permite agregar un segundo nombre al campo SecondName
func (c *User) AddSecondName(SecondName string) {
	c.SecondName.SecondName = SecondName
}

// Método CreatePassword de la estructura User que permite establecer la contraseña.
func (c *User) CreatePassword(password int) {
	c.Password = password
}

func main() {
	// Crear una instancia de User llamada "poncho".
	poncho := User{}
	// Agregar nombres a la instancia "poncho".
	poncho.AddSecondName("poncho")

	poncho.AddSecondName("Viajero")
	// Establecer la contraseña para la instancia "poncho".
	poncho.CreatePassword(12345)
	// Solicitar al usuario que ingrese la contraseña para continuar.
	fmt.Println("introduce your password to continue")

	var passwordtemp int
	// Leer la contraseña ingresada por el usuario.
	fmt.Scan(&passwordtemp)
	// Verificar si la contraseña ingresada coincide con la contraseña almacenada en "poncho".
	switch passwordtemp {
	case poncho.Password:
		fmt.Println("Welcome, Poncho!")
	default:
		fmt.Println("Wrong Password")
	}
}
