package main

import "fmt"

/*
	1. Es una direccion de memoria.
	2. En lugar de valor, tengo la direccion en donde esta el valor.

	Ej:
	*T => *int, *string, *Struct

	Valor zero de un puntero => nil
*/
func main() {

	// Puntero a un entero
	var x, y *int
	enteroNuevo := 5

	// Puntero X apunta a la direccion de memoria de enteroNuevo
	x = &enteroNuevo
	y = &enteroNuevo

	fmt.Println(x)
	fmt.Println(y)

	// Modificamos el valor contenido en la direccion de memoria a la que apunta x
	*x = 6

	// Mostramos el valor de la direccion de memoria a la que apunta x
	fmt.Println(*x)
	fmt.Println(*y)
}
