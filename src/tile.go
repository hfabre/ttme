package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tile struct {
	index int
	properties []tileProperty
}

func (t tile) Index32() int32 {
	return int32(t.index)
}

func (t tile) GetTilsetPosition(tileset tileset) (float32, float32) {
	tilesetWidth := tileset.TilesByLine()

	tileX := float32((t.Index32() % tilesetWidth) * int32(tileset.tileHeight))
	tileY := float32((t.Index32() / tilesetWidth) * int32(tileset.tileHeight))

	return tileX, tileY
}

func (t *tile) AddProperty(property tileProperty) {
	t.properties = append(t.properties, property)
}

func (t tile) Draw(x, y int, tileset tileset) {

	if t.index == -1 {
		r.DrawRectangle(x, y, tileset.tileWidth, tileset.tileHeight, r.Black)
	} else {
		x32 := float32(x)
		y32 := float32(y)

		tileX, tileY := t.GetTilsetPosition(tileset)

		pos := r.Vector2{X: x32, Y: y32}
		subRec := r.Rectangle{X: tileX, Y: tileY, Width: float32(tileset.tileWidth), Height: float32(tileset.tileHeight)}
		color := r.White

		if len(t.properties) > 0 {
			color = t.properties[0].color
		}

		r.DrawTextureRec(tileset.texture, subRec, pos, color)
	}
}
