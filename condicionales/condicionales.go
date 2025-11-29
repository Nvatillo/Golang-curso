package condicionales

import "fmt"

func CalcularMadurez(edad int) {

	switch {
	case edad < 13:
		fmt.Println("Niño")
	case edad >= 13 || edad <= 17:
		fmt.Println("Adolescente")
	case edad >= 18 || edad <= 64:
		fmt.Println("Adulto")
	case edad > 64:
		fmt.Println("Adulto mayor")
	}
}
