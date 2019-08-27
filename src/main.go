package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/sqweek/dialog"
)

// Tile size
const TileHeight = 16
const TileWidth = 16

// Gui layout
const BoardOffsetX = 300
const BoardOffsetY = 10


func main() {
	rl.InitWindow(800, 600, "Mapper")
	rl.SetTargetFPS(60)

	tileset := makeTileset("./assets/tilesetpkm.png")
	mcw := makeMapConfigurationWidget(10, 10)
	b := emptyBoard(mcw.width, mcw.height, BoardOffsetX, BoardOffsetY)
	ts := tileSelector{tileset, 10, 70, tile{0}}
	ew := exportWidget{200, 80}
	mousePressed := false

	for !rl.WindowShouldClose() {

		// Updating

		// Find a tileset
		// TODO: Write a widget and make sure user can specify tiles attributes (height and width)

		if raygui.Button(rl.Rectangle{float32(ts.offsetX), float32(ts.offsetY) - 20, 80, 15}, "Select tileset") {
			tilesetPath, _ := dialog.File().Load()
			tileset = makeTileset(tilesetPath)
			ts = tileSelector{tileset, 10, 70, tile{0}}
		}

		// Check map size changes

		if mcw.hasChanges {
			newBoard := emptyBoard(mcw.width, mcw.height, BoardOffsetX, BoardOffsetY)
			b.copy(newBoard)
			b = newBoard
			mcw.hasChanges = false
		}

		// Handle Mouse inputs

		mousePosition := rl.GetMousePosition()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mousePressed = true

			if ts.contains(mousePosition.X, mousePosition.Y) {
				ts.selectTile(mousePosition.X, mousePosition.Y)
			}
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			mousePressed = false
		}

		// Handle tile pasting

		if mousePressed && b.contains(mousePosition.X, mousePosition.Y) {
			b.setTileFromPos(mousePosition.X, mousePosition.Y, ts.selectedTile)
		}



		// Drawing

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		mcw.draw()
		b.draw(tileset)
		ts.draw()
		ew.draw(mcw, *b)

		rl.EndDrawing()
	}

	tileset.unload()
	rl.CloseWindow()
}
