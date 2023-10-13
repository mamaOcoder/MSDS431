package main

import (
	"fmt"
)

func getStringKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func getIntKeys(m map[int]interface{}) []int {
	keys := make([]int, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func main() {
	//startTime := time.Now()

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

	fmt.Printf("vegetableSet keys: %+v\n", getStringKeys(vegetableSet))
	fmt.Printf("fruitRank keys: %+v\n", getIntKeys(fruitRank))

	//endTime := time.Now()

	// Calculate the execution time
	//executionTime := endTime.Sub(startTime)

	//fmt.Println("Execution time:", executionTime)
}
