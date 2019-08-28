package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type tileSelectorWidget struct {
	tileset tileset
	offsetX float32
	offsetY float32
	selectedTile tile
}

func (ts *tileSelectorWidget) selectTile(x, y float32) {
	tileX := (x - ts.offsetX) / TileWidth
	tileY := (y - ts.offsetY) / TileHeight
	tilePos := int32(tileY) * ts.tileset.tilesByLine() + int32(tileX)

	ts.selectedTile = tile{int(tilePos)}
}

func (ts tileSelectorWidget) contains(x, y float32) bool {
	point := rl.Vector2{x, y}
	rect := rl.Rectangle{ts.offsetX, ts.offsetY, float32(ts.tileset.texture.Width), float32(ts.tileset.texture.Height)}

	return rl.CheckCollisionPointRec(point, rect)
}

func (ts tileSelectorWidget) draw() {
	rl.DrawTexture(ts.tileset.texture, int32(ts.offsetX), int32(ts.offsetY), rl.White)

	tileX, tileY := ts.selectedTile.getTilsetPosition(ts.tileset)
	x := ts.offsetX + tileX
	y := ts.offsetY + tileY

	rl.DrawRectangleLines(int32(x), int32(y), TileWidth, TileHeight, rl.Red)
}
