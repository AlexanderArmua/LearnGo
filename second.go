package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	params := getParamsAsInt()

	fmt.Printf("Ingreso por parametros: %v", params)
}

func getParamsAsInt() []int {
	parametrosConsola, error := getParametrosFromConsola()
	if !error {
		parametrosInt, err := convertirArrayStringToInt(parametrosConsola)
		if !err {
			return parametrosInt
		}
	}
	fmt.Printf("ERROR, hay parametros faltantes o que no son numeros")

	return []int{}
}

func convertirArrayStringToInt(arrayString []string) ([]int, bool) {
	arrayInt := make([]int, 0, len(arrayString))

	for i := 0; i < len(arrayString); i++ {
		numero, err := strconv.Atoi(arrayString[i])
		if err == nil {
			arrayInt = append(arrayInt, numero)
		} else {
			return []int{}, true
		}
	}

	return arrayInt, false
}

func getParametrosFromConsola() ([]string, bool) {
	if len(os.Args) > 1 {
		return os.Args[1:], false
	}
	return []string{}, true
}
