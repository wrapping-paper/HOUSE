package DataRooms

import (
	"GolandProjects/HOME/PrintFILE"
	"GolandProjects/HOME/Rooms"
)

func StorageDataRooms() {

	var ListOfData []interface{}

	LivingRoom := Rooms.Space{
		Name:     "LivingRoom",
		Landlord: "Family",
		Features: Rooms.Features{
			Side:   "East",
			Square: 1,
		},
	}
	ListOfData = append(ListOfData, LivingRoom)
	Kitchen := Rooms.Space{
		Name:     "Kitchen",
		Landlord: "Family",
		Features: Rooms.Features{
			Side:   "Middle",
			Square: 1,
		},
	}
	ListOfData = append(ListOfData, Kitchen)
	Wardrobe := Rooms.Space{
		Name:     "Wardrobe",
		Landlord: "Family",
		Features: Rooms.Features{
			Side:   "North-West",
			Square: 1,
		},
	}
	ListOfData = append(ListOfData, Wardrobe)
	Bathroom := Rooms.Space{
		Name:     "Bathroom",
		Landlord: "Family",
		Features: Rooms.Features{
			Side:   "South-West",
			Square: 1,
		},
	}
	ListOfData = append(ListOfData, Bathroom)
	SisterRoom := Rooms.Space{
		Name:     "Sister's Room",
		Landlord: "Sister",
		Features: Rooms.Features{
			Side:   "South",
			Square: 1,
		},
	}
	ListOfData = append(ListOfData, SisterRoom)
	MyRoom := Rooms.Space{
		Name:     "My Room",
		Landlord: "Me",
		Features: Rooms.Features{
			Side:   "South-East",
			Square: 1,
		},
	}
	ListOfData = append(ListOfData, MyRoom)
	ParentsRoom := Rooms.Space{
		Name:     "Parents' Room",
		Landlord: "Parents",
		Features: Rooms.Features{
			Side:   "North",
			Square: 1,
		},
	}
	ListOfData = append(ListOfData, ParentsRoom)
	PrintFILE.PrintCONSOLE(ListOfData)
}
