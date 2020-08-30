package ttme

import (
	"fmt"
	r "github.com/lachee/raylib-goplus/raylib"
)

type tilemapWidget struct {
	x, y, width, height int
	tilemap tilemap
	panelScroll r.Vector2
	targetTexture r.RenderTexture2D
	view r.Rectangle
}

func NewTilemapWidget(x, y, width, height int, tilemap tilemap) *tilemapWidget {
	newWidget := tilemapWidget{x: x, y: y, tilemap: tilemap}
	newWidget.width = width
	newWidget.height = height
	newWidget.panelScroll = r.Vector2{X: 0, Y: 0}
	newWidget.targetTexture = r.LoadRenderTexture(tilemap.PixelWidth(), tilemap.PixelHeight())

	return &newWidget
}

func (tmw *tilemapWidget) Draw() {
	content := r.Rectangle{X: 0, Y: 0, Width: float32(tmw.tilemap.PixelWidth()), Height: float32(tmw.tilemap.PixelHeight())}
	tmw.view, tmw.panelScroll = r.GuiScrollPanel(tmw.Bounds(), content, tmw.panelScroll)

	r.BeginTextureMode(tmw.targetTexture)
	tmw.tilemap.Draw()
	r.EndTextureMode()

	// Note: We have to take back one because scroll offset starts at one
	verticalOffset := tmw.tilemap.PixelHeight() - int(tmw.view.Height) - 1

	// Note: Since we vertically invert texture we need to invert scroll (Y) as well,
	mapView := r.Rectangle{X: -(tmw.panelScroll.X - 1), Y: tmw.panelScroll.Y + float32(verticalOffset), Width: tmw.view.Width, Height: -tmw.view.Height}
	r.DrawTextureRec(tmw.targetTexture.Texture, mapView, r.Vector2{X: float32(tmw.x), Y: float32(tmw.y)}, r.White)
}

func (tmw tilemapWidget) Contains(x, y float32) bool {
	point := r.Vector2{x, y}

	return r.CheckCollisionPointRec(point, tmw.view)
}

func (tmw *tilemapWidget) SetTile(x, y int, tile tile) {
	tmw.tilemap.tiles[y][x] = tile
}

// Note: panelScroll Offset
func (tmw *tilemapWidget) SetTileFromPos(x, y float32, tile tile) {
	tileX := (x - tmw.view.X - tmw.panelScroll.X) / float32(tmw.tilemap.tileset.tileWidth)
	tileY := (y - tmw.view.Y - tmw.panelScroll.Y) / float32(tmw.tilemap.tileset.tileHeight)

	tmw.SetTile(int(tileX), int(tileY), tile)
}

func (tmw *tilemapWidget) SetTileProperty(x, y int, property tileProperty) {
	tmw.tilemap.tiles[y][x].AddProperty(property)
}

// Note: panelScroll Offset
func (tmw *tilemapWidget) SetPropertyFromPos(x, y float32, property tileProperty) {
	tileX := (x - tmw.view.X - tmw.panelScroll.X) / float32(tmw.tilemap.tileset.tileWidth)
	tileY := (y - tmw.view.Y - tmw.panelScroll.Y) / float32(tmw.tilemap.tileset.tileHeight)

	tmw.SetTileProperty(int(tileX), int(tileY), property)
}

func (tmw tilemapWidget) ShowInfo() {
	scrollInfo := fmt.Sprintf("Scroll state: %f - %f", tmw.panelScroll.X, tmw.panelScroll.Y)
	viewInfo := fmt.Sprintf("View rect: %f - %f - %f - %f", tmw.view.X, tmw.view.Y, tmw.view.Width, tmw.view.Height)
	mapInfo := fmt.Sprintf("Map size: %d - %d", tmw.tilemap.PixelWidth(), tmw.tilemap.PixelHeight())

	r.DrawText(scrollInfo, tmw.x + 10, tmw.y + 10, 10, r.Gray)
	r.DrawText(viewInfo, tmw.x + 10, tmw.y + 30, 10, r.Gray)
	r.DrawText(mapInfo, tmw.x + 10, tmw.y + 50, 10, r.Gray)
}

func (tmw tilemapWidget) Bounds() r.Rectangle {
	return r.Rectangle{X: float32(tmw.x), Y: float32(tmw.y), Width: float32(tmw.width), Height: float32(tmw.height)}
}
