package ttme

import (
	"encoding/json"
	"fmt"
	r "github.com/lachee/raylib-goplus/raylib"
	"github.com/sqweek/dialog"
	"io/ioutil"
)

type tilemapWidget struct {
	x, y, width, height int
	tilemap *tilemap
	panelScroll r.Vector2
	targetTexture r.RenderTexture2D
	view r.Rectangle
}

func NewTilemapWidget(x, y, width, height int, tilemap *tilemap) *tilemapWidget {
	newWidget := tilemapWidget{x: x, y: y, tilemap: tilemap}
	newWidget.width = width
	newWidget.height = height
	newWidget.panelScroll = r.Vector2{X: 0, Y: 0}
	newWidget.targetTexture = r.LoadRenderTexture(tilemap.PixelWidth(), tilemap.PixelHeight())

	return &newWidget
}

func (tmw *tilemapWidget) Update() {
	r.UnloadRenderTexture(tmw.targetTexture)
	tmw.targetTexture = r.LoadRenderTexture(tmw.tilemap.PixelWidth(), tmw.tilemap.PixelHeight())
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

	if r.GuiButton(r.Rectangle{X: float32(tmw.x + tmw.width - 90), Y: float32(tmw.y + tmw.height + 10), Width: 90, Height: 25}, "Export to JSON") {
		path, _ := dialog.File().Filter("JSON files", "json").Title("Export to JSON").Save()
		file, _ := json.Marshal(*tmw.tilemap)
		_ = ioutil.WriteFile(path, file, 0644)
	}
}

func (tmw tilemapWidget) Contains(x, y float32) bool {
	point := r.Vector2{x, y}

	return r.CheckCollisionPointRec(point, tmw.view)
}

func (tmw *tilemapWidget) SetTile(x, y int, tile tile) {
	tmw.tilemap.Tiles[y][x] = tile
}

func (tmw tilemapWidget) GetTileXFromPos(x float32)  int {
	return int((x - tmw.view.X - tmw.panelScroll.X) / float32(tmw.tilemap.Tileset.TileWidth))
}

func (tmw tilemapWidget) GetTileYFromPos(y float32)  int {
	return int((y - tmw.view.Y - tmw.panelScroll.Y) / float32(tmw.tilemap.Tileset.TileHeight))
}

// Note: panelScroll Offset
func (tmw *tilemapWidget) SetTileFromPos(x, y float32, tile tile) {
	tileX := tmw.GetTileXFromPos(x)
	tileY := tmw.GetTileYFromPos(y)

	tmw.SetTile(int(tileX), int(tileY), tile)
}

func (tmw *tilemapWidget) SetTileProperty(x, y int, property tileProperty) {
	tmw.tilemap.Tiles[y][x].AddProperty(property)
}

// Note: panelScroll Offset
func (tmw *tilemapWidget) SetPropertyFromPos(x, y float32, property tileProperty) {
	tileX := (x - tmw.view.X - tmw.panelScroll.X) / float32(tmw.tilemap.Tileset.TileWidth)
	tileY := (y - tmw.view.Y - tmw.panelScroll.Y) / float32(tmw.tilemap.Tileset.TileHeight)

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
