package main

import (
	"sort"
)

func SortNative(slice []int) {
	sort.Ints(slice)
}

func SortImmutable(slice []int) []int {

	copySlice := make([]int, len(slice))
	copy(copySlice, slice)

	sort.Ints(copySlice)
	return copySlice
}
