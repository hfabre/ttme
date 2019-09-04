package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type newMapScreen struct {
	tileWidthInput, tileHeightInput string
	widthInput, heightInput string
	mc mapConfiguration
}

func makeNewMapScreen() newMapScreen {
	return newMapScreen{"16", "16", "20", "20", mapConfiguration{}}
}

func (nms newMapScreen) load() {
	// Implement Screen interface
}

func (nms newMapScreen) unload() {
	// Implement Screen interface
}

func (nms *newMapScreen) tick() {

	// Width Text box
	raygui.Label(rl.Rectangle{10, 10, 30, 10}, "Map Width (tile): ")
	nms.widthInput = raygui.TextBox(rl.Rectangle{170, 5, 20, 15}, nms.widthInput)

	// Height Text box
	raygui.Label(rl.Rectangle{10, 40, 30, 10}, "Map Height (tile): ")
	nms.heightInput = raygui.TextBox(rl.Rectangle{170, 35, 20, 15}, nms.heightInput)

	// Width Text box
	raygui.Label(rl.Rectangle{10, 90, 30, 10}, "Tile Width (px): ")
	nms.tileWidthInput = raygui.TextBox(rl.Rectangle{170, 85, 20, 15}, nms.tileWidthInput)

	// Height Text box
	raygui.Label(rl.Rectangle{10, 110, 30, 10}, "Tile Height (px): ")
	nms.tileHeightInput = raygui.TextBox(rl.Rectangle{170, 105, 20, 15}, nms.tileHeightInput)

	if raygui.Button(rl.Rectangle{10, 170, 60, 15}, "Create") {
		nms.mc.width, _ = strconv.Atoi(nms.widthInput)
		nms.mc.height, _ = strconv.Atoi(nms.heightInput)
		nms.mc.tileWidth, _ = strconv.Atoi(nms.tileWidthInput)
		nms.mc.tileHeight, _ = strconv.Atoi(nms.tileHeightInput)

		smInstance().switchScreen("map_creation")
	}
}
