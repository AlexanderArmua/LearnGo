/**
*	Hacer un programa que reciba por parámetro una lista de enteros,
*	los almacene en un slice,
*	popule un arbol binario y
*	después imprima los valores en orden ascendente.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Valor       int
	Left, Right *Node
}

func main() {
	params := getParamsAsInt()
	fmt.Printf("Ingreso por parametros: %v\n\n", params)

	nodo1 := new(Node)
	loadTree(nodo1, params)

	printTree(nodo1, 0, "*")
}

func loadTree(arbol *Node, elements []int) (*Node, bool) {
	if len(elements) == 0 {
		return new(Node), true
	}

	arbol.Valor = elements[0]

	for i := 1; i < len(elements); i++ {
		addValueToNode(arbol, elements[i])
	}

	return arbol, false
}

// Solo imprime hacia la derecha
func printTree(arbol *Node, pos int, char string) {
	if arbol != nil {
		fmt.Print(char + " -> ")

		fmt.Printf("%v", arbol.Valor)
		fmt.Print("\n")

		if arbol.Right != nil {
			printTree(arbol.Right, pos+1, char+"R.")
		}

		if arbol.Left != nil {
			printTree(arbol.Left, pos+1, char+"L.")
		}
	}
}

// Crear el arbol segun ingresa usuario (falta re-ordenar punteros)
func addValueToNode(arbol *Node, valor int) *Node {
	if arbol == nil {
		return createNewNode(valor)
	}

	addNodeToTree(arbol, valor)

	return arbol
}

func createNewNode(valor int) *Node {
	arbol := new(Node)
	arbol.Valor = valor
	return arbol
}

func addNodeToTree(arbol *Node, valor int) {
	if valor < arbol.Valor {
		arbol.Left = addValueToNode(arbol.Left, valor)
	} else if valor > arbol.Valor {
		arbol.Right = addValueToNode(arbol.Right, valor)
	}
}

// PUNTO 1, 2 y 3
func getParamsAsInt() []int {
	parametrosConsola, error := getParametrosFromConsolaAsInt()

	if !error {
		return parametrosConsola
	}

	fmt.Printf("ERROR, hay parametros faltantes o que no son numeros")

	return []int{}
}

func getParametrosFromConsolaAsInt() ([]int, bool) {
	if len(os.Args) > 1 {
		sliceRetornar := make([]int, len(os.Args), 2*len(os.Args))
		sliceRetornar = convertirArrayStringToInt(os.Args[1:])

		return sliceRetornar, false
	}

	return []int{}, true
}

func convertirArrayStringToInt(arrayString []string) []int {
	arrayInt := make([]int, 0, len(arrayString))

	for i := 0; i < len(arrayString); i++ {
		numero, err := strconv.Atoi(arrayString[i])
		if err == nil {
			arrayInt = append(arrayInt, numero)
		} else {
			fmt.Printf("ERROR AL CONVERTIR")
		}
	}

	return arrayInt
}

/*
func getParametrosFromConsola() ([]string, bool) {
	if len(os.Args) > 1 {
		return os.Args[1:], false
	}
	return []string{}, true
}
*/
