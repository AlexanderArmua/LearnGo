package main

import (
	"fmt"

	"./lib"
)

func main() {
	suma := lib.GetSumaDivisores(0, 10000000)

	fmt.Printf("La suma de 0a 10M es: %d", suma)

	suma2 := lib.GetSumaDivisores2Paralelo(0, 10000000)

	fmt.Printf("La suma de 0a 10M es: %d", suma2)
}
