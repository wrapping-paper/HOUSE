package Rooms

import "fmt"

type Space struct {
	Name, Landlord string
	Features
}
type Features struct {
	Side   string
	Square float32
}

func PrintCONSOLE() {
	fmt.Print("\n\tROOMS\n\n")
}
