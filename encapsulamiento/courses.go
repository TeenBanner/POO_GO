package main

import "fmt"

//setters y getters no existen como tal pero si podemos simularlos creando metodos

func (c *Course) SetClasses(classes map[uint]string) {
	c.Classes = classes
}

// identificadores exportados/ no exportados a nivel de paquete
type Course struct {
	Name    string
	Price   float64
	isFree  bool // si se cambia la primera letra por minuscula este ya no puede ser exportado
	UserIDs []uint
	Classes map[uint]string
}

// la diferencia entre estos 2 es que los exportados su primer caracter empieza con mayuscula
// metodo constructor(funcion contructora)
func NewCourse(name string, price float64, IsFree bool) *Course { // esta funcion constructora me devuelve un puntero a curso
	return &Course{
		Name:   name,
		Price:  price,
		isFree: IsFree,
	}
}

// si llegaramos a nesecitar unmetodo constructor podemos crear una funcion new la cual le va a asignar valores por defecto
// con esta nos retornaria una funcion inicializada con los valores que pongamos
// no es nesesario un metodo constructor ya que si instanciamoos un objeto nuevo y no definimos el valor de sus atributos estos por defecto tendren un valor cero
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c *Course) PrintClasses() {
	text := "las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func (c *Course) SetName(name string) {
	c.Name = name
}

//. en el caso de los getter nos es nesesario poner get de prefijo sino que solo un nombre que hace referencia al campo
func (c *Course) NAME() string {
	return c.Name
}

func main() {
	Go := Course{
		Name:    "Go desde cero",
		Price:   12.34,
		isFree:  false,
		UserIDs: []uint{12, 56, 99},
		Classes: map[uint]string{
			1: "instruccion",
			2: "estructuras",
			3: "Maps",
		},
	}
	fmt.Println(Go.Price)
	// instanciamos un objeto con la funcion constructora
	HtmlCourse := NewCourse("html desde cero", 14.99, false)

	fmt.Println(HtmlCourse.Name)
	// vamos a usar setter cuando no queramos que los datos se puedan modificar/acceder desde afuera del paquete pero si es nesesario proovememos un metodo para modificar el valor
	Go.SetClasses(map[uint]string{
		1: "variables y tipos de datos",
		2: "ciclos",
		3: "condicionales",
		4: "final",
	})
}
