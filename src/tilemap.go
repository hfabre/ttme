package ttme

type tilemap struct {
	Width   int      `json:"width"`  // in tile
	Height  int      `json:"height"` //in tile
	Tileset *tileset `json:"tileset"`
	Tiles   [][]tile `json:"tiles"`
	// actions (Array<Object> (custom tiles action) [name, color, attributes(json)])
}

func NewTilemap(width, height int, tileset *tileset) *tilemap {
	newTilemap := tilemap{Width: width, Height: height, Tileset: tileset}
	newTilemap.FillEmptyBoard()

	return &newTilemap
}

func (tm *tilemap) ChangeSize(width, height int) {
	tm.Width = width
	tm.Height = height
	newBoard := make([][]tile, tm.Height)

	for y := 0; y < tm.Height; y++ {
		newBoard[y] = make([]tile, tm.Width)

		for x := 0; x < tm.Width; x++ {
			if y < len(tm.Tiles) && x < len(tm.Tiles[y]) {
				newBoard[y][x] = tm.Tiles[y][x]
			} else {
				newBoard[y][x] = tile{Index: -1}
			}
		}
	}

	tm.Tiles = newBoard
}

func (tm *tilemap) FillEmptyBoard() {
	tm.Tiles = make([][]tile, tm.Height)
	for y := 0; y < tm.Height; y++ {
		tm.Tiles[y] = make([]tile, tm.Width)
		for x := 0; x < tm.Width; x++ {
			tm.Tiles[y][x] = tile{Index: -1}
		}
	}
}

func (tm tilemap) Draw() {
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			tm.Tiles[y][x].Draw(x * tm.Tileset.TileWidth, y * tm.Tileset.TileHeight, *tm.Tileset)
		}
	}
}

func (tm tilemap) PixelWidth() int {
	return tm.Width * tm.Tileset.TileWidth
}

func (tm tilemap) PixelHeight() int {
	return tm.Height * tm.Tileset.TileHeight
}