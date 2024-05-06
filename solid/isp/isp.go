package main

import "fmt"

type Document struct{}

// type Machine interface {
// 	Print(d Document)
// 	Scan(d Document)
// 	Fax(d Document)
// }

// type MultifunctionalPrinter struct{}

// func (m MultifunctionalPrinter) Print(d Document) {
// 	fmt.Println(d)
// }

// func (m MultifunctionalPrinter) Scan(d Document) {
// 	fmt.Println(d)
// }

// func (m MultifunctionalPrinter) Fax(d Document) {
// 	fmt.Println(d)
// }

// type OldFashionedPrinter struct{}

// func (o OldFashionedPrinter) Print(d Document) {
// 	fmt.Println(d)
// }

// func (o OldFashionedPrinter) Scan(d Document) {
// 	panic("I can not do this!")
// }

// func (o OldFashionedPrinter) Fax(d Document) {
// 	panic("I can not do this!")
// }

//ISP

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Fax interface {
	Fax(d Document)
}

type MyPrinter struct{}

func (p MyPrinter) Print(d Document) {
	fmt.Println(d)
}

type MyScanner struct{}

func (s MyScanner) Scan(d Document) {
	fmt.Println(d)
}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {
	fmt.Println(d)
}
func (p Photocopier) Scan(d Document) {
	fmt.Println(d)
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	Fax
}

// Decorator
type MultiFunctionMachine struct {
	scanner Scanner
	printer Printer
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}
func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	d := &Document{}
	p := &MyPrinter{}
	s := &MyScanner{}
	m := MultiFunctionMachine{s, p}

	m.Scan(*d)
	m.Print(*d)
}
