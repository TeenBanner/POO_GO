package main

import "fmt"

func main() {
	// instanciar structuras (creamos instancias de la estructura cursos)
	Go := Course{ // debemos de pasar un puntero a curso a la hora de llamar a sus metodos
		Name:    "Go desde cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 99},
		classes: map[uint]string{
			1: "instruccion",
			2: "estructuras",
			3: "Maps",
		},
	}
	// acceder a los campos
	fmt.Println(Go.Name) // imprime el nombre del curso
	// podemos declarar los valores que contiene el campo sin declarar el campo de la estrucuta siempre que los declaremos en el mismo orden que en la estructura
	Go2 := Course{
		"go desde cero:",
		12.46,
		false,
		[]uint{15, 23, 42},
		map[uint]string{
			1: "introduccion",
			2: "arrays",
			3: "maps",
		},
	}
	fmt.Println(Go2.Price)
	// si no queremos utilizar todos los campos de la estructura debemos especificarlos al instanciar el objeto
	css := Course{Name: "css desde cero", IsFree: true}
	// al no usar todos los campos los que queden sin asignar tomaran su valor cero
	js := Course{}
	// podemos darle valor despues de instamciar el objetoa los atributos
	js.Name = "curso javascript"
	fmt.Println(css.IsFree)
	fmt.Printf("%+v", css)

	Go.PrintClasses()
	// apuntamos al objeto go para modificar el valor
	(&Go).ChangePrice(67.12) // esto es automatico pero tenemos que tener en cuenta este proceso
	fmt.Println(Go.Price)
}
