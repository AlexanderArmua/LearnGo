package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tenista struct {
	ProbabilidadAcierto float32
	Puntos              int
	Nombre              string
}

func (self *Tenista) Raquetazo(tenistaChannel chan bool) {
	// TODO: Evitar creer seed cada vez que llaman a la función.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	acerto := self.ProbabilidadAcierto > r1.Float32()

	tenistaChannel <- acerto
}

func main() {
	delPotro := new(Tenista)
	delPotro.ProbabilidadAcierto = 0.9
	delPotro.Nombre = "Del Potro"
	djokovic := new(Tenista)
	djokovic.ProbabilidadAcierto = 0.3
	djokovic.Nombre = "Djokovic"

	var ganadorTenisChannel chan string = make(chan string)

	go func() {
		nombreGanador := comenzarPartido(delPotro, djokovic)

		ganadorTenisChannel <- nombreGanador
	}()

	fmt.Println("Ganador:", <-ganadorTenisChannel)
}

func comenzarPartido(tenista1 *Tenista, tenista2 *Tenista) string {
	var tenistaChannel chan bool = make(chan bool)
	var puntosParaTerminarPartido int = 5
	var nombreGanador string

	for {
		go func() {
			tenista1.Raquetazo(tenistaChannel)

			if !<-tenistaChannel {
				tenista2.Puntos++
			}
		}()

		go func() {
			tenista2.Raquetazo(tenistaChannel)

			if !<-tenistaChannel {
				tenista1.Puntos++
			}
		}()

		// Calcular si alguien gano despues de cada raquetazo
		tenistaGanador, ganoAlguien := getTenistaGanador(tenista1, tenista2, puntosParaTerminarPartido)
		if ganoAlguien {
			nombreGanador = tenistaGanador.Nombre
			break
		}
	}

	return nombreGanador
}

func getTenistaGanador(tenista1 *Tenista, tenista2 *Tenista, puntosRequeridos int) (*Tenista, bool) {
	if tenista1.Puntos == puntosRequeridos {
		return tenista1, true
	} else if tenista2.Puntos == puntosRequeridos {
		return tenista2, true
	}

	return nil, false
}
