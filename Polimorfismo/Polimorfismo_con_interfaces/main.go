package main

import "fmt"

// factory method

type PayMethod interface {
	Pay()
}

// paypal
type Paypal struct {
}

func (p Paypal) Pay() {
	fmt.Println("pagado por Paypal")
}

// tarjeta
type CreditCart struct {
}

func (c CreditCart) Pay() {
	fmt.Println("pagado con tarjeta")
}

// cash
type Cash struct {
}

func (C Cash) Pay() {
	fmt.Println("pagado en cash")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return CreditCart{}
	case 3:
		return Cash{}
	default:
		return nil
	}
}

// polimorfismo es mandar una señal a diferentes objetos y que estos tengan la capaciad de responder
func main() {
	fmt.Println("digite un metodo de pago: ")

	var opciones string = "1 para payapl 2 para tarjeta 3 para cash"
	fmt.Println(opciones)
	var methodpay uint

	_, err := fmt.Scanln(&methodpay)
	if err != nil {
		panic("introduce un metodo valido")
	}

	if methodpay > 3 {
		fmt.Println("no es un metodo de pago correcto")
	}

	payMethod := Factory(methodpay)
	payMethod.Pay()
}
