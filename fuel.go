package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fuel Calculator")
}

func ModuleFuel(moduleMass int) int {
	// since this is an int type, the fractional part is left off. so we are already rounding down
	moduleMass = moduleMass / 3
	moduleMass = moduleMass - 2
	return moduleMass
}
