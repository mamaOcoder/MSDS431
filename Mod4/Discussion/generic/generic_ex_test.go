package main

import (
	"testing"
)

// benchmarking for generic function
func BenchmarkgetKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vegetableSet := map[string]bool{
			"potato":  true,
			"cabbage": true,
			"carrot":  true,
		}

		fruitRank := map[int]string{
			1: "strawberry",
			2: "raspberry",
			3: "blueberry",
		}

		getKeys(vegetableSet)
		getKeys(fruitRank)
	}
}
