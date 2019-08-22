package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

// Tile size
const TileHeight = 16
const TileWidth = 16

// Gui layout
const BoardOffsetX = 300
const BoardOffsetY = 10


func main() {
	rl.InitWindow(800, 450, "Mapper")
	rl.SetTargetFPS(60)

	tileset := makeTileset("./assets/tilesetpkm.png")
	mcw := makeMapConfigurationWidget(10, 10)
	b := emptyBoard(mcw.width, mcw.height, BoardOffsetX, BoardOffsetY)
	ts := tileSelector{tileset, 10, 50, tile{0}}
	ew := exportWidget{200, 80}
	mousePressed := false

	for !rl.WindowShouldClose() {

		// Updating

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
