package shapes

import (
    "fmt"
    "geometry/utils"
)

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func NewRectangle() *Rectangle {
    fmt.Print("Enter the length of the rectangle: ")
	var height = utils.ScanNumber()
    fmt.Print("Enter the width of the rectangle: ")
	var width = utils.ScanNumber()
	return &Rectangle{Width: width, Height: height}
}