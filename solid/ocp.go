package main

import (
	"fmt"
)

/*
		The Open/Closed Principle (OCP)

		Software entities (classes, modules, functions, etc.) should be open for extension,
	but closed for modification. Such an entity can allow its behaviour to be extended without
	modifying its source code.

		- A module will be said to be open if it is still available for extension.
		For example, it should be possible to add fields to the data structures it contains,
		or new elements to the set of functions it performs.

		- A module will be said to be closed if it is available for use by other modules.
		This assumes that the module has been given a well-defined, stable description (the interface
		in the sense of information hiding).

	Ref: https://en.wikipedia.org/wiki/Open/closed_principle
 	Ref: http://joelabrahamsson.com/a-simple-example-of-the-openclosed-principle/
*/

type Size int
type Color int

const (
	Small Size = iota
	Medium
	Large
	Giant
)

const (
	Red Color = iota
	Green
	Blue
	Yellow
	Black
	White
)

type Product struct {
	name  string
	size  Size
	color Color
}

func (p *Product) GetName() string {
	return p.name
}

func NewProduct(name string, size Size, color Color) *Product {
	return &Product{
		name:  name,
		size:  size,
		color: color,
	}
}

// !!! The wrong way
type WrongFilter struct{}

func (pf *WrongFilter) BySize(products []*Product, size Size) []*Product {
	result := []*Product{}

	for _, product := range products {
		if product.size == size {
			result = append(result, product)
		}
	}

	return result
}

func (pf *WrongFilter) ByColor(products []*Product, color Color) []*Product {
	result := []*Product{}

	for _, product := range products {
		if product.color == color {
			result = append(result, product)
		}
	}

	return result
}

func (pf *WrongFilter) BySizeAndColor(products []*Product, size Size, color Color) []*Product {
	result := []*Product{}

	for _, product := range products {
		if product.size == size && product.color == color {
			result = append(result, product)
		}
	}

	return result
}

// !!! The better way
type ISpecification interface {
	IsSatisfied(*Product) bool
}

type ColorSpecification struct {
	color Color
}

func (cs *ColorSpecification) IsSatisfied(product *Product) bool {
	return cs.color == product.color
}

func NewColorSpecification(color Color) ISpecification {
	return &ColorSpecification{color: color}
}

type SizeSpecification struct {
	size Size
}

func (cs *SizeSpecification) IsSatisfied(product *Product) bool {
	return cs.size == product.size
}

func NewSizeSpecification(size Size) ISpecification {
	return &SizeSpecification{size: size}
}

type AndSpecification struct {
	first  ISpecification
	second ISpecification
}

func (as *AndSpecification) IsSatisfied(product *Product) bool {
	return as.first.IsSatisfied(product) && as.second.IsSatisfied(product)
}

func NewAndSpecification(first, second ISpecification) ISpecification {
	return &AndSpecification{first: first, second: second}
}

type IFilter interface {
	Filter([]*Product, ISpecification) []*Product
}

type RightFilter struct{}

func (rf *RightFilter) Filter(products []*Product, specification ISpecification) []*Product {
	result := []*Product{}

	for _, product := range products {
		if specification.IsSatisfied(product) {
			result = append(result, product)
		}
	}

	return result
}

func main() {
	bike := NewProduct("Bike", Small, Blue)
	motorcycle := NewProduct("Motorcycle", Small, Green)
	car := NewProduct("Car", Medium, Green)
	truck := NewProduct("Truck", Large, Red)
	train := NewProduct("Train", Large, Yellow)

	products := []*Product{bike, motorcycle, car, truck, train}

	fmt.Println("--The wrong way to filter things...")
	wrongFilter := &WrongFilter{}

	fmt.Println()
	fmt.Println("--Small things:")
	smallThings := wrongFilter.BySize(products, Small)
	for _, thing := range smallThings {
		fmt.Println(thing.GetName())
	}

	fmt.Println()
	fmt.Println("--Green things:")
	greenThings := wrongFilter.ByColor(products, Green)
	for _, thing := range greenThings {
		fmt.Println(thing.GetName())
	}

	fmt.Println()
	fmt.Println("--Large and yellow things:")
	largeYellowThings := wrongFilter.BySizeAndColor(products, Large, Yellow)
	for _, thing := range largeYellowThings {
		fmt.Println(thing.GetName())
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("--The right way to filter things...")
	rightFilter := &RightFilter{}

	fmt.Println()
	fmt.Println("--Small things:")
	smallSpecification := NewSizeSpecification(Small)
	smallThings = rightFilter.Filter(products, smallSpecification)
	for _, thing := range smallThings {
		fmt.Println(thing.GetName())
	}

	fmt.Println()
	fmt.Println("--Green things:")
	greenSpecification := NewColorSpecification(Green)
	greenThings = rightFilter.Filter(products, greenSpecification)
	for _, thing := range greenThings {
		fmt.Println(thing.GetName())
	}

	fmt.Println()
	fmt.Println("--Large and yellow things:")
	yellowSpecification := NewColorSpecification(Yellow)
	largeSpecification := NewSizeSpecification(Large)
	largeYellowThingsSpecification := NewAndSpecification(largeSpecification, yellowSpecification)
	largeYellowThings = rightFilter.Filter(products, largeYellowThingsSpecification)
	for _, thing := range largeYellowThings {
		fmt.Println(thing.GetName())
	}
}
