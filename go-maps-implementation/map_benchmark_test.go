package main

import (
	"strconv"
	"testing"
)

// setupMap creates a map with a specified number of elements.
func setupMap(size int) map[string]int {
	m := make(map[string]int, size)
	for i := 0; i < size; i++ {
		key := "key_" + strconv.Itoa(i)
		m[key] = i
	}
	return m
}

// BenchmarkMapLookup benchmarks the performance of map lookups.
func BenchmarkMapLookup(b *testing.B) {
	b.ReportAllocs() // Explicitly enable memory allocation reporting for this benchmark.
	const mapSize = 100000
	m := setupMap(mapSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key_" + strconv.Itoa(i%mapSize)
		_ = m[key]
	}
}

// BenchmarkMapInsert benchmarks the performance of map insertions.
func BenchmarkMapInsert(b *testing.B) {
	b.ReportAllocs() // Explicitly enable memory allocation reporting.
	m := make(map[string]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "new_key_" + strconv.Itoa(i)
		m[key] = i
	}
}

// BenchmarkMapDelete benchmarks the performance of map deletions.
func BenchmarkMapDelete(b *testing.B) {
	b.ReportAllocs() // Explicitly enable memory allocation reporting.
	const mapSize = 100000
	m := setupMap(mapSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key_" + strconv.Itoa(i%mapSize)
		delete(m, key)
	}
}
