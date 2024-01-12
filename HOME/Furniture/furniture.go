package Furniture

import "fmt"

type Stuff struct {
	Name string
	Features
}
type Features struct {
	Material, Colour      string
	Length, Height, Width float32
}

func PrintCONSOLE() {
	fmt.Println("\tFURNITURE")
}
