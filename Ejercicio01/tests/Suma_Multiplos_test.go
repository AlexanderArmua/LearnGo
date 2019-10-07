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

func BenchmarkGetSumaDivisores(b *testing.B) {
	want := 23333341666668

	for i := 0; i < b.N; i++ {
		b.Run("0 a 10.000.000", func(b *testing.B) {
			got := lib.GetSumaDivisores(0, 10000000)
			if got != want {
				b.Errorf("got %d want %d", got, want)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					lib.GetSumaDivisores(0, 10000000)
				}
			},
		)
	}
}

func ExampleGetSumaDivisores() {
	suma := lib.GetSumaDivisores(0, 10000)

	fmt.Printf("La suma de 0 a 10.000 es: %d", suma)
}
