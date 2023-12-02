package house

import (
	"fmt"
	"haus/structures/haus_parts/devices"
	"haus/structures/haus_parts/furniture"
	"haus/structures/haus_parts/types"
)

type House struct {
	Devices       []devices.Devices
	Furniture     []furniture.Furniture
	Rooms         int
	Square        types.SquareMeter
	CeilingHeight types.Centimeter
}

func (h House) Print() {
	fmt.Print("Комнаты: ", h.Rooms, "\nПлощадь: ", h.Square, "\nВысота потолков: ", h.CeilingHeight, "\n")

	for _, tmpDevices := range h.Devices {
		tmpDevices.Print()
	}
	for _, tmpFurniture := range h.Furniture {
		tmpFurniture.Print()
	}
}
func Make() House {
	table := furniture.Furniture{
		Type:          "Стол",
		Length:        120,
		Width:         60,
		Color:         "White",
		ComfortRating: 3,
	}
	chair := furniture.Furniture{
		Type:          "Стул",
		Length:        40,
		Width:         40,
		Color:         "Зеленый",
		ComfortRating: 1,
	}
	bed := furniture.Furniture{
		Type:          "Кровать",
		Length:        200,
		Width:         120,
		Color:         "Белый",
		ComfortRating: 4,
	}
	wardrobe := furniture.Furniture{
		Type:          "Шкаф",
		Length:        80,
		Width:         50,
		Color:         "Белый",
		ComfortRating: 5,
	}

	Peka := devices.Devices{
		Type:         "ПК",
		Length:       35,
		Width:        20,
		Color:        "Черный",
		VoiceControl: false,
	}

	phone := devices.Devices{
		Type:         "Телефон",
		Length:       15,
		Width:        7,
		Color:        "Черный",
		VoiceControl: true,
	}

	watch := devices.Devices{
		Type:         "Часы",
		Length:       4,
		Width:        1,
		Color:        "Черный",
		VoiceControl: false,
	}

	house := House{
		Devices:       []devices.Devices{Peka, phone, watch},
		Furniture:     []furniture.Furniture{table, chair, wardrobe, bed},
		Rooms:         1,
		Square:        20,
		CeilingHeight: 250,
	}
	return house
}
