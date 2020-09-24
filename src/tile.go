package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tile struct {
	Index      int            `json:"index"`
	Properties []tileProperty `json:"properties"`
}

func (t tile) Index32() int32 {
	return int32(t.Index)
}

func (t tile) GetTilsetPosition(tileset tileset) (float32, float32) {
	tilesetWidth := tileset.TilesByLine()

	tileX := float32((t.Index32() % tilesetWidth) * int32(tileset.TileHeight))
	tileY := float32((t.Index32() / tilesetWidth) * int32(tileset.TileHeight))

	return tileX, tileY
}

func (t *tile) AddProperty(property tileProperty) {
	t.Properties = append(t.Properties, property)
}

func (t tile) Draw(x, y int, tileset tileset) {

	if t.Index == -1 {
		r.DrawRectangle(x, y, tileset.TileWidth, tileset.TileHeight, r.Black)
	} else {
		x32 := float32(x)
		y32 := float32(y)

		tileX, tileY := t.GetTilsetPosition(tileset)

		pos := r.Vector2{X: x32, Y: y32}
		subRec := r.Rectangle{X: tileX, Y: tileY, Width: float32(tileset.TileWidth), Height: float32(tileset.TileHeight)}
		color := r.White

		if len(t.Properties) > 0 {
			color = t.Properties[0].Color
		}

		r.DrawTextureRec(tileset.texture, subRec, pos, color)
	}
}
