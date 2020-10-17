package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tileset struct {
	TileWidth  int    `json:"tileWidth"`  // in pixel
	TileHeight int    `json:"tileHeight"` // in pixel
	ImagePath  string `json:"imagePath"`
	texture    r.Texture2D
	loaded, needsRedraw bool
}

func NewTileset(tileWidth, tileHeight int, imagePath string) *tileset {
	newTileset := tileset{TileWidth: tileWidth, TileHeight: tileHeight, ImagePath: imagePath}

	if len(newTileset.ImagePath) > 0 {
		newTileset.LoadTexture()
	}

	return &newTileset
}

func (t *tileset) LoadTexture() {
	t.texture = r.LoadTexture(t.ImagePath)
	t.loaded = true
}

func (t *tileset) ChangeImage(imagePath string) {
	r.UnloadTexture(t.texture)
	t.loaded = false
	t.ImagePath = imagePath
	t.LoadTexture()
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
