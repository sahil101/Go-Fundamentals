# Go-Fundamentals

## Map Benchmark Tests

This repository includes benchmark tests for Go map operations (insert, lookup, delete) with 10 million rows.

### Running the Benchmarks

To run the map operation benchmarks:

```bash
cd go-maps-implementation
go test -bench=BenchmarkMap -benchmem -v
```

### Individual Benchmark Tests

Run specific benchmarks:

```bash
# Insert benchmark
go test -bench=BenchmarkMapInsert -benchmem

# Lookup benchmark  
go test -bench=BenchmarkMapLookup -benchmem

# Delete benchmark
go test -bench=BenchmarkMapDelete -benchmem
```