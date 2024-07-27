package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(size) // Valores aleatórios até o tamanho do slice
	}
	return slice
}

func BenchmarkSort5Native(b *testing.B) {
	data := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortNative(data)
	}
}

func BenchmarkSort5Immutable(b *testing.B) {
	data := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortImmutable(data)
	}
}

func BenchmarkSort1000Native(b *testing.B) {
	data := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortNative(data)
	}
}

func BenchmarkSort1000Immutable(b *testing.B) {
	data := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortImmutable(data)
	}
}

func BenchmarkSort10KNative(b *testing.B) {
	data := generateSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortNative(data)
	}
}

func BenchmarkSort10KImmutable(b *testing.B) {
	data := generateSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortImmutable(data)
	}
}

func BenchmarkSort1MNative(b *testing.B) {
	data := generateSlice(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortNative(data)
	}
}

func BenchmarkSort1MImmutable(b *testing.B) {
	data := generateSlice(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortImmutable(data)
	}
}
