package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Nota struct {
	ID        int
	Contenido string
	Fecha     time.Time
}

var idActual int = 0
var notas []Nota

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

func crearNota() Nota {
	var contenidoNota string
	idActual++
	fmt.Println("Introduce el contenido de la nota:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		contenidoNota = scanner.Text()
	}
	nuevaNota := Nota{
		ID:        idActual,
		Contenido: contenidoNota,
		Fecha:     time.Now(),
	}
	fmt.Println("Nota creada correctamente")
	return nuevaNota
}

func editarNota() {
	var contenidoNuevoNota string
	var ID int
	consultarNotas()
	fmt.Println("¿Qué nota quieres editar?")
	fmt.Scan(&ID)
	notaEncontrada := encontrarNotaPorID(ID)
	fmt.Println("Cambia el contenido de la nota:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		contenidoNuevoNota = scanner.Text()
	}
	notaEncontrada.Contenido = contenidoNuevoNota
	fmt.Println("Nota editada correctamente")
}

func consultarNotas() {
	for _, nota := range notas {
		fmt.Printf("ID: %d, Contenido: %s, Fecha: %s\n", nota.ID, nota.Contenido, nota.Fecha)
	}
}

func borrarNota() []Nota {
	var ID int
	consultarNotas()
	fmt.Println("¿Qué nota quieres borrar?")
	fmt.Scan(&ID)
	var nuevasNotas []Nota
	for _, nota := range notas {
		if nota.ID != ID {
			nuevasNotas = append(nuevasNotas, nota)
		}
	}
	fmt.Println("Nota borrada correctamente")
	return nuevasNotas
}

func opcionIncorrecta() {
	fmt.Println("No has introducido una opción correcta.")
}

func menu(opcionSeleccionada int) {
	switch opcionSeleccionada {
	case 1:
		notaCreada := crearNota()
		notas = append(notas, notaCreada)
	case 2:
		editarNota()
	case 3:
		consultarNotas()
	case 4:
		notas := borrarNota()
		cantidadNotas := len(notas)
		fmt.Println("Notas en total:", cantidadNotas)
	default:
		opcionIncorrecta()
	}
}

func encontrarNotaPorID(id int) *Nota {
	for i := range notas {
		if notas[i].ID == id {
			return &notas[i]
		}
	}
	return nil
}
