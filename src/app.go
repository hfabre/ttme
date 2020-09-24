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

// TODO: Handle mouse press when clickable is also scrollable and scrolled
func (a *app) Start() {
	tileset := *NewTileset(16, 16, "./assets/tilesetpkm.png")
	tilesetWidget := NewTilesetWidget(30, 425, 370, 370, &tileset)
	tilemap := NewTilemap(50, 50, &tileset)
	tilemapWidget := NewTilemapWidget(420, 95, 800, 700, tilemap)
	tilsetConfigurationWidget := NewTilesetConfigurationWidget(30, 750, &tileset)
	tilePropertiesWidget := NewTilePropertiesWidget(25, 200)
	tilemapConfigurationWidget := NewTilemapConfigurationWidget(80, 100, tilemap)

	for !r.WindowShouldClose() {

		// Handle Mouse inputs

		a.mousePosition = r.GetMousePosition()

		if r.IsMouseButtonPressed(r.MouseLeftButton) {
			a.mousePressed = true

			if !tilePropertiesWidget.editMode && tilesetWidget.Contains(a.mousePosition.X, a.mousePosition.Y) {
				tilePropertiesWidget.Unset()
				tilesetWidget.SelectTile(a.mousePosition.X, a.mousePosition.Y)
			}
		}

		if r.IsMouseButtonReleased(r.MouseLeftButton) {
			a.mousePressed = false
		}

		// Handle tile pasting

		if a.mousePressed && tilemapWidget.Contains(a.mousePosition.X, a.mousePosition.Y) {
			if tilePropertiesWidget.Selected() {
				tilemapWidget.SetTileFromPos(a.mousePosition.X, a.mousePosition.Y, tilesetWidget.selectedTile)
			} else {
				tilemapWidget.SetPropertyFromPos(a.mousePosition.X, a.mousePosition.Y, tilePropertiesWidget.SelectedProperty())
			}
		}

		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		tilemapWidget.Draw()
		tilesetWidget.Draw()
		tilsetConfigurationWidget.Draw(tilemapWidget)
		tilemapConfigurationWidget.Draw(tilemapWidget)

		a.ShowInfo()
		mouseTilePosInfo := fmt.Sprintf("Mouse tile position: %v - %v", tilemapWidget.GetTileXFromPos(a.mousePosition.X), tilemapWidget.GetTileYFromPos(a.mousePosition.Y))
		tilesetinfo := fmt.Sprintf("Tileset: %v - %v / %v", tileset.tileWidth, tileset.tileHeight, tileset.imagePath)
		tilemapinfo := fmt.Sprintf("Tilemap: %v - %v", tilemap.width, tilemap.height)

		r.DrawText(mouseTilePosInfo, 10, 50, 10, r.Black)
		r.DrawText(tilesetinfo, 10, 70, 10, r.Black)
		r.DrawText(tilemapinfo, 10, 90, 10, r.Black)
		tilePropertiesWidget.Draw()
		r.EndDrawing()
	}
	r.CloseWindow()
}