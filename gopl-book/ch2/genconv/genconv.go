// Package genconv performs general conversions.
package genconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
func (f Foot) String() string       { return fmt.Sprintf("%gfeet", f) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%gkg", k) }
