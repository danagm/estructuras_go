package main

import (
	"fmt"

	"./multimedia"
)

func main() {
	am := multimedia.ArchivoMultimedia{
		Archivos: []multimedia.Archivo{},
	}

	var opc int32
	var opc1 int32
	for opc != 3 {
		fmt.Println("\n1.Ingresar Contenido Web\n2.Mostrar Contenido Web\n3.Terminar programa\n: ")
		fmt.Scanln(&opc)
		if opc == 1 {
			for opc1 != 4 {
				fmt.Println("\n1.Imagen\n2.Video\n3.Audio\n4.Regresar\n: ")
				fmt.Scanln(&opc1)
				if opc1 == 1 {
					var titulo, formato, canales string
					fmt.Println("-IMAGEN-")
					fmt.Print("titulo: ")
					fmt.Scanln(&titulo)
					fmt.Print("formato: ")
					fmt.Scanln(&formato)
					fmt.Print("canales: ")
					fmt.Scanln(&canales)
					img := multimedia.Imagen{titulo, formato, canales}
					am.Archivos = append(am.Archivos, &img)
				} else if opc1 == 2 {
					var titulo, formato string
					var frames int
					fmt.Println("-VIDEO-")
					fmt.Print("titulo: ")
					fmt.Scanln(&titulo)
					fmt.Print("formato: ")
					fmt.Scanln(&formato)
					fmt.Print("frames: ")
					fmt.Scanln(&frames)
					vid := multimedia.Video{titulo, formato, frames}
					am.Archivos = append(am.Archivos, &vid)
				} else if opc1 == 3 {
					var titulo, formato string
					var duracion int
					fmt.Println("-AUDIO-")
					fmt.Print("titulo: ")
					fmt.Scanln(&titulo)
					fmt.Print("formato: ")
					fmt.Scanln(&formato)
					fmt.Print("duración: ")
					fmt.Scanln(&duracion)
					aud := multimedia.Video{titulo, formato, duracion}
					am.Archivos = append(am.Archivos, &aud)
				}
			}
		} else if opc == 2 {
			am.Mostrar()
		}
	}
}
