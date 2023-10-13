package main

import (
	"fmt"
)

func getKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	//startTime := time.Now()

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

	fmt.Printf("vegetableSet keys: %+v\n", getKeys(vegetableSet))
	fmt.Printf("fruitRank keys: %+v\n", getKeys(fruitRank))

	//endTime := time.Now()

	// Calculate the execution time
	//executionTime := endTime.Sub(startTime)

	//fmt.Println("Execution time:", executionTime)
}
