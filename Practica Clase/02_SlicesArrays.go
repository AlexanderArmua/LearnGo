/*
	Slices: Estructuras basadas en arrays que se pueden redimensionar.
*/

package main

import (
	"fmt"
)

func main() {
	////////// ARRAY
	var array1 [3]int // Array de tamaño 3, vacio
	fmt.Println(array1)

	array2 := [3]int{1, 2, 3} // Array de tamaño 3 con contenido
	fmt.Println(array2)

	////////// SLICE
	slice1 := []int{1, 2, 3, 4} // Slice de tamaño 4
	fmt.Println(slice1)

	slice2 := slice1[:2] // Nuevo Slice desde 0 hasta 2
	fmt.Println(slice2)

	slice3 := []string{"A", "B", "C", "D"}

	slice4 := slice3[1:2] // Arranca en 1 (inclusive) hasta 2 (no-inclusive)
	fmt.Println(slice4)

	////////// MAKE
	slice5 := make([]int, 3, 5) // Crea un slice de len 3 y capacidad maxima 5
	fmt.Println(slice5)
	fmt.Println(cap(slice5)) // Capacidad del Slice

	// Añade un nuevo elemento sin la necesidad de reconstruir un nuevo slice
	slice5 = append(slice5, 27)
	fmt.Println(slice5)

	// x2
	slice5 = append(slice5, 22)
	fmt.Println(slice5)

	// Desborda su capacidad y se vuelve menos eficiente, de array pasa a slice
	slice5 = append(slice5, 11)
	fmt.Println(slice5)

	////////// COPY
	// COPIA EL MINIMO DE ELEMENTOS EN AMBOS ARRAYS
	slice6 := []int{15, 25, 35}
	slice7Copy := make([]int, 3)
	slice8CopyChico := make([]int, 2)

	// Funciona porque los dos son de tamaño 3
	copy(slice7Copy, slice6) // Copia el array 2 al 1

	// Solo copia las posiciones que posean los dos.
	/*
		copy([ , ,]. [1, 2, 3]) ==> [1, 2]
	*/
	copy(slice8CopyChico, slice6)

	fmt.Println(slice6)
	fmt.Println(slice7Copy)
	fmt.Println(slice8CopyChico)

	// Nos aseguramos que el que sea copia siempre tenga el mismo tamaño
	slice9Copy := make([]int, len(slice6))
	copy(slice9Copy, slice6)

	fmt.Println(slice9Copy)
}
