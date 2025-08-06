package main

import (
	"math/rand"
	"strconv"
	"testing"
)

const numRows = 10_000_000

func generateTestMap(n int) map[string]int {
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		m[strconv.Itoa(i)] = rand.Int()
	}
	return m
}

func BenchmarkMapInsert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m := make(map[string]int)
		for j := 0; j < numRows; j++ {
			m[strconv.Itoa(j)] = j
		}
	}
}

func BenchmarkMapLookup(b *testing.B) {
	m := generateTestMap(numRows)
	keys := make([]string, numRows)
	for i := 0; i < numRows; i++ {
		keys[i] = strconv.Itoa(i)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		key := keys[rand.Intn(numRows)]
		_ = m[key]
	}
}

func BenchmarkMapDelete(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m := generateTestMap(numRows)
		b.StartTimer()
		for j := 0; j < numRows; j++ {
			delete(m, strconv.Itoa(j))
		}
		b.StopTimer()
	}
}
