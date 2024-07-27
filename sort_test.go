package main

import (
	"testing"
)

func TestSortNative(t *testing.T) {
	data := []int{3, 1, 4, 1, 5, 9}
	SortNative(data)
	expected := []int{1, 1, 3, 4, 5, 9}
	for i, v := range expected {
		if data[i] != v {
			t.Errorf("sortNative() = %v; want %v", data, expected)
		}
	}
}

func TestSortImmutable(t *testing.T) {
	data := []int{3, 1, 4, 1, 5, 9}
	sortedData := SortImmutable(data)
	expected := []int{1, 1, 3, 4, 5, 9}
	for i, v := range expected {
		if sortedData[i] != v {
			t.Errorf("sortImmutable() = %v; want %v", sortedData, expected)
		}
	}
	// Verifica se a coleção original não foi alterada
	for i, v := range data {
		if v != []int{3, 1, 4, 1, 5, 9}[i] {
			t.Errorf("sortImmutable() modified original slice; got %v, want %v", data, []int{3, 1, 4, 1, 5, 9})
			break
		}
	}
}
