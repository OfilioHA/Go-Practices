package shapes

func FactoryShape(shapeType string) Shape {
	switch shapeType {
		case "rectangle":
			return NewRectangle(5, 10)
		case "square":
			return NewSquare(5)
		default:
			return nil
	}
}
