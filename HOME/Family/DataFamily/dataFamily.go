package DataFamily

import (
	"GolandProjects/HOME/Family"
	"GolandProjects/HOME/PrintFILE"
)

func StorageDataFamily() {

	var ListOfData []interface{}

	Father := Family.Subject{
		FirstName: "Anton",
		Role:      "Father",
		Features: Family.Features{
			Age:       45,
			BloodType: 2,
			Hobby:     "Car",
		},
	}
	ListOfData = append(ListOfData, Father)
	Mother := Family.Subject{
		FirstName: "Julia",
		Role:      "Mother",
		Features: Family.Features{
			Age:       45,
			BloodType: 2,
			Hobby:     "Cooking",
		},
	}
	ListOfData = append(ListOfData, Mother)
	Sister := Family.Subject{
		FirstName: "Anastasia",
		Role:      "Sister",
		Features: Family.Features{
			Age:       24,
			BloodType: 2,
			Hobby:     "Art",
		},
	}
	ListOfData = append(ListOfData, Sister)
	Brother := Family.Subject{
		FirstName: "Lev",
		Role:      "Brother",
		Features: Family.Features{
			Age:       4,
			BloodType: 2,
			Hobby:     "Cartoons",
		},
	}
	ListOfData = append(ListOfData, Brother)
	PrintFILE.PrintCONSOLE(ListOfData)
}
