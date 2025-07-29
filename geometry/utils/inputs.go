package utils

import "fmt"

func ScanNumber() float64 {
	for {
		var number float64
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Entrada no válida. Por favor, ingresa un número.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		return number
	}
}
