package tempconv

import "fmt"

// Celsius is a temperature scale
type Celsius float64

// Fahrenheit is a temperature scale
type Fahrenheit float64

// Kelvin is a temperature scale
type Kelvin float64

// Const Temperature Var
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// This method means type Celsius implements the interface I,
// but we don't need to explicitly declare that it does so.
func (c Celsius) String() string {
	return fmt.Sprintf("%g˚C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g˚F", f)
}

func (f Kelvin) String() string {
	return fmt.Sprintf("%gK", f)
}

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}
