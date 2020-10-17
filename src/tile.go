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
	alreadyContains := false

	for i := 0; i < len(t.Properties); i++ {
		if t.Properties[i].Name == property.Name {
			alreadyContains = true
		}
	}

	if !alreadyContains {
		t.Properties = append(t.Properties, property)
	}
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
		r.DrawTextureRec(tileset.texture, subRec, pos, r.White)

		switch len(t.Properties) {
		case 0:
			// Nothing to do
		case 1:
			r.DrawRectangle(x, y, tileset.TileWidth, tileset.TileHeight, t.PropertyColor(0))
		case 2:
			w := tileset.TileWidth / 2
			r.DrawRectangle(x, y, w, tileset.TileHeight, t.PropertyColor(0))
			r.DrawRectangle(x + w, y, w, tileset.TileHeight, t.PropertyColor(1))
		case 3:
			w := tileset.TileWidth / 2
			h := tileset.TileHeight / 2
			r.DrawRectangle(x, y, w, h, t.PropertyColor(0))
			r.DrawRectangle(x + w, y, w, h, t.PropertyColor(1))
			r.DrawRectangle(x, y + h, tileset.TileWidth, h, t.PropertyColor(2))
		case 4:
			w := tileset.TileWidth / 2
			h := tileset.TileHeight / 2
			r.DrawRectangle(x, y, w, h, t.PropertyColor(0))
			r.DrawRectangle(x + w, y, w, h, t.PropertyColor(1))
			r.DrawRectangle(x, y + h, w, h, t.PropertyColor(2))
			r.DrawRectangle(x + w, y + h, w, h, t.PropertyColor(3))
		default:
			// Find a way to handle (or block) more than 4 properties
		}
	}
}

func (t tile) PropertyColor(propertyIndex int) r.Color {
	color := t.Properties[propertyIndex].Color
	return r.NewColor(color.R, color.G, color.B, 100)
}
