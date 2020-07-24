package ttme
import r "github.com/lachee/raylib-goplus/raylib"

func Start() {
	r.InitWindow(1280, 900, "Raylib Go Plus")
	tileset := *NewTileset(16, 16, "toto")
	tilemap := tilemap{tileset: tileset, width: 50, height: 50}
	tilemapWidget := NewTilemapWidget(10, 10, tilemap)

	for !r.WindowShouldClose() {
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		tilemapWidget.Draw()
		r.EndDrawing()
	}
	r.CloseWindow()
}