package main

import (
	"fmt"
	"./contenidoWeb"
)

const opcAgregarObjMultimedia int = 1
const opcMostrarObjetosMultimedia int = 2
const opcionSalir int = 3
const opcImagen int = 1
const opcAudio int = 2
const opcVideo int = 3
const opcRegresar int = 4

func main() {
	cw := new(contenidoWeb.ContenidoWeb)
	var opcionMenu int
	var opcMultimedia int
	continuarPrograma := true
	regresarMenu := false

	for continuarPrograma {
		fmt.Println("\n\t***Contenido Web***")
		fmt.Println("\n1. Agregar objeto multimedia")
		fmt.Println("2. Mostrar objetos multimedia")
		fmt.Println("3. Salir")
		fmt.Printf("Ingrese una opción: ")
		fmt.Scan(&opcionMenu)

		switch opcionMenu {
		case opcAgregarObjMultimedia:
			fmt.Println("\n\n\t***Agregar un objeto multimedia***")
			fmt.Println("1. Imagen")
			fmt.Println("2. Audio")
			fmt.Println("3. Video")
			fmt.Println("4. Regresar")
			fmt.Printf("Ingrese la opción del objeto multimedia que desea agregar: ")
			fmt.Scan(&opcMultimedia)

			switch opcMultimedia {
			case opcImagen:
				cw.CapturarImagen()
			
			case opcAudio:
				cw.CapturarAudio()
			
			case opcVideo:
				cw.CapturarVideo()
			
			case opcRegresar:
				regresarMenu = true
			
			default:
				fmt.Println("\nOpción no válida")
			}

		case opcMostrarObjetosMultimedia:
			cw.Mostrar()

		case opcionSalir:
			fmt.Println("\nGracias por utilizar el programa...")
			continuarPrograma = false
		
		default:
			fmt.Println("\nOpción no válida")
		}

		if continuarPrograma {
			if regresarMenu {
				regresarMenu = false
			} else {
				fmt.Println("\nPresione enter para continuar...")
				fmt.Scanf("\n\n")
			}
		}
	}
}