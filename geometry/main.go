package main

import (
	"fmt"
	"geometry/shapes"
)

func printShapeInfo(s shapes.Shape) {
	fmt.Printf("My type is: %T\n", s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	var shapeType string
	var list []shapes.Shape;

	for {
		fmt.Println("\nWrite the new shape type to add")
		fmt.Scanln(&shapeType)

		if shapeType == "quit"{
			break
		}

		shape := shapes.FactoryShape(shapeType)
		
		if shape != nil {
			list = append(list, shape)
		} else{
			fmt.Println("Invalid shape type")
		}
	}

	var listLenght = len(list)
	fmt.Printf("\nList lenght: %d\n", listLenght)

	for _, shape  := range list {
		printShapeInfo(shape)
	}
}
