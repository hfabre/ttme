package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tilemapConfigurationWidget struct {
	x, y int
	tilemap *tilemap
	mapWidth, mapHeight int
	mapWidthBoxEditMode, mapHeightBoxEditMode bool
}

func NewTilemapConfigurationWidget(x, y int, tilemap *tilemap) *tilemapConfigurationWidget {
	newWidget := tilemapConfigurationWidget{x: x, y: y, tilemap: tilemap, mapWidth: tilemap.width, mapHeight: tilemap.height}
	newWidget.mapWidthBoxEditMode = false
	newWidget.mapHeightBoxEditMode = false

	return &newWidget
}

// TODO: implement a messaging system to avoid passing other widget ?
func (tmcw *tilemapConfigurationWidget) Draw(widget *tilemapWidget) {
	tmcw.mapWidthBoxEditMode, tmcw.mapWidth = r.GuiSpinner(r.Rectangle{X: float32(tmcw.x), Y: float32(tmcw.y), Width: 125, Height: 25}, "Map width: ", tmcw.mapWidth, 1, 124, tmcw.mapWidthBoxEditMode)
	tmcw.mapHeightBoxEditMode, tmcw.mapHeight = r.GuiSpinner(r.Rectangle{X: float32(tmcw.x), Y: float32(tmcw.y) + 30, Width: 125, Height: 25}, "Map height: ", tmcw.mapHeight, 1, 124, tmcw.mapHeightBoxEditMode)

	if (r.GuiButton(r.Rectangle{X: float32(tmcw.x) + 140, Y: float32(tmcw.y) + 30, Width: 50, Height: 25}, "Save")) {
		tmcw.tilemap.ChangeSize(tmcw.mapWidth, tmcw.mapHeight)
		widget.Update()
	}
}