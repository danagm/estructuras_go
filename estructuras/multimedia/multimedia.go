package multimedia

import "fmt"

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

func (i *Imagen) Mostrar() {
	fmt.Println("Imagen")
	fmt.Println("Título:", i.Titulo, "Formato:", i.Formato, "Canales:", i.Canales)
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int
}

func (a *Audio) Mostrar() {
	fmt.Println("Audio:")
	fmt.Println("Título:", a.Titulo, "Formato:", a.Formato, "Duración:", a.Duracion)
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int
}

func (v *Video) Mostrar() {
	fmt.Println("Video:")
	fmt.Println("Título:", v.Titulo, "Formato:", v.Formato, "Frames:", v.Frames)
}

type Archivo interface {
	Mostrar()
}

type ArchivoMultimedia struct {
	Archivos []Archivo
}

func (am *ArchivoMultimedia) Mostrar() {
	for _, a := range am.Archivos {
		a.Mostrar()
	}
}
