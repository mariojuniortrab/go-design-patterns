package main

import "fmt"

//OCP
//Specification pattern

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

// ============ Old filter ============
type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// ====================================

//specification:

type FilterSpecication interface {
	ConditionIsSatisfied(*Product) bool
}

type ColorFilter struct {
	color Color
}

func (f *ColorFilter) ConditionIsSatisfied(p *Product) bool {
	return f.color == p.color
}

type SizeFilter struct {
	size Size
}

func (f *SizeFilter) ConditionIsSatisfied(p *Product) bool {
	return f.size == p.size
}

// Composition
type FilterAnd struct {
	filters []FilterSpecication
}

func (f *FilterAnd) Add(spec FilterSpecication) {
	f.filters = append(f.filters, spec)
}

func (f *FilterAnd) ConditionIsSatisfied(p *Product) bool {

	for i := range f.filters {
		if !f.filters[i].ConditionIsSatisfied(p) {
			return false
		}
	}

	return true
}

func BetterFilter(products []Product, filterSpecification FilterSpecication) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if filterSpecification.ConditionIsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Println("Green products (old)")

	filter := Filter{}
	for _, v := range filter.FilterByColor(products, green) {
		fmt.Printf(" - %s is green! \n", v.name)
	}

	fmt.Println("Green products (New)")

	greenSpec := ColorFilter{green}
	largeSpec := SizeFilter{large}

	and := FilterAnd{}
	and.Add(&greenSpec)
	and.Add(&largeSpec)
	for _, v := range BetterFilter(products, &and) {
		fmt.Printf(" - %s found! \n", v.name)
	}
}
