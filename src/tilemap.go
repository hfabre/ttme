package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tileset struct {
	tileWidth, tileHeight int // in pixel
	imagePath string
	//image r.Texture2D
}

func NewTileset(tileWidth, tileHeight int, imagePath string) *tileset {
	return &tileset{tileWidth:  tileWidth, tileHeight: tileHeight, imagePath: imagePath}
}

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
			case 10:
				color = r.Blue
			case 20:
				color = r.Green
			case 30:
				color = r.Yellow
			case 40:
				color = r.Orange
			case 50:
				color = r.Red
			default:
				color = r.Black
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