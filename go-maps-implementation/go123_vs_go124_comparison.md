# Go 1.23 vs Go 1.24 Map Benchmark Comparison

**Test Environment:**
- CPU: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
- OS: macOS (darwin/amd64)
- Dataset: 10,000,000 rows
- Test Date: August 6, 2025

## Performance Comparison Summary

### Go 1.23 Results (from previous runs)
Based on earlier benchmark runs in this session:

| Operation | Time | Memory | Allocations |
|-----------|------|--------|-------------|
| **Insert** | ~5.63s | ~1.07 GB | ~10.3M |
| **Lookup** | ~528ns | 0 B/op | 0 |
| **Delete** | ~5.21s | ~623 MB | ~20M |

### Go 1.24 Results (current)

| Operation | Time | Memory | Allocations |
|-----------|------|--------|-------------|
| **Insert** | ~4.29s | ~929 MB | ~10.1M |
| **Lookup** | ~385ns | 0 B/op | 0 |
| **Delete** | ~4.51s | ~579 MB | ~20M |

## Performance Improvements in Go 1.24

### 🚀 Insert Operations
- **Speed Improvement**: 23.8% faster (5.63s → 4.29s)
- **Memory Reduction**: 13.2% less memory (1.07GB → 929MB)
- **Allocation Efficiency**: Slightly fewer allocations (~10.3M → ~10.1M)

### ⚡ Lookup Operations  
- **Speed Improvement**: 27.1% faster (528ns → 385ns)
- **Memory Usage**: Unchanged (0 allocations)
- **Consistency**: Maintained O(1) performance

### 🗑️ Delete Operations
- **Speed Improvement**: 13.4% faster (5.21s → 4.51s)
- **Memory Reduction**: 7.1% less memory (623MB → 579MB)
- **Allocation Count**: Similar allocation patterns

## Key Performance Gains

### Overall Improvements:
1. **Faster Execution**: 13-27% speed improvements across all operations
2. **Better Memory Efficiency**: 7-13% reduction in memory usage
3. **Improved Hash Table**: Better internal algorithms for map operations
4. **Optimized Allocator**: More efficient memory allocation patterns

### Throughput Comparison:
| Operation | Go 1.23 (ops/sec) | Go 1.24 (ops/sec) | Improvement |
|-----------|-------------------|-------------------|-------------|
| Insert    | ~1.78M            | ~2.33M            | +31% |
| Lookup    | 1.89M/µs          | 2.60M/µs          | +37% |
| Delete    | ~1.92M            | ~2.22M            | +16% |

## Technical Analysis

### What Changed in Go 1.24:
- **Runtime Optimizations**: Improved garbage collector and memory allocator
- **Map Implementation**: Enhanced hash table algorithms
- **Compiler Improvements**: Better code generation and optimization
- **Memory Layout**: More efficient data structure layouts

### Real-World Impact:
- **Applications with heavy map usage** will see immediate performance benefits
- **Memory-constrained environments** benefit from reduced allocation overhead
- **High-frequency lookup operations** experience significant speed improvements

## Conclusion

Go 1.24 shows substantial improvements in map operations:
- **Best Improvement**: Lookups (27% faster)
- **Most Consistent**: All operations improved
- **Memory Efficiency**: Reduced memory footprint across all operations

These improvements make Go 1.24 particularly beneficial for applications that rely heavily on map operations, such as caches, indexes, and data processing pipelines.

---
*Benchmark comparison based on 3-run averages for statistical accuracy*
