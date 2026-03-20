package main

import (
	"fmt"

	mergesorted "merge-sorted"
)

func main() {
	collection1 := []int{9, 5, 1}  // descending
	collection2 := []int{2, 6, 8}  // ascending
	collection3 := []int{3, 4, 7}  // ascending

	result := mergesorted.Merge(collection1, collection2, collection3)
	fmt.Println("Input collection1 (desc):", collection1)
	fmt.Println("Input collection2 (asc): ", collection2)
	fmt.Println("Input collection3 (asc): ", collection3)
	fmt.Println("Merged result (asc):     ", result)
}
