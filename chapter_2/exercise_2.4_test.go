package main

import "testing"

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
  var count int = 0
  for i := 0; i < 64; i++ {
    if byte(x&1) == 1 {
      count += 1
    }
    x = x>>1
  }
	return count
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
    PopCount(123)
	}
}
