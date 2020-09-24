package ttme

import r "github.com/lachee/raylib-goplus/raylib"
import "github.com/sqweek/dialog"

type tilesetConfigurationWidget struct {
	x, y int
	tileset *tileset
	tileWidth, tileHeight int
	tileWidthBoxEditMode, tileHeightBoxEditMode bool
	newTilsesetPath string
}

func NewTilesetConfigurationWidget(x, y int, tileset *tileset) *tilesetConfigurationWidget {
	newWidget := tilesetConfigurationWidget{x: x, y: y, tileset: tileset}
	newWidget.tileWidth = tileset.TileWidth
	newWidget.tileHeight = tileset.TileHeight
	newWidget.tileWidthBoxEditMode = false
	newWidget.tileHeightBoxEditMode = false
	newWidget.newTilsesetPath = ""

	return &newWidget
}

// TODO: implement a messaging system to avoid passing other widget ?
func (tscw *tilesetConfigurationWidget) Draw(widget *tilemapWidget) {
	if (r.GuiButton(r.Rectangle{X: float32(tscw.x), Y: float32(tscw.y) + 50, Width: 75, Height: 25}, "Select tileset")) {
		tscw.newTilsesetPath, _ = dialog.File().Load()
	}

	tscw.tileWidthBoxEditMode, tscw.tileWidth = r.GuiSpinner(r.Rectangle{X: float32(tscw.x) + 180, Y: float32(tscw.y) + 50, Width: 125, Height: 25}, "Tile width: ", tscw.tileWidth, 1, 124, tscw.tileWidthBoxEditMode)
	tscw.tileHeightBoxEditMode, tscw.tileHeight = r.GuiSpinner(r.Rectangle{X: float32(tscw.x) + 180, Y: float32(tscw.y) + 80, Width: 125, Height: 25}, "Tile height: ", tscw.tileHeight, 1, 124, tscw.tileHeightBoxEditMode)

	if (r.GuiButton(r.Rectangle{X: float32(tscw.x) + 320, Y: float32(tscw.y) + 80, Width: 50, Height: 25}, "Save")) {
		tscw.tileset.TileWidth = tscw.tileWidth
		tscw.tileset.TileHeight = tscw.tileHeight
		if tscw.newTilsesetPath != "" {
			tscw.tileset.ChangeImage(tscw.newTilsesetPath)
			tscw.newTilsesetPath = ""
		}

		widget.Update()
	}
}