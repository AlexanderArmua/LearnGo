package lib

import (
	"runtime"
	"sync"
)

type Parte struct {
	Inicio int
	Fin    int
}

// GetSumaDivisores retorna la suma desde 0 hasta el numero que se ingrese de los multiplos de 3 o 5
func GetSumaDivisores(desde int, hasta int) int {
	return obtenerSumaDesdeHasta(desde, hasta)
}

// GetSumaDivisoresParalelizandoCPU retorna la suma desde 0 hasta el numero creando tantas goroutines como pueda
func GetSumaDivisoresParalelizandoCPU(desde int, hasta int) int {
	processors := runtime.GOMAXPROCS(0)
	return GetSumaDivisoresParalelizando(desde, hasta, processors)
}

// GetSumaDivisoresParalelizando retorna la suma desde 0 hasta el numero pero se puede decirle en cuantas partes paralelizar
func GetSumaDivisoresParalelizando(desde int, hasta int, partes int) int {
	var wg sync.WaitGroup

	if partes == 0 {
		partes = 1
	}

	wg.Add(partes)

	partesTomar := armarSlicePartes(desde, hasta, partes)

	resultados := make([]int, 0, partes)
	for i := 0; i < len(partesTomar); i++ {
		go func(value Parte) {
			defer wg.Done()

			resultados = append(resultados, obtenerSumaDesdeHasta(value.Inicio, value.Fin))
		}(partesTomar[i])
	}

	wg.Wait()

	return getSumaFromSlice(resultados)
}

func obtenerSumaDesdeHasta(desde int, hasta int) (suma int) {
	divisores := []int{3, 5}
	for v := desde; v <= hasta; v++ {
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

func armarSlicePartes(desde int, hasta int, partes int) []Parte {
	/*
		Crea todas las partes en las que se va a trabajar.
		Ejemplo: 0 a 100 en 4 partes
		{0 a 25}, {26 a 50}, {51 a 75}, {76 a 100}
	*/
	tamanoParte := (hasta - desde) / partes

	partesTomar := make([]Parte, 0, partes)
	partesTomar = append(partesTomar, Parte{Inicio: desde, Fin: desde + tamanoParte})

	for i := 1; i < partes; i++ {
		nuevoInicio := partesTomar[i-1].Fin + 1
		nuevoFin := partesTomar[i-1].Fin + tamanoParte
		partesTomar = append(partesTomar, Parte{Inicio: nuevoInicio, Fin: nuevoFin})
	}

	fixOverflowUltimoElemento(partesTomar, hasta)

	return partesTomar
}

/*
	Soluciona problemas cuando la division no da numeros exactos. Por ej:
	Si el numero ingresado es impar y se divide por 4.
	Las partes quedan {0, 24}; +24{25, 48}; +24{49, 72}; +24{73, 96};
	96 - (96 - 97) = 97
*/
func fixOverflowUltimoElemento(partesTomar []Parte, maximoPermitido int) {
	posUltimoElemento := len(partesTomar) - 1
	partesTomar[posUltimoElemento].Fin = partesTomar[posUltimoElemento].Fin - (partesTomar[posUltimoElemento].Fin - maximoPermitido)
}

func getSumaFromSlice(numeros []int) int {
	sum := 0
	for _, i := range numeros {
		sum += i
	}

	return sum
}
