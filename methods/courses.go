package main

import "fmt"

//  usuamos structuras en lugar de clases
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	classes map[uint]string
}

// metodos con receptores de puntero
// los metodos se comportan como las funciones asi que cuando le pasamos un valor no es el valor original sino que es una copia
func (c *Course) ChangePrice(price float64) {
	c.Price = price // cuando modificamos valores es casi obliagtorio usar punteros
}

// metodos p1
// los metodos son funciones que se declaran fuera de la estructura
// asi vinculamos un metodo a una estructura es mala practica usar self o this para referenciar
// cuando un metodo usa un puntero a la estructura como receptor es preferible que todos los demas metodos modifique o no valores se usen tambien punteros
func (c *Course) PrintClasses() { // metodo que tiene un receptor de tipo course es impoertante referenciar correctamente a la estructura correspondiente
	text := "las clases son: "
	for _, class := range c.classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// podemos crear metodos en otros archivos siempre y cuando estos pertenezcan al mismo paquete
