package calculadora

import (
	"errors"
)

func Calcular(num1 float64, num2 float64, operacion string) (float64, error) {
	switch operacion {
	case "sumar":
		return sumar(num1, num2)
	case "restar":
		return restar(num1, num2)
	case "dividir":
		return dividir(num1, num2)
	case "multiplicar":
		return multiplicar(num1, num2)
	default:
		return 0, errors.New("operacion no existente")
	}
}

func sumar(num1 float64, num2 float64) (float64, error) {
	return num1 + num2, nil
}

func restar(num1 float64, num2 float64) (float64, error) {
	return num1 - num2, nil
}

func dividir(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return num1 / num2, nil
}

func multiplicar(num1 float64, num2 float64) (float64, error) {
	return num1 * num2, nil
}
