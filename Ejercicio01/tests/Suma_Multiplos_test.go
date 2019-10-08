package main_test

import (
	"fmt"
	"testing"

	"../lib"
)

func TestGetSumaDivisores(t *testing.T) {
	t.Run("0 a 10", func(t *testing.T) {
		got := lib.GetSumaDivisores(0, 10)
		want := 33

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("-5 a 5", func(t *testing.T) {
		got := lib.GetSumaDivisores(-5, 5)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestBenchmarkGetSumaDivisoresParalelizando(t *testing.T) {
	t.Run("0 a 10", func(t *testing.T) {
		got := lib.GetSumaDivisoresParalelizando(0, 10, 2)
		want := 33

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("-5 a 5", func(t *testing.T) {
		got := lib.GetSumaDivisoresParalelizando(-5, 5, 2)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

/*
	Nota: Como se puede ver en los benchmark.
	Aproximadamente de 0 a 10M tarda lo mismo que la suma de (0 a 5M) + (5M hasta 10M),
	por lo tanto, si se hicieran en paralelo podria cortarse considerablemente el tiempo
	de ejecución.
*/
func BenchmarkGetSumaDivisores(b *testing.B) {
	fmt.Println("\t1 Routines")
	want1 := 23333341666668
	want2 := 1458335416668
	want3 := 5833334166668
	want4 := 17500012500000

	for i := 0; i < b.N; i++ {
		calcSumaDivisores(0, 10000000, 1, want1, b)
		calcSumaDivisores(0, 2500000, 1, want2, b)
		calcSumaDivisores(0, 5000000, 1, want3, b)
		calcSumaDivisores(5000000, 10000000, 1, want4, b)
	}
}

func BenchmarkGetSumaDivisoresParalelizando(b *testing.B) {
	iterateAndCalculate(4, b)
	iterateAndCalculate(5, b)
	iterateAndCalculate(8, b)
	iterateAndCalculate(11, b)
	iterateAndCalculate(16, b)
	iterateAndCalculate(23, b)
	iterateAndCalculate(32, b)
}

func iterateAndCalculate(routines int, b *testing.B) {
	want1 := 23333341666668
	want2 := 1458335416668
	want3 := 5833334166668
	want4 := 17500012500000

	fmt.Printf("\t%d Routines\n", routines)
	for i := 0; i < b.N; i++ {
		calcSumaDivisores(0, 10000000, routines, want1, b)
		calcSumaDivisores(0, 2500000, routines, want2, b)
		calcSumaDivisores(0, 5000000, routines, want3, b)
		calcSumaDivisores(5000000, 10000000, routines, want4, b)
	}
}

func calcSumaDivisores(desde int, hasta int, routines int, want int, b *testing.B) {
	b.Run(fmt.Sprintf("Desde: %d Hasta: %d Routines: %d", desde, hasta, routines), func(b *testing.B) {
		got := lib.GetSumaDivisoresParalelizando(desde, hasta, routines)
		if got != want {
			b.Errorf("got %d want %d", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				lib.GetSumaDivisoresParalelizando(desde, hasta, routines)
			}
		},
	)
}

func ExampleGetSumaDivisores() {
	suma := lib.GetSumaDivisores(0, 10000)

	fmt.Printf("La suma de 0 a 10.000 es: %d", suma)
}
