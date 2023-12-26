package main

import "fmt"

func main() {
	go loop(0, 1000000)
	for i := 0; i < 1000000; i++ {
		fmt.Println("proceso main iteracion numero: ", i)
	}
}

func loop(n, n2 int) {
	valor_usuario := n2

	for i := n; i < valor_usuario; i++ {
		fmt.Println("proceso 2 iteracion: ", i)
	}
}
