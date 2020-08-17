package main

import (
	"flag"
	"fmt"
	"golangbook/chap2/tempcov"
)

type celsiusFlag struct { tempcov.Celsius }

func (cflag *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		cflag.Celsius = tempcov.Celsius(value)
	case "F":
		cflag.Celsius = tempcov.FToC(tempcov.Fahrenheit(value))
	case "K":
		cflag.Celsius = tempcov.KToC(tempcov.Kelvin(value))
	default:
		return fmt.Errorf("invalid unit %q", unit)
	}
	return nil
}

func CelsiusFlag(name string, defaultVal tempcov.Celsius, usage string) *tempcov.Celsius {
	f := celsiusFlag{defaultVal}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", tempcov.Celsius(20.0), "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}