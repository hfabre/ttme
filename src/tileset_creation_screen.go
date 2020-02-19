package main
//
//import (
//	"fmt"
//	"github.com/gen2brain/raylib-go/raygui"
//	rl "github.com/gen2brain/raylib-go/raylib"
//	"image"
//	"image/draw"
//	"os"
//)
//
//const defaultColumn = 10
//const defaultLine = 30
//
//type tilesetCreationScreen struct {
//	mc mapConfiguration
//	tiles []image.Image
//}
//
//func makeTilesetCreationScreen(mc mapConfiguration) tilesetCreationScreen {
//	tcs := tilesetCreationScreen{mc: mc}
//
//	// Start tileset with a transparent tile
//	tcs.addTile(tcs.emptyImg())
//
//	return tcs
//}
//
//func (tcs tilesetCreationScreen) load() {
//	// Implement Screen interface
//}
//
//func (tcs tilesetCreationScreen) unload() {
//	// Implement Screen interface
//}
//
//func (tcs tilesetCreationScreen) emptyImg() image.Image {
//	return image.NewRGBA(image.Rect(0, 0, tcs.mc.tileWidth, tcs.mc.tileHeight))
//}
//
//func (tcs *tilesetCreationScreen) addTile(tile image.Image) {
//	tcs.tiles = append(tcs.tiles, tile)
//}
//
//func (tcs tilesetCreationScreen) imgFromPath(path string) image.Image {
//	file, _ := os.Open(path)
//	img, _, _ := image.Decode(file)
//
//	return img
//}
//
//func (tcs tilesetCreationScreen) tilesetImg() image.Image {
//	tileNb := len(tcs.tiles)
//	newImgWidth := tileNb % defaultColumn
//	newImgHeight := (tileNb / defaultColumn) + 1
//	newImg := image.NewRGBA(image.Rect(0, 0, newImgWidth * tcs.mc.tileWidth, newImgHeight * tcs.mc.tileHeight))
//
//	for i := 0; i < tileNb; i++ {
//		src := tcs.tiles[i]
//		b := src.Bounds()
//		pos := image.Point{(i % defaultColumn) * tcs.mc.tileWidth, (i / defaultColumn) * tcs.mc.tileHeight}
//
//		draw.Draw(newImg, b, src, pos, draw.Src)
//	}
//
//	return newImg
//}
//
//func (tcs *tilesetCreationScreen) tick() {
//	raygui.Label(rl.Rectangle{10, 10, 80, 10}, fmt.Sprintf("Tile configuration (pixel): %d (w) - %d (h)", tcs.mc.tileWidth, tcs.mc.tileHeight))
//	raygui.Label(rl.Rectangle{10, 30, 30, 10}, fmt.Sprintf("Current tileset: "))
//}

