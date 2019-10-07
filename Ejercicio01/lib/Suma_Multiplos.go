package lib

// GetSumaDivisores retorna la suma desde 0 hasta el numero que se ingrese de los multiplos de 3 o 5
func GetSumaDivisores(desde int, hasta int) int {
	return obtenerSumaDesdeHasta(desde, hasta)
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
