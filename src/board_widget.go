package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type boardWidget struct {
	mc *mapConfiguration
	offsetX, offsetY int
	tiles [][]tile
}

func (b *boardWidget) initTiles() {
	b.tiles = make([][]tile, b.mc.height)
	for y := 0; y < b.mc.height; y++ {
		b.tiles[y] = make([]tile, b.mc.width)
		for x := 0; x < b.mc.width; x++ {
			b.tiles[y][x] = tile{0}
		}
	}
}

func (b boardWidget) draw(tileset tileset) {
	for y := 0; y < b.mc.height; y++ {
		for x := 0; x < b.mc.width; x++ {
			b.tiles[y][x].draw(b.offsetX + x * b.mc.tileWidth, b.offsetY + y * b.mc.tileHeight, tileset)
		}
	}

	b.drawGrid()
}

// FIXME: This only draw the top and left lines
func (b boardWidget) drawGrid() {
	for x := 0; x < b.mc.width; x++ {
		rl.DrawLine(int32(b.offsetX + x * b.mc.tileWidth), int32(b.offsetY), int32(b.offsetX + x * b.mc.tileWidth), int32(b.offsetY + b.mc.height * b.mc.tileHeight), rl.Red)
	}

	for y := 0; y < b.mc.height; y++ {
		rl.DrawLine(int32(b.offsetX), int32(b.offsetY + y * b.mc.tileHeight), int32(b.offsetX + b.mc.width * b.mc.tileWidth), int32(b.offsetY + y * b.mc.tileHeight), rl.Red)
	}
}

func (b boardWidget) contains(x, y float32) bool {
	point := rl.Vector2{x, y}
	rect := rl.Rectangle{float32(b.offsetX), float32(b.offsetY), float32(b.mc.width * b.mc.tileWidth), float32(b.mc.height * b.mc.tileHeight)}

	return rl.CheckCollisionPointRec(point, rect)
}

func (b *boardWidget) setTile(x, y int, tile tile) {
	b.tiles[y][x] = tile
}

func (b *boardWidget) setTileFromPos(x, y float32, tile tile) {
	tileX := (x - float32(b.offsetX)) / float32(b.mc.tileWidth)
	tileY := (y - float32(b.offsetY)) / float32(b.mc.tileHeight)

	b.setTile(int(tileX), int(tileY), tile)
}

func (b *boardWidget) clear() {
	for y := 0; y < b.mc.height; y++ {
		for x := 0; x < b.mc.width; x++ {
			b.tiles[y][x] = tile{0}
		}
	}
}

// TODO: Guess this could almost be done automagically using slice ?
func (b boardWidget) copy(newBoard *boardWidget) {
	for y := 0; y < len(b.tiles); y++ {
		if len(newBoard.tiles) > y {
			for x := 0; x < len(b.tiles[y]); x++ {
				if len(newBoard.tiles[y]) > x {
					newBoard.tiles[y][x] = b.tiles[y][x]
				}
			}
		}
	}
}

func emptyBoard(mc *mapConfiguration, offsetX, offsetY int) *boardWidget {
	b := &boardWidget{mc: mc, offsetX: offsetX, offsetY: offsetY}
	b.initTiles()
	return b
}
