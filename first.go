// "go command"
// Clase en ML el día: 11/09/2019
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
}

// Se ejecuta todo aca al iniciar
func init() {
	valor, error := strconv.Atoi(os.Args[1])
	if error == nil {
		pedirNumeroYMostrar(valor)
	} else {
		fmt.Printf("ERROR: '%s' no es un valor entero", os.Args[1])
	}

}

// HelloWorld (public)
func HelloWorld() {
	mensaje := "Hello World 3"
	fmt.Println(mensaje)

	const otroMensaje string = "Mensaje 2222"
	fmt.Printf(otroMensaje)
}

func printZeroValue() {
	var c complex128
	fmt.Printf("\n\t%v\n", c)
}

func ifearUnNombreAlgo(parametro string) {
	if parametro == "Alexander" {
		fmt.Println("Que acelga")
	} else {
		fmt.Println("Buenas tardes")
	}
}

func loopearTuO(veces int) {
	var o string
	for i := 0; i < veces; i++ {
		o += "o"
		fmt.Print("Hell")
		fmt.Println(o)
	}
}

// fail
func loopWhileMultiplo(multiplo int) {
	i := 1
	for {
		if i%multiplo == 0 {
			break
		}

		println("Vuelta nº: %s", i)

		i++
	}
}

func imprimirAlReves(x, y string) (string, int, otro string) {
	otro = "asdasda"
	return y, x, otro
}

func pedirNumeroYMostrar(numero int) {
	//for i := 1; i < len(os.Args); i++ {
	//	fmt.Println(os.Args[i])
	//}
	valor := obtenerSuma(numero)
	fmt.Printf("La suma es: %d", valor)
}

func obtenerSuma(p int) (suma int) {
	for v := 0; v <= p; v++ {
		if esMultiploDe3o5(v) {
			suma += v
		}
	}

	return suma
}

func esMultiploDe3o5(valor int) bool {
	return esMultiploDe(valor, 3) || esMultiploDe(valor, 5)
}

func esMultiploDe(valor, divisor int) bool {
	return valor%divisor == 0
}
