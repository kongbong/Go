package tempconv

import "fmt"

// Celsius is celsius temperature
type Celsius float64
// Fahrenheit is fahrenheit temperature
type Fahrenheit float64
// Kelvin is Kelvin scale temperature
type Kelvin float64


const (
    // AbsoluteZeroC ...
    AbsoluteZeroC   Celsius = -273.15
    // FreezingC ...
    FreezingC       Celsius = 0
    // BoilingC ...
    BoilingC        Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g℃", c)}
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f)}
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k)}
