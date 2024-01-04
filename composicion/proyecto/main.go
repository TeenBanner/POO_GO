package main

import "fmt"

// customer pk
type Customer struct {
	name   string
	addres string
	phone  string
}

// new nos permite
func NewCustomer(name, address, phone string) Customer {
	return Customer{name, address, phone}
}

//invoice item
// contains info about invoice item
type Item struct {
	id      uint
	product string
	value   float64
}

//value getter
func (i Item) Value() float64 {
	return i.value
}

// new returns a new item
func NewItem(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// invoice is the struct of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  Customer
	items   []Item // decimos que tenemos una relacion de uno a muchos (obtenemos un slice de tipo item)
}

// set total is a setter of invoice.Total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}

// NewInvoic returns a new invoice
func NewInvoice(country, city string, client Customer, items []Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

func main() {
	// make new Invoice
	i := NewInvoice("Colombia", "popaya",
		NewCustomer("alejandro", "cl 123", "+40 233 572 3024"),
		[]Item{NewItem(1, "Curso de Go", 12.34),
			NewItem(2, "Curso POO GO", 54.23),
			NewItem(3, "Curso de Testing con Go", 90.00)})

	// Calculate Total
	i.SetTotal()

	// Print the invoice
	fmt.Printf("%+v", i)
}
