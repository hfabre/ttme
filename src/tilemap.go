package ttme

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

type tilemap struct {
	width, height int // in tile
	tileset tileset
	// actions (Array<Object> (custom tiles action) [name, color, attributes(json)])
}

func (tm tilemap) Draw() {
	color := r.Black

	for y := 0; y < tm.height; y++ {
		for x := 0; x < tm.width; x++ {
			switch y {
			case 0:
				color = r.White
			case 9:
				color = r.Blue
			case 19:
				color = r.Green
			case 29:
				color = r.Yellow
			case 39:
				color = r.Orange
			case 49:
				color = r.Red
			default:
				color = r.Black
			}

			if x == 45 {
				color = r.Gray
			}
			r.DrawRectangle(x * tm.tileset.tileWidth, y * tm.tileset.tileHeight, tm.tileset.tileWidth, tm.tileset.tileHeight, color)
		}
	}
}

func (tm tilemap) PixelWidth() int {
	return tm.width * tm.tileset.tileWidth
}

func (tm tilemap) PixelHeight() int {
	return tm.height * tm.tileset.tileHeight
}