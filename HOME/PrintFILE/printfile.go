package PrintFILE

import (
	"GolandProjects/HOME/Family"
	"GolandProjects/HOME/Furniture"
	"GolandProjects/HOME/Rooms"
	"fmt"
)

func PrintCONSOLE(i []interface{}) {
	for j := 0; j < len(i); j++ {
		switch k := i[j].(type) {
		case Family.Subject:
			fmt.Println("\nFirst Name:", k.FirstName)
			fmt.Println("Role:", k.Role)
			fmt.Println("Age:", k.Age)
			fmt.Println("Blood Type:", k.BloodType)
			fmt.Println("Hobby:", k.Hobby)
		case Rooms.Space:
			fmt.Println("Name:", k.Name)
			fmt.Println("Landlord:", k.Landlord)
			fmt.Println("Side:", k.Side)
			fmt.Print("Square: ", k.Square, "\n\n")
		case Furniture.Stuff:
			fmt.Println("\nName:", k.Name)
			fmt.Println("Material:", k.Material)
			fmt.Println("Colour:", k.Colour)
			fmt.Println("Length:", k.Length)
			fmt.Println("Height:", k.Height)
			fmt.Println("Width:", k.Width)
		}
	}
}
