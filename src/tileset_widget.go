package ttme

import (
	"fmt"
	r "github.com/lachee/raylib-goplus/raylib"
)

type tilesetWidget struct {
	x, y, width, height int
	tileset *tileset
	panelScroll r.Vector2
	targetTexture r.RenderTexture2D
	view r.Rectangle
	selectedTile tile
}

func NewTilesetWidget(x, y, width, height int, tileset *tileset) *tilesetWidget {
	newWidget := tilesetWidget{x: x, y: y, tileset: tileset}
	newWidget.width = width
	newWidget.height = height
	newWidget.panelScroll = r.Vector2{X: 0, Y: 0}
	newWidget.targetTexture = r.LoadRenderTexture(tileset.PixelWidth(), tileset.PixelHeight())
	newWidget.selectedTile = tile{index: -1}

	return &newWidget
}

func (tsw *tilesetWidget) SelectTile(x, y float32) {
	tileX := (x - float32(tsw.x)) / float32(tsw.tileset.tileWidth)
	tileY := (y - float32(tsw.y)) / float32(tsw.tileset.tileHeight)
	tilePos := int32(tileY) * tsw.tileset.TilesByLine() + int32(tileX)

	tsw.selectedTile = tile{index: int(tilePos)}
}

func (tsw tilesetWidget) Contains(x, y float32) bool {
	point := r.Vector2{X: x, Y: y}

	return r.CheckCollisionPointRec(point, tsw.view)
}

func (tsw *tilesetWidget) Draw() {
	content := r.Rectangle{X: 0, Y: 0, Width: float32(tsw.tileset.PixelWidth()), Height: float32(tsw.tileset.PixelHeight())}
	tsw.view, tsw.panelScroll = r.GuiScrollPanel(tsw.Bounds(), content, tsw.panelScroll)
	tileX, tileY := tsw.selectedTile.GetTilsetPosition(*tsw.tileset)

	r.BeginTextureMode(tsw.targetTexture)
	tsw.tileset.Draw()
	r.DrawRectangleLines(int(tileX), int(tileY), tsw.tileset.tileWidth, tsw.tileset.tileHeight, r.Red)
	r.EndTextureMode()

	// Note: We have to take back one because scroll offset starts at one
	verticalOffset := tsw.tileset.PixelHeight() - int(tsw.view.Height) - 1

	// Note: Since we vertically invert texture we need to invert scroll (Y) as well,
	mapView := r.Rectangle{X: -(tsw.panelScroll.X - 1), Y: tsw.panelScroll.Y + float32(verticalOffset), Width: tsw.view.Width, Height: -tsw.view.Height}
	r.DrawTextureRec(tsw.targetTexture.Texture, mapView, r.Vector2{X: float32(tsw.x), Y: float32(tsw.y)}, r.White)
}

func (tsw tilesetWidget) ShowInfo() {
	scrollInfo := fmt.Sprintf("Scroll state: %f - %f", tsw.panelScroll.X, tsw.panelScroll.Y)
	viewInfo := fmt.Sprintf("View rect: %f - %f - %f - %f", tsw.view.X, tsw.view.Y, tsw.view.Width, tsw.view.Height)
	mapInfo := fmt.Sprintf("set size: %d - %d", tsw.tileset.PixelWidth(), tsw.tileset.PixelHeight())
	selectedTileInfo := fmt.Sprintf("Selected tile: %v", tsw.selectedTile.index)

	r.DrawText(scrollInfo, tsw.x + 10, tsw.y + 10, 14, r.Red)
	r.DrawText(viewInfo, tsw.x + 10, tsw.y + 30, 14, r.Red)
	r.DrawText(mapInfo, tsw.x + 10, tsw.y + 50, 14, r.Red)
	r.DrawText(selectedTileInfo, tsw.x + 10, tsw.y + 70, 14, r.Red)
}

func (tsw tilesetWidget) Bounds() r.Rectangle {
	return r.Rectangle{X: float32(tsw.x), Y: float32(tsw.y), Width: float32(tsw.width), Height: float32(tsw.height)}
}
