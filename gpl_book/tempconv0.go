// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

// Celsius lol
type Celsius float64

// Fahrenheit dota
type Fahrenheit float64

// lol
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF convert
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC convert
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
