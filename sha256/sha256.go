package main

import (
    "fmt"
    "crypto/sha256"
)

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    
    var d int    
    for i := 0; i < sha256.Size; i++ {
        e := c1[i]^c2[i]
        for e > 0 {
            d++
            e = e&(e-1)
        }
    }
    
    fmt.Printf("different bits count : %d\n", d)
}