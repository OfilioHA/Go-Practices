package shapes

type Square struct {
    Side float64
}

func (s *Square) Area() float64 {
    return s.Side * s.Side
}

func (s *Square) Perimeter() float64 {
    return 4 * s.Side
}

func NewSquare(side float64) *Square {
	return &Square{Side: side}
}