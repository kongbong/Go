package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
    return int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] + 
        pc[byte(x>>(3*8))] + 
        pc[byte(x>>(4*8))] + 
        pc[byte(x>>(5*8))] + 
        pc[byte(x>>(6*8))] + 
        pc[byte(x>>(7*8))])
}

// PopCount2 returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
    var total int
    var i uint
    for i = 0; i < 8; i++ { 
        total += int(pc[byte(x>>(i*8))])
    }
    return total
}

// PopCount3 returns the population count (number of set bits) of x.
func PopCount3(x uint64) int {
    var total int
    for x > 0 {
        if x & 1 == 1 {
            total++
        }
        x = x >> 1        
    }
    return total
}

// PopCount4 returns the population count (number of set bits) of x.
func PopCount4(x uint64) int {
    var total int
    for x > 0 {
        total ++
        x = x&(x-1)
    }
    return total
}