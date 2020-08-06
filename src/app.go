package ttme

import (
	"fmt"
	r "github.com/lachee/raylib-goplus/raylib"
)

type app struct {
	width, height int
	name string
	mousePressed bool
	mousePosition r.Vector2
}

func NewApp(width, height int, name string) *app {
	r.InitWindow(width, height, name)

	return &app{
		width:         width,
		height:        height,
		name:          name,
		mousePressed:  false,
		mousePosition: r.NewVector2Zero(),
	}
}

func (a app) ShowInfo() {
	mousePosInfo := fmt.Sprintf("Mouse position: %f - %f", a.mousePosition.X, a.mousePosition.Y)
	mouseStateInfo := fmt.Sprintf("Mouse pressed: %v", a.mousePressed)

	r.DrawText(mousePosInfo, 10, 10, 10, r.Black)
	r.DrawText(mouseStateInfo, 10, 30, 10, r.Black)
}

func (a *app) Start() {
	tileset := *NewTileset(16, 16, "./assets/tilesetpkm.png")
	tilesetWidget := NewTilesetWidget(30, 325, 370, 470, tileset)
	tilemap := tilemap{tileset: tileset, width: 50, height: 50}
	tilemapWidget := NewTilemapWidget(420, 95, 800, 700, tilemap)

	for !r.WindowShouldClose() {

		// Handle Mouse inputs

		a.mousePosition = r.GetMousePosition()

		if r.IsMouseButtonPressed(r.MouseLeftButton) {
			a.mousePressed = true

			if tilesetWidget.Contains(a.mousePosition.X, a.mousePosition.Y) {
				println("Select tile")
				tilesetWidget.SelectTile(a.mousePosition.X, a.mousePosition.Y)
			}
		}

		if r.IsMouseButtonReleased(r.MouseLeftButton) {
			a.mousePressed = false
		}

		// Handle tile pasting
		//
		//if mcs.mousePressed && mcs.bw.contains(mousePosition.X, mousePosition.Y) {
		//	mcs.bw.setTileFromPos(mousePosition.X, mousePosition.Y, mcs.tsw.selectedTile)
		//}

		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		tilemapWidget.Draw()
		tilesetWidget.Draw()
		a.ShowInfo()
		r.EndDrawing()
	}
	r.CloseWindow()
}