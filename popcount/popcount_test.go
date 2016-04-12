package popcount

import (
    "testing"
    "math/rand"
)

// TestPopCount testPopCount func
func TestPopCount(t *testing.T) {
    r := uint64(rand.Uint32())
    
    r1 := PopCount(r)
    r2 := PopCount2(r)
    r3 := PopCount3(r)
    r4 := PopCount4(r)
    
    if r1 == r2 && r1 == r3 && r1 == r4 {
        t.Log("PopCount1,2,3 are equal")
    } else {
        t.Errorf("PopCount1,2,3 should be equal")
    }    
}

// BenchmarkPopCount PopCoount
func BenchmarkPopCount(b *testing.B) {   
    for i := 0; i < b.N; i++ {
        PopCount(uint64(rand.Uint32()))
    }
}

// BenchmarkPopCount2 PopCoount2
func BenchmarkPopCount2(b *testing.B) {   
    for i := 0; i < b.N; i++ {
        PopCount2(uint64(rand.Uint32()))
    }
}

// BenchmarkPopCount3 PopCoount3
func BenchmarkPopCount3(b *testing.B) {   
    for i := 0; i < b.N; i++ {
        PopCount3(uint64(rand.Uint32()))
    }
}

// BenchmarkPopCount4 PopCoount4
func BenchmarkPopCount4(b *testing.B) {   
    for i := 0; i < b.N; i++ {
        PopCount4(uint64(rand.Uint32()))
    }
}