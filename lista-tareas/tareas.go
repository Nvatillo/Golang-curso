package listatareas

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Tarea struct {
	ID     int
	Titulo string
	Estado string
}

var tareas []Tarea

func CrearTarea(nombre string) string {
	nueva := Tarea{
		ID:     len(tareas),
		Titulo: nombre,
		Estado: "pendiente",
	}
	tareas = append(tareas, nueva)
	return "creada"
}

func Listar() {
	for _, tarea := range tareas {
		fmt.Println(tarea.ID, tarea.Titulo, tarea.Estado)
	}
}

func CompletarTarea(id int) error {
	for i := range tareas {
		if tareas[i].ID == id {
			tareas[i].Estado = "completada"
			return nil
		}
	}

	return errors.New("no existe la tarea con ese ID")
}

func EliminarPorIndice(i int) error {
	if i < 0 || i >= len(tareas) {
		return errors.New("índice fuera de rango")
	}

	tareas = append(tareas[:i], tareas[i+1:]...)
	return nil
}

func GuardarEnArchivo(ruta string) error {
	data, err := json.MarshalIndent(tareas, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(ruta, data, 0644)
}

func CargarDesdeArchivo(ruta string) error {
	data, err := os.ReadFile(ruta)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &tareas)
}
