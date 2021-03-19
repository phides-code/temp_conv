// Package uniconv performs all kinds of conversions.
package uniconv

import (
    "fmt"
    "os"
    "strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin(((f - 32) * 5 / 9) +273.15) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func main() {
    
    inputTemp, _ := strconv.ParseFloat(os.Args[1], 64)
    
    fmt.Println(Celsius(inputTemp).String(), "is", CToF(Celsius(inputTemp)).String(), "and", CToK(Celsius(inputTemp)).String() )
    
    
    fmt.Println(Fahrenheit(inputTemp).String(), "is", FToC(Fahrenheit(inputTemp)).String(), "and", FToK(Fahrenheit(inputTemp)).String() )
    
    
    fmt.Println(Kelvin(inputTemp).String(), "is", KToF(Kelvin(inputTemp)).String(), "and", KToC(Kelvin(inputTemp)).String() )
    
}