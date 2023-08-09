package main

import (
	"fmt"
	"time"
)

type Nota struct {
	ID        int
	Contenido string
	Fecha     time.Time
}

func main() {
	for true {

		var opcion int

		fmt.Println("Bienvenido a la aplicación de notas")
		fmt.Println("-----------------------------------")

		fmt.Println("Crear una nota (1)")
		fmt.Println("Editar una nota (2)")
		fmt.Println("Consultar notas (3)")
		fmt.Println("Borrar una nota (4)")
		fmt.Println("Salir (5)")
		fmt.Println("Selecciona una de las siguientes opciones del 1 al 5:")
		fmt.Scan(&opcion)

		if opcion >= 1 && opcion <= 4 {
			menu(opcion)
		} else if opcion == 5 {
			break
		} else {
			opcionIncorrecta()
		}
	}

	fmt.Println("¡Hasta luego!")
}

func crearNota() {

	fmt.Println("Función crearNota")
}

func editarNota() {
	fmt.Println("Función editarNota")
}

func consultarNotas() {
	fmt.Println("Función consultarNotas")
}

func borrarNota() {
	fmt.Println("Función borrarNota")
}

func opcionIncorrecta() {
	fmt.Println("No has introducido una opción correcta.")
}

func menu(opcionSeleccionada int) {
	switch opcionSeleccionada {
	case 1:
		crearNota()
	case 2:
		editarNota()
	case 3:
		consultarNotas()
	case 4:
		borrarNota()
	default:
		opcionIncorrecta()
	}
}
