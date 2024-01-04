package main

import (
	"fmt"
	"strings"
)

// las inrfaces vacias nos sirven para controlar metodos o funciones deconocidad

func wrapper(i interface{}) {
	fmt.Printf("Valor %v, Tipo: %T\n", i, i)

	// type assertion nos permite validar el tipo concreto de la interfz vacia

	v, ok := i.(string)
	if ok {
		fmt.Println(strings.ToUpper(v))
	}

	// switch

	switch v := i.(type) {
	case string:
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	case uint:
		fmt.Println(v * 2)
	}
}

func main() {

	wrapper(12)

	wrapper("hola")
}
