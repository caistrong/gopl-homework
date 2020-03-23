package lenconv

import "fmt"

// Inch 英寸
type Inch float64

// Meter 米
type Meter float64

// Foot 英尺
type Foot float64

// This method means type Celsius Inch the interface I(has method String),
// but we don't need to explicitly declare that it does so.
func (i Inch) String() string {
	return fmt.Sprintf("%g英寸", i)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g米", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g英尺", f)
}

// FToM convert Foot to Meter
func FToM(f Foot) Meter {
	return Meter(f * 0.3048)
}

// FToI convert Foot to Inch
func FToI(f Foot) Inch {
	return Inch(f * 12)
}
