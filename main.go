package main

import (
	"fmt"
	"golang-curso/calculadora"
	"golang-curso/condicionales"
	listatareas "golang-curso/lista-tareas"
)

func main() {
	nombre := "Carlos"
	edad := 30
	altura := 1.82
	activo := true
	const PI = 3.14159
	const NombreApp = "MiApp"

	fmt.Println(nombre, edad, altura, activo, PI, NombreApp)

	if edad < 29 {
		fmt.Println("es menor")
	} else {
		fmt.Println("es mayor")
	}

	if n := len(nombre); n < 5 {
		fmt.Println("tiene menos de 5 letras de largo el nombre Carlos, pero con if con declaracion")
	} else {
		fmt.Println("tiene menos de 5 letras de largo el nombre Carlos, pero con if con declaracion")
	}

	dia := 3

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	}

	switch {
	case edad < 18:
		fmt.Println("eres menor de edad")
	case edad > 18:
		fmt.Println("eres mayor de edad")
	}

	switch dia := 6; dia {
	case 1, 2, 3, 4, 5:
		fmt.Println("dia de semana")
	case 6, 7:
		fmt.Println("fin de semana")
	}

	for i := 1; i < 6; i++ {
		fmt.Println("vamos en el numero", i)
	}

	// iteracion for como while

	i := 0

	for i < 3 {
		fmt.Println("iteracion numero", i)
		i++
	}

	// for infinito

	for {
		fmt.Println("Esto corre para siempre")
		break
	}

	// for range

	nombres := []string{"Ana", "Luis", "Pedro"}

	for i, nombre := range nombres {
		fmt.Println(i, nombre)
	}

	// for range si no necesito indice lo reemplazo por _

	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	edades := map[string]int{
		"Ana":  20,
		"Luis": 25,
	}

	for nombre, edad := range edades {
		fmt.Println(nombre, edad)
	}

	condicionales.CalcularMadurez(13)
	resultado, err := calculadora.Calcular(1, 1, "dividir")

	if err != nil {
		fmt.Println("Ocurrió un error:", err)
		return
	}
	fmt.Println(resultado)

	fmt.Println(listatareas.CrearTarea("lavar"))
	fmt.Println(listatareas.CrearTarea("estudiar"))
	listatareas.Listar()
	completarTarea := listatareas.CompletarTarea(1)
	if completarTarea == nil {
		fmt.Println("tarea completada")
	}

	eliminarTarea := listatareas.EliminarPorIndice(1)

	if eliminarTarea == nil {
		fmt.Println("tarea eliminada")
	}

	listatareas.CargarDesdeArchivo("tareas.json")

	listatareas.CrearTarea("comer")
	listatareas.CrearTarea("aspirar")

	listatareas.GuardarEnArchivo("tareas.json")

	listatareas.Listar()
}
