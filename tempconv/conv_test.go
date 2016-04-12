package tempconv

import ( 
    "testing"
    "fmt"
)

func TestCToF(t *testing.T) {
    fmt.Println(CToF(AbsoluteZeroC))
}

func TestFToC(t *testing.T) {
    fmt.Println(FToC(CToF(BoilingC)))
}

func TestKToC(t *testing.T) {
    c := KToC(0)
    if c != -273.15 {
        t.Error("0 K is not -273.15 C")
    }
    
    c = KToC(1000)
    if c != 1 {
        t.Error("1000 K is not 1 C")
    }
}