package Family

import (
	"fmt"
)

type Subject struct {
	FirstName, Role string
	Features
}
type Features struct {
	Age, BloodType int
	Hobby          string
}

func PrintCONSOLE() {
	fmt.Println("\n\tFAMILY MEMBERS")
}
