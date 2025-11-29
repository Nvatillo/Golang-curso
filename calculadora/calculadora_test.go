package calculadora

import (
	"testing"
)

func TestSumarExito(t *testing.T) {
	resultado, _ := Calcular(1, 1, "sumar")

	if resultado != 2 {
		t.Errorf("esperado 2, pero se obtuvo %f", resultado)
	}

}

func TestRestarExito(t *testing.T) {
	resultado, _ := Calcular(1, 1, "restar")

	if resultado != 0 {
		t.Errorf("esparado 0, pero se obtuvo %f", resultado)
	}
}

func TestMultiplicarExito(t *testing.T) {
	resultado, _ := Calcular(1, 1, "multiplicar")

	if resultado != 1 {
		t.Errorf("esparado 1, pero se obtuvo %f", resultado)
	}
}

func TestDividirExito(t *testing.T) {
	resultado, _ := Calcular(10, 10, "dividir")

	if resultado != 1 {
		t.Errorf("esparado 1, pero se obtuvo %f", resultado)
	}
}

func TestDividirErrorDividirPorCero(t *testing.T) {
	_, err := Calcular(1, 0, "dividir")

	if err != nil {
		t.Errorf("no podemos dividir por 0")
	}
}

func TestCalculoInvalido(t *testing.T) {
	_, err := Calcular(1, 0, "coso")

	if err != nil {
		t.Errorf("no podemos manejar coso")
	}
}
