/*
	Modelar la funcionalidad para un sistema de cines
	que calcule los ingresos netos de una función en base a los asistentes
	y al precio base de la entrada.
	Existen tres tipos de asistentes y tienen las siguientes características:

	General: Pagan el 100%.
	Jubilados: Tienen un 50% de descuento.
	Invitados: No pagan nada.
*/

package main

import "fmt"

type general struct {
}

func (a general) Pagar(valor float32) float32 {
	return valor
}

type jubilado struct {
}

func (a jubilado) Pagar(valor float32) float32 {
	return 0.5 * valor
}

type invitado struct {
}

func (a invitado) Pagar(valor float32) float32 {
	return 0
}

type asistente interface {
	Pagar(float32) float32
}

func main() {
	// TODO: Seria bueno que no este hardcodeado y que al menos tome por parametro al ejecutar
	gente := []asistente{general{}, general{}, jubilado{}, jubilado{}, invitado{}}
	var precioBase float32 = 20.55

	ingresoNeto := float32(0)
	for _, elem := range gente {
		ingresoNeto += elem.Pagar(precioBase)
	}

	fmt.Printf("\n\tGanancia Neta: %v", ingresoNeto)
}
