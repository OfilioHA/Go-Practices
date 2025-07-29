package shapes

import (
    "fmt"
    "geometry/utils"
)

type Square struct {
    Side float64
}

func (s *Square) Area() float64 {
    return s.Side * s.Side
}

func (s *Square) Perimeter() float64 {
    return 4 * s.Side
}

func NewSquare() *Square {
    fmt.Print("Ingrese el lado del cuadrado: ")
	var side = utils.ScanNumber()
    return &Square{Side: side}
}