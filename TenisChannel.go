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

func (self *Tenista) Raquetazo() bool {
	// TODO: Evitar creer seed cada vez que llaman a la función.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return self.ProbabilidadAcierto > r1.Float32()
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

	fmt.Print(<-ganadorTenisChannel)
}

func jugarLosTantos(t1 *Tenista, t2 *Tenista, c chan *Tenista) {
	var puntosParaTerminarPartido int = 5

	tenistas := [2]*Tenista{t1, t2}

	for i := 0; i < (2 * puntosParaTerminarPartido); i++ {
		nroTenista := i % 2

		if !tenistas[nroTenista].Raquetazo() {
			c <- tenistas[(i+1)%2]
		}
	}

	// Borrar
	close(c)
}

func comenzarPartido(tenista1 *Tenista, tenista2 *Tenista) string {
	var tenistaChanel = make(chan *Tenista)

	retornar := ""
	go jugarLosTantos(tenista1, tenista2, tenistaChanel)

	for {
		tenistaQueSumaPuntos, ok := <-tenistaChanel

		if ok == false {
			// Calcular quien se pego mas raquetazos
			tenistaGanador, ganoAlguien := getTenistaGanador(tenista1, tenista2)
			if ganoAlguien {
				retornar = tenistaGanador.Nombre
				break
			}
		} else {
			tenistaQueSumaPuntos.Puntos++
		}
	}

	return retornar
}

func getTenistaGanador(tenista1 *Tenista, tenista2 *Tenista) (*Tenista, bool) {
	if tenista1.Puntos > tenista2.Puntos {
		return tenista1, true
	} else if tenista2.Puntos > tenista1.Puntos {
		return tenista2, true
	}

	return nil, false
}
