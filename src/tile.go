package main

import rl "github.com/gen2brain/raylib-go/raylib"

type tile struct {
	index int
}

func (t tile) index32() int32 {
	return int32(t.index)
}

func (t tile) getTilsetPosition(tileset tileset) (float32, float32) {
	tilesetWidth := tileset.tilesByLine()

	tileX := float32((t.index32() % tilesetWidth) * int32(tileset.mc.tileHeight))
	tileY := float32((t.index32() / tilesetWidth) * int32(tileset.mc.tileHeight))

	return tileX, tileY
}

func (t tile) draw(x, y int, tileset tileset) {
	x32 := float32(x)
	y32 := float32(y)

	tileX, tileY := t.getTilsetPosition(tileset)

	pos := rl.Vector2{x32, y32}
	subRec := rl.Rectangle{tileX, tileY, float32(tileset.mc.tileWidth), float32(tileset.mc.tileHeight)}

	rl.DrawTextureRec(tileset.texture, subRec, pos, rl.White)
}