package DataFurniture

import (
	"GolandProjects/HOME/Furniture"
	"GolandProjects/HOME/PrintFILE"
)

func StorageDataFurniture() {

	var ListOfData []interface{}

	Sofa := Furniture.Stuff{
		Name: "Sofa",
		Features: Furniture.Features{
			Material: "Leather",
			Colour:   "Brown",
			Length:   1,
			Height:   1,
			Width:    1,
		},
	}
	ListOfData = append(ListOfData, Sofa)
	Table := Furniture.Stuff{
		Name: "Table",
		Features: Furniture.Features{
			Material: "Spruce Wood",
			Colour:   "Rust",
			Length:   1,
			Height:   1,
			Width:    1,
		},
	}
	ListOfData = append(ListOfData, Table)
	Armchair := Furniture.Stuff{
		Name: "Armchair",
		Features: Furniture.Features{
			Material: "Cotton",
			Colour:   "Cream",
			Length:   1,
			Height:   1,
			Width:    1,
		},
	}
	ListOfData = append(ListOfData, Armchair)
	Bed := Furniture.Stuff{
		Name: "Bed",
		Features: Furniture.Features{
			Material: "Wood Slab",
			Colour:   "Dark Brown",
			Length:   1,
			Height:   1,
			Width:    1,
		},
	}
	ListOfData = append(ListOfData, Bed)
	Chest := Furniture.Stuff{
		Name: "Chest",
		Features: Furniture.Features{
			Material: "Pine Wood",
			Colour:   "Cream",
			Length:   1,
			Height:   1,
			Width:    1,
		},
	}
	ListOfData = append(ListOfData, Chest)
	FloorLamp := Furniture.Stuff{
		Name: "Floor Lamp",
		Features: Furniture.Features{
			Material: "Aluminium",
			Colour:   "Black",
			Length:   1,
			Height:   1,
			Width:    1,
		},
	}
	ListOfData = append(ListOfData, FloorLamp)
	PrintFILE.PrintCONSOLE(ListOfData)
}
