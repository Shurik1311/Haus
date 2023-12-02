package devices

import (
	"fmt"
	"haus/structures/haus_parts/types"
)

type Devices struct {
	Type         string
	Length       types.Centimeter
	Width        types.Centimeter
	Color        string
	VoiceControl bool
}

func (d Devices) Print() {
	fmt.Print("Тип устройства: ", d.Type, "\nДлина устройства: ", d.Length, "\nШирина устройства: ", d.Width, "\nЦвет: ", d.Color, "\n")
}
