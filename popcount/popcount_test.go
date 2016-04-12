package popcount

import (
    "testing"
    "math/rand"
)

// BenchmarkPopCount PopCoount
func BenchmarkPopCount(b *testing.B) {   
    for i := 0; i < b.N; i++ {
        PopCount(uint64(rand.Uint32()))
    }
}

// BenchmarkPopCount PopCoount2
func BenchmarkPopCount2(b *testing.B) {   
    for i := 0; i < b.N; i++ {
        PopCount2(uint64(rand.Uint32()))
    }
}