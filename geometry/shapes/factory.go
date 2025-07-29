package shapes

func FactoryShape(shapeType string) Shape {
	switch shapeType {
		case "rectangle":
			return NewRectangle()
		case "square":
			return NewSquare()
		default:
			return nil
	}
}
