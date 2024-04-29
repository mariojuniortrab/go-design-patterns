package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)
}

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func NewSquare(size int) *Square {
	square := Square{}
	square.SetHeight(size)
	square.SetWidth(size)
	return &square
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)

	expectedArea := width * 10
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Println("Expected area:", expectedArea, ", but we got", actualArea)
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)
	s := NewSquare(15)
	UseIt(s)
}
