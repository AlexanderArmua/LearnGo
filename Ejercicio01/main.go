package main

import (
	"fmt"

	"./lib"
)

func main() {
	suma := lib.GetSumaDivisores(0, 10000000)

	fmt.Printf("La suma de 0 a 10000000 es: %d", suma)
}
