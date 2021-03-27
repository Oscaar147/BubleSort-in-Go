package main

import "fmt"

type Imagen struct {
	titulo  string
	formato string
	canales string
}
type Audio struct {
	titulo   string
	formato  string
	duracion int64
}
type Video struct {
	titulo  string
	formato string
	frames  int64
}

type Multimedia interface {
	Mostrar()
}

type ContenidoWeb struct {
	M []Multimedia
}

func (i *Imagen) Mostrar() {
	fmt.Println("\nIMAGEN")
	fmt.Printf("Titulo: %s\n", i.titulo)
	fmt.Printf("Formato: %s\n", i.formato)
	fmt.Printf("Canales: %s\n", i.canales)
}

func (a *Audio) Mostrar() {
	fmt.Println("\nAUDIO")
	fmt.Printf("Titulo: %s\n", a.titulo)
	fmt.Printf("Formato: %s\n", a.formato)
	fmt.Printf("Duracion: %d\n", a.duracion)
}

func (v *Video) Mostrar() {
	fmt.Println("\nVIDEO")
	fmt.Printf("Titulo: %s\n", v.titulo)
	fmt.Printf("Formato: %s\n", v.formato)
	fmt.Printf("Frames: %d\n", v.frames)
}

func main() {
	terminar := false
	c := ContenidoWeb{}

	for !terminar {
		MenuPrincipal()
		op := OpcionDataUser()
		switch op {
		case 1:
			imprimirMenuAñadir()
			contenido := OpcionDataUser()
			if contenido == 1 {
				c.M = append(c.M, añadirImagen())
			} else if contenido == 2 {
				c.M = append(c.M, añadirAudio())

			} else {
				c.M = append(c.M, añadirVideo())
			}
			fmt.Println("Agregado Correctamente!\n\n")
			break
		case 2:
			mostrar(c.M...)
			break
		case 3:
			terminar = true
			break
		default:
			fmt.Print("Opcion incorrecta intenta de nuevo")
			break
		}
	}
}

func MenuPrincipal() {
	fmt.Println("\nBienvenido")
	fmt.Println("1. Agregar multimedia")
	fmt.Println("2. Mostrar multimedia")
	fmt.Println("3. Salir")
	fmt.Print("Opcion: ")
}

func imprimirMenuAñadir() {
	fmt.Print("\t1. Imagen\n")
	fmt.Print("\t2. Audio\n")
	fmt.Print("\t3. Video\n")
	fmt.Print("\tOpcion: ")
}

func OpcionDataUser() int64 {
	var opc int64
	fmt.Scan(&opc)
	return opc
}

func añadirImagen() *Imagen {
	var titulo, formato, canales string
	fmt.Printf("\nAgregar una nueva imagen\n")
	fmt.Print("Titulo: ")
	titulo = GetUserInfo()
	fmt.Print("Formato: ")
	formato = GetUserInfo()
	fmt.Print("Canales: ")
	canales = GetUserInfo()
	return &Imagen{titulo, formato, canales}
}

func añadirAudio() *Audio {
	var titulo, formato string
	var duracion int64

	fmt.Printf("\nAgregar un audio\n")
	fmt.Print("Titulo: ")
	titulo = GetUserInfo()
	fmt.Print("Formato: ")
	formato = GetUserInfo()
	fmt.Print("Duracion: ")
	duracion = OpcionDataUser()
	return &Audio{titulo, formato, duracion}
}

func añadirVideo() *Video {
	var titulo, formato string
	var frames int64
	fmt.Printf("\nAgregar un video\n")
	fmt.Print("Titulo: ")
	titulo = GetUserInfo()
	fmt.Print("Formato: ")
	formato = GetUserInfo()
	fmt.Print("Frames: ")
	frames = OpcionDataUser()
	return &Video{titulo, formato, frames}
}

func mostrar(contenido ...Multimedia) {
	for _, c := range contenido {
		c.Mostrar()
	}

}
func GetUserInfo() string {
	var line string
	fmt.Scan(&line)
	return line
}
