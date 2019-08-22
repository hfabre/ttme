package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type mapConfigurationWidget struct {
	offsetX, offsetY int
	widthInput, heightInput string
	width, height int
	hasChanges bool
}

func makeMapConfigurationWidget(offsetX, offsetY int) mapConfigurationWidget {
	return mapConfigurationWidget{offsetX, offsetY, "20", "20", 20, 20, false}
}

func (mcw *mapConfigurationWidget) draw() {
	// Width Text box
	raygui.Label(rl.Rectangle{float32(mcw.offsetX), float32(mcw.offsetY), 30, 10}, "Map Width (unit is a tile): ")
	mcw.widthInput = raygui.TextBox(rl.Rectangle{float32(mcw.offsetX + 160), float32(mcw.offsetY - 5), 20, 15}, mcw.widthInput)

	// Height Text box
	raygui.Label(rl.Rectangle{float32(mcw.offsetX), float32(mcw.offsetY + 30), 30, 10}, "Map Height (unit is a tile): ")
	mcw.heightInput = raygui.TextBox(rl.Rectangle{float32(mcw.offsetX + 160), float32(mcw.offsetY + 25), 20, 15}, mcw.heightInput)

	if raygui.Button(rl.Rectangle{float32(mcw.offsetX + 200), float32(mcw.offsetY + 50), 20, 15}, "Update") {
		mcw.update()
	}
}

func (mcw *mapConfigurationWidget) update() {
	mcw.height, _ = strconv.Atoi(mcw.heightInput)
	mcw.width, _ = strconv.Atoi(mcw.widthInput)
	mcw.hasChanges = true
}
