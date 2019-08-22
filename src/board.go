package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type board struct {
	width, height int
	offsetX, offsetY int
	tiles [][]tile
}

func (b *board) initTiles() {
	b.tiles = make([][]tile, b.height)
	for y := 0; y < b.height; y++ {
		b.tiles[y] = make([]tile, b.width)
		for x := 0; x < b.width; x++ {
			b.tiles[y][x] = tile{0}
		}
	}
}

func (b board) show() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			fmt.Printf("%d", b.tiles[y][x].index)
		}

		fmt.Printf("%c", '\n')
	}
}

func (b board) draw(tileset tileset) {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			b.tiles[y][x].draw(b.offsetX + x * TileWidth, b.offsetY + y * TileHeight, tileset)
		}
	}

	b.drawGrid()
}

// FIXME: This only draw the top and left lines
func (b board) drawGrid() {
	for x := 0; x < b.width; x++ {
		rl.DrawLine(int32(b.offsetX + x * TileWidth), int32(b.offsetY), int32(b.offsetX + x * TileWidth), int32(b.offsetY + b.height * TileHeight), rl.Red)
	}

	for y := 0; y < b.height; y++ {
		rl.DrawLine(int32(b.offsetX), int32(b.offsetY + y * TileHeight), int32(b.offsetX + b.width * TileWidth), int32(b.offsetY + y * TileHeight), rl.Red)
	}
}

func (b board) contains(x, y float32) bool {
	point := rl.Vector2{x, y}
	rect := rl.Rectangle{float32(b.offsetX), float32(b.offsetY), float32(b.width * TileWidth), float32(b.height * TileHeight)}

	return rl.CheckCollisionPointRec(point, rect)
}

func (b *board) setTile(x, y int, tile tile) {
	b.tiles[y][x] = tile
}

func (b *board) setTileFromPos(x, y float32, tile tile) {
	tileX := (x - float32(b.offsetX)) / TileWidth
	tileY := (y - float32(b.offsetY)) / TileHeight

	b.setTile(int(tileX), int(tileY), tile)
}

func (b *board) clear() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			b.tiles[y][x] = tile{0}
		}
	}
}

func (b board) copy(newBoard *board) {
	for y := 0; y < b.height; y++ {
		if len(newBoard.tiles) > y {
			for x := 0; x < b.width; x++ {
				if len(newBoard.tiles[y]) > x {
					newBoard.tiles[y][x] = b.tiles[y][x]
				}
			}
		}
	}
}

func emptyBoard(width, height, offsetX, offsetY int) *board {
	b := &board{width: width, height: height, offsetX: offsetX, offsetY: offsetY}
	b.initTiles()
	return b
}
