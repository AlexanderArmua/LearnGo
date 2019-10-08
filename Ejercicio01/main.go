package main

import (
	"fmt"
	"os"
	"strconv"

	"./lib"
)

func main() {
	minimo, _ := strconv.Atoi(os.Args[1])
	maximo, _ := strconv.Atoi(os.Args[2])
	partes, _ := strconv.Atoi(os.Args[3])

	suma := lib.GetSumaDivisores(minimo, maximo)

	fmt.Printf("La suma de %d a %d con 1 routine es:\t%d\n", minimo, maximo, suma)

	suma3 := lib.GetSumaDivisoresParalelizando(minimo, maximo, partes)

	fmt.Printf("La suma de %d a %d con %d routines es:\t%d\n", minimo, maximo, partes, suma3)
}
