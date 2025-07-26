package main

import (
	"fmt"
	"geometry/shapes"
)

func printShapeInfo(s shapes.Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	var shapeType string

	for (shapeType != "quit"){
		fmt.Scanln(&shapeType)

		shape := shapes.FactoryShape(shapeType)		
		if shape != nil {
			printShapeInfo(shape)
		} else if shapeType != "quit" {
			fmt.Println("Invalid shape type")
		}
	}
}
