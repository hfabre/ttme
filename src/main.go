package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

// Tile size
const TileHeight = 16
const TileWidth = 16


func main() {
	rl.InitWindow(800, 600, "TTME")
	rl.SetTargetFPS(60)

	currentScreen := makeMapCreationScreen()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		currentScreen.tick()

		rl.EndDrawing()
	}

	currentScreen.unload()
	rl.CloseWindow()
}
