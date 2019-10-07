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

/*
	Nota: Como se puede ver en los benchmark.
	Aproximadamente de 0 a 10M tarda lo mismo que la suma de (0 a 5M) + (5M hasta 10M),
	por lo tanto, si se hicieran en paralelo podria cortarse considerablemente el tiempo
	de ejecución.
*/
func BenchmarkGetSumaDivisores(b *testing.B) {
	want1 := 23333341666668
	want2 := 1458335416668
	want3 := 5833334166668
	want4 := 17500012500000

	for i := 0; i < b.N; i++ {
		b.Run("0 a 10M", func(b *testing.B) {
			got := lib.GetSumaDivisores(0, 10000000)
			if got != want1 {
				b.Errorf("got %d want %d", got, want1)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores(0, 10000000)
				}
			},
		)

		b.Run("0 a 2.5M", func(b *testing.B) {
			got := lib.GetSumaDivisores(0, 2500000)
			if got != want2 {
				b.Errorf("got %d want %d", got, want2)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores(0, 2500000)
				}
			},
		)

		b.Run("0 a 5M", func(b *testing.B) {
			got := lib.GetSumaDivisores(0, 5000000)
			if got != want3 {
				b.Errorf("got %d want %d", got, want3)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores(0, 5000000)
				}
			},
		)

		b.Run("5M a 10M", func(b *testing.B) {
			got := lib.GetSumaDivisores(5000000, 10000000)
			if got != want4 {
				b.Errorf("got %d want %d", got, want4)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores(5000000, 10000000)
				}
			},
		)
	}
}

func BenchmarkGetSumaDivisores2Paralelo(b *testing.B) {
	want1 := 23333341666668
	want2 := 1458335416668
	want3 := 5833334166668
	want4 := 17500012500000

	for i := 0; i < b.N; i++ {
		b.Run("0 a 10M", func(b *testing.B) {
			got := lib.GetSumaDivisores2Paralelo(0, 10000000)
			if got != want1 {
				b.Errorf("got %d want %d", got, want1)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores2Paralelo(0, 10000000)
				}
			},
		)

		b.Run("0 a 2.5M", func(b *testing.B) {
			got := lib.GetSumaDivisores2Paralelo(0, 2500000)
			if got != want2 {
				b.Errorf("got %d want %d", got, want2)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores2Paralelo(0, 2500000)
				}
			},
		)

		b.Run("0 a 5M", func(b *testing.B) {
			got := lib.GetSumaDivisores2Paralelo(0, 5000000)
			if got != want3 {
				b.Errorf("got %d want %d", got, want3)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores2Paralelo(0, 5000000)
				}
			},
		)

		b.Run("5M a 10M", func(b *testing.B) {
			got := lib.GetSumaDivisores2Paralelo(5000000, 10000000)
			if got != want4 {
				b.Errorf("got %d want %d", got, want4)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores2Paralelo(5000000, 10000000)
				}
			},
		)
	}
}

func ExampleGetSumaDivisores() {
	suma := lib.GetSumaDivisores(0, 10000)

	fmt.Printf("La suma de 0 a 10.000 es: %d", suma)
}
