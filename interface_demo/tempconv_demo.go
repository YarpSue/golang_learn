package interface_demo

import (
	"fmt"
	"testmode/tempconv"
)

type celsiusFlag struct {
	tempconv.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}