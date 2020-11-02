package contenidoWeb

import "fmt"

type ContenidoWeb struct {
	Multimedias []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {
	cantObjetos := len(cw.Multimedias)

	fmt.Println("\n\t***Mostrar objetos Multimedia***")
	if cantObjetos > 0 {
		for _, v := range cw.Multimedias {
			v.Mostrar()
			fmt.Println("----------------------------")
		}
	} else {
		fmt.Println("\nNo se encontraron objetos multimedia")
	}
}

func (cw *ContenidoWeb) CapturarImagen() {
	var titulo string
	var formato string
	var canales int64

	fmt.Println("\n\t***Agregar Imagen***")
	fmt.Println("Ingrese los siguientes datos de la Imagen:")
	fmt.Printf("Título: ")
	fmt.Scan(&titulo)
	fmt.Printf("Formato: ")
	fmt.Scan(&formato)
	fmt.Printf("Cantidad de canales: ")
	fmt.Scan(&canales)

	cw.AgregarImagen(titulo, formato, canales)
	fmt.Println("\nImagen capturada exitosamente")
}

func (cw *ContenidoWeb) CapturarAudio() {
	var titulo string
	var formato string
	var duracion float64

	fmt.Println("\n\t***Agregar Audio***")
	fmt.Println("Ingrese los siguientes datos del Audio:")
	fmt.Printf("Título: ")
	fmt.Scan(&titulo)
	fmt.Printf("Formato: ")
	fmt.Scan(&formato)
	fmt.Printf("Duración en segundos: ")
	fmt.Scan(&duracion)
	
	cw.AgregarAudio(titulo, formato, duracion)
	fmt.Println("\nAudio capturado exitosamente")
}

func (cw *ContenidoWeb) CapturarVideo() {
	var titulo string
	var formato string
	var frames int64
	fmt.Println("\n\t***Agregar Video***")
	fmt.Println("Ingrese los siguientes datos del Video:")
	fmt.Printf("Título: ")
	fmt.Scan(&titulo)
	fmt.Printf("Formato: ")
	fmt.Scan(&formato)
	fmt.Printf("Frames: ")
	fmt.Scan(&frames)

	cw.AgregarVideo(titulo, formato, frames)
	fmt.Println("\nVideo capturado exitosamente")
}

func (cw *ContenidoWeb) AgregarImagen(titulo string, formato string, canales int64) {
	i := Imagen{titulo, formato, canales}
	cw.Multimedias = append(cw.Multimedias, &i)
}

func (cw *ContenidoWeb) AgregarAudio(titulo string, formato string, duracion float64) {
	a := Audio{titulo, formato, duracion}
	cw.Multimedias = append(cw.Multimedias, &a)
}

func (cw *ContenidoWeb) AgregarVideo(titulo string, formato string, frames int64) {
	v := Video{titulo, formato, frames}
	cw.Multimedias = append(cw.Multimedias, &v)
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales int64
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion float64
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int64
}

func (i *Imagen) Mostrar() {
	fmt.Println("***Imagen***")
	fmt.Println("Título:", i.Titulo)
	fmt.Println("Formato:", i.Formato)
	fmt.Println("Canales:", i.Canales)
}

func (a *Audio) Mostrar() {
	fmt.Println("***Audio***")
	fmt.Println("Título:", a.Titulo)
	fmt.Println("Formato:", a.Formato)
	fmt.Println("Duración:", a.Duracion)
}

func (v *Video) Mostrar() {
	fmt.Println("***Video***")
	fmt.Println("Título:", v.Titulo)
	fmt.Println("Formato:", v.Formato)
	fmt.Println("Frames:", v.Frames)
}