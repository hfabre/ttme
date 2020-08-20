package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tileset struct {
	tileWidth, tileHeight int // in pixel
	imagePath string
	texture r.Texture2D
}

func NewTileset(tileWidth, tileHeight int, imagePath string) *tileset {
	newTileset := tileset{tileWidth:  tileWidth, tileHeight: tileHeight, imagePath: imagePath}
	newTileset.texture = r.LoadTexture(imagePath)

	return &newTileset
}

func (t *tileset) ChangeImage(imagePath string) {
	r.UnloadTexture(t.texture)
	t.imagePath = imagePath
	t.texture = r.LoadTexture(t.imagePath)
}

func (t tileset) PixelWidth() int {
	return int(t.texture.Width)
}

func (t tileset) PixelHeight() int {
	return int(t.texture.Height)
}

func (t tileset) TilesByLine() int32 {
	return t.texture.Width / int32(t.tileWidth)
}

func (t tileset) Draw() {
	r.DrawTexture(t.texture, 0, 0, r.White)
}
