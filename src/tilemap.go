package ttme

type tilemap struct {
	width, height int // in tile
	tileset *tileset
	tiles [][]tile
	// actions (Array<Object> (custom tiles action) [name, color, attributes(json)])
}

func NewTilemap(width, height int, tileset *tileset) *tilemap {
	newTilemap := tilemap{width: width, height: height, tileset: tileset}
	newTilemap.FillEmptyBoard()

	return &newTilemap
}

func (tm *tilemap) ChangeSize(width, height int) {
	tm.width = width
	tm.height = height
	newBoard := make([][]tile, tm.height)

	for y := 0; y < tm.height; y++ {
		newBoard[y] = make([]tile, tm.width)

		for x := 0; x < tm.width; x++ {
			if y < len(tm.tiles) && x < len(tm.tiles[y]) {
				newBoard[y][x] = tm.tiles[y][x]
			} else {
				newBoard[y][x] = tile{index: -1}
			}
		}
	}

	tm.tiles = newBoard
}

func (tm *tilemap) FillEmptyBoard() {
	tm.tiles = make([][]tile, tm.height)
	for y := 0; y < tm.height; y++ {
		tm.tiles[y] = make([]tile, tm.width)
		for x := 0; x < tm.width; x++ {
			tm.tiles[y][x] = tile{index: -1}
		}
	}
}

func (tm tilemap) Draw() {
	for y := 0; y < tm.height; y++ {
		for x := 0; x < tm.width; x++ {
			tm.tiles[y][x].Draw(x * tm.tileset.tileWidth, y * tm.tileset.tileHeight, *tm.tileset)
		}
	}
}

func (tm tilemap) PixelWidth() int {
	return tm.width * tm.tileset.tileWidth
}

func (tm tilemap) PixelHeight() int {
	return tm.height * tm.tileset.tileHeight
}