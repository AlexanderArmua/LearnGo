/*
	go run NombreArchivo.go NUMERO
	Se imprime la suma de todos los numeros multiplos de 3 y de 5 desde 0 hasta NUMERO.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

}

func init() {
	if len(os.Args) > 1 {
		valor, error := strconv.Atoi(os.Args[1])
		if error == nil {
			pedirNumeroYMostrar(valor)
		} else {
			fmt.Printf("ERROR: '%s' no es un valor entero", os.Args[1])
		}
	} else {
		fmt.Printf("ERROR: Ingrese un parametro")
	}
}

func pedirNumeroYMostrar(numero int) {
	valor := obtenerSuma(numero)
	fmt.Printf("La suma es: %d", valor)
}

func obtenerSuma(p int) (suma int) {
	divisores := []int{3, 5}
	for v := 0; v <= p; v++ {
		if esMultiploDeAlguno(v, divisores) {
			suma += v
		}
	}

	return suma
}

func esMultiploDeAlguno(valor int, valores []int) bool {
	for i := 0; i < len(valores); i++ {
		if esMultiploDe(valor, valores[i]) {
			return true
		}
	}
	return false
}

func esMultiploDe(valor, divisor int) bool {
	return valor%divisor == 0
}
