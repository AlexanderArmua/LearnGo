package lib

import "sync"

// GetSumaDivisores retorna la suma desde 0 hasta el numero que se ingrese de los multiplos de 3 o 5
func GetSumaDivisores(desde int, hasta int) int {
	return obtenerSumaDesdeHasta(desde, hasta)
}

// TODO: Por ahora solo toma numeros pares
// GetSumaDivisores2Paralelo igual que GetSumaDivisores pero lo paraleliza y realiza en menos tiempo
func GetSumaDivisores2Paralelo(desde int, hasta int) int {
	var wg sync.WaitGroup

	wg.Add(2)

	desde1 := desde
	hasta1 := hasta / 2
	desde2 := hasta1 + 1
	hasta2 := hasta

	var resultado1 int
	var resultado2 int

	go func() {
		defer wg.Done()
		resultado1 = obtenerSumaDesdeHasta(desde1, hasta1)
	}()

	go func() {
		defer wg.Done()
		resultado2 = obtenerSumaDesdeHasta(desde2, hasta2)
	}()

	wg.Wait()

	return resultado1 + resultado2
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
