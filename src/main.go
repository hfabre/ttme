package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "TTME")
	rl.SetTargetFPS(60)

	nms := makeNewMapScreen()
	mcs := makeMapCreationScreen(&nms.mc)

	smInstance().addScreen("new_map", &nms)
	smInstance().addScreen("map_creation", &mcs)

	smInstance().switchScreen("new_map")

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		smInstance().tick()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
