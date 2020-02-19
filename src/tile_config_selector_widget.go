package main

import rl "github.com/gen2brain/raylib-go/raylib"

type tileConfig struct {
	name string
	mc mapConfiguration
	color rl.Color
}

func (tc tileConfig) draw() {

}

//type tileConfigSelectorWidget struct {
//	mc *mapConfiguration
//	offsetX float32
//	offsetY float32
//	selectedTile tileConfig
//}
//
//func (tsc *tileConfigSelectorWidget) selectTile(x, y float32) {
//	tileX := (x - tsc.offsetX) / float32(tsc.mc.tileWidth)
//	tileY := (y - tsc.offsetY) / float32(tsc.mc.tileHeight)
//	tilePos := int32(tileY) * tsc.tileset.tilesByLine() + int32(tileX)
//
//	tsc.selectedTile = tile{int(tilePos)}
//	rl.whi
//}
//
//func (tsc tileConfigSelectorWidget) contains(x, y float32) bool {
//	point := rl.Vector2{x, y}
//	rect := rl.Rectangle{tsc.offsetX, tsc.offsetY, float32(tsc.tileset.texture.Width), float32(tsc.tileset.texture.Height)}
//
//	return rl.CheckCollisionPointRec(point, rect)
//}
//
//func (tsc tileConfigSelectorWidget) draw() {
//	rl.DrawTexture(tsc.tileset.texture, int32(tsc.offsetX), int32(tsc.offsetY), rl.White)
//
//	tileX, tileY := tsc.selectedTile.getTilsetPosition(tsc.tileset)
//	x := tsc.offsetX + tileX
//	y := tsc.offsetY + tileY
//
//	rl.DrawRectangleLines(int32(x), int32(y), int32(tsc.mc.tileWidth), int32(tsc.mc.tileHeight), rl.Red)
//}

