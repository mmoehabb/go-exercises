package main

import "testing"

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
  var count int = 0
  for x != 0 {
    x = x&(x-1)
    count++
  }
	return count
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
    PopCount(123)
	}
}
