package main

import "fmt"

/*
		The Liskov Substitution Principle (LSP)

		Substitutability is a principle in object-oriented programming stating that,
	in a computer program, if S is a subtype of T, then objects of type T may be replaced with objects of type S
	(i.e. an object of type T may be substituted with any object of a subtype S) without altering any of
	the desirable properties of T (correctness, task performed, etc.).

		More formally, the Liskov substitution principle (LSP) is a particular definition of a subtyping relation,
	called (strong) behavioral subtyping, that was initially introduced by Barbara Liskov in a 1987 conference keynote
	address titled Data abstraction and hierarchy. It is a semantic rather than merely syntactic relation because
	it intends to guarantee semantic interoperability of types in a hierarchy, object types in particular.


	Ref: https://en.wikipedia.org/wiki/Liskov_substitution_principle
 	Ref: https://www.infragistics.com/community/blogs/dhananjay_kumar/archive/2015/06/30/simplifying-the-liskov-substitution-principle-of-solid-in-c.aspx
*/

type IRectangle interface {
	GetWidth() int
	GetHeight() int
	SetWidth(int)
	SetHeight(int)
	Area() int
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) Area() int {
	return r.width * r.height
}

type Square struct {
	Rectangle
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func NewRectangle(width, height int) IRectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func NewSquare(size int) IRectangle {
	result := &Square{}
	result.SetHeight(size)
	return result
}

func main() {
	rectangle := NewRectangle(10, 10)
	width := rectangle.GetWidth()
	rectangle.SetHeight(20)
	fmt.Printf("\nExpected %d, got %d\n", width*20, rectangle.Area())

	// LSP violated. Square should not inherit from rectangle.
	square := NewSquare(10)
	width = square.GetWidth()
	square.SetHeight(20)
	fmt.Printf("\nExpected %d, got %d\n", width*20, square.Area())
}
