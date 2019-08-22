package main

import (
	"encoding/json"
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"io/ioutil"
)

type exportWidget struct {
	offsetX, offsetY int
}

type jsonExport struct {
	Width, Height int
	Board [][]int
}

func (ew exportWidget) draw(mcw mapConfigurationWidget, board board) {
	if raygui.Button(rl.Rectangle{float32(ew.offsetX), float32(ew.offsetY), 20, 15}, "Export as json") {
		ew.exportJSON("./assets/map.json", mcw, board)
	}
}

func (ew exportWidget) exportJSON(path string, mcw mapConfigurationWidget, board board) {
	export := jsonExport{Width: mcw.width, Height: mcw.height}

	export.Board = make([][]int, export.Height)
	for y := 0; y < export.Height; y++ {
		export.Board[y] = make([]int, export.Width)
		for x := 0; x < export.Width; x++ {
			export.Board[y][x] = board.tiles[y][x].index
		}
	}

	file, _ := json.Marshal(export)
	_ = ioutil.WriteFile(path, file, 0644)
}
