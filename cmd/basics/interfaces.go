package main

import "fmt"

// Shape interface defines methods for calculating area and perimeter of a shape.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Payment Gateway Interface example

type PaymentProcesser interface {
	pay(amount float64)
}

type Finix struct {
}

func (f Finix) pay(amount float64) {
	fmt.Printf("Paying the Amount %f using Finix", amount)
}

type TabaPay struct {
}

func (t TabaPay) pay(amount float64) {
	fmt.Printf("Paying the Amount %f using Tabapay", amount)
}

func processPayment(p PaymentProcesser, amount float64) {
	p.pay(amount)
}

func processShapeInterface() {
	fmt.Print("\n")
	var s Shape
	s = Rectangle{Width: 5, Height: 10}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())

	s = Circle{Radius: 7}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())
}

func ProcessPaymentProcessorInterface() {
	var payment PaymentProcesser
	payment = Finix{}

	processPayment(payment, 234.33)
	payment = TabaPay{}
	processPayment(payment, 123.9)

}
func InterfacesExample() {
	ProcessPaymentProcessorInterface()
	processShapeInterface()
}
