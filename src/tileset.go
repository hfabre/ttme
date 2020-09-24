package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tileset struct {
	TileWidth  int    `json:"tileWidth"`  // in pixel
	TileHeight int    `json:"tileHeight"` // in pixel
	ImagePath  string `json:"imagePath"`
	texture    r.Texture2D
}

func NewTileset(tileWidth, tileHeight int, imagePath string) *tileset {
	newTileset := tileset{TileWidth: tileWidth, TileHeight: tileHeight, ImagePath: imagePath}
	newTileset.texture = r.LoadTexture(imagePath)

	return &newTileset
}

func (t *tileset) ChangeImage(imagePath string) {
	r.UnloadTexture(t.texture)
	t.ImagePath = imagePath
	t.texture = r.LoadTexture(t.ImagePath)
}

func (t tileset) PixelWidth() int {
	return int(t.texture.Width)
}

func (t tileset) PixelHeight() int {
	return int(t.texture.Height)
}

func (t tileset) TilesByLine() int32 {
	return t.texture.Width / int32(t.TileWidth)
}

func (t tileset) Draw() {
	r.DrawTexture(t.texture, 0, 0, r.White)
}
