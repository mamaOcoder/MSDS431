package main

import (
	"testing"
)

// benchmarking for interface functions
func BenchmarkgetKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vegetableSet := map[string]interface{}{
			"potato":  true,
			"cabbage": true,
			"carrot":  true,
		}

		fruitRank := map[int]interface{}{
			1: "strawberry",
			2: "raspberry",
			3: "blueberry",
		}

		getStringKeys(vegetableSet)
		getIntKeys(fruitRank)
	}
}
