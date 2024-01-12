package HOME

import (
	"GolandProjects/HOME/Family"
	"GolandProjects/HOME/Family/DataFamily"
	"GolandProjects/HOME/Furniture"
	"GolandProjects/HOME/Furniture/DataFurniture"
	"GolandProjects/HOME/Rooms"
	"GolandProjects/HOME/Rooms/DataRooms"
)

func Make() {
	Family.PrintCONSOLE()
	DataFamily.StorageDataFamily()
	Rooms.PrintCONSOLE()
	DataRooms.StorageDataRooms()
	Furniture.PrintCONSOLE()
	DataFurniture.StorageDataFurniture()
}
