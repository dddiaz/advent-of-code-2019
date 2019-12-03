package main

//func main() {
//	fmt.Println("Fuel Calculator")
//}

func ModuleFuel(moduleMass int64) int64 {
	// since this is an int type, the fractional part is left off. so we are already rounding down
	moduleMass = moduleMass / 3
	moduleMass = moduleMass - 2
	return moduleMass
}
