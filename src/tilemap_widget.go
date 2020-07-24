package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type tilemapWidget struct {
	x, y, width, height int
	tilemap tilemap
	panelScroll r.Vector2
}

func NewTilemapWidget(x, y int, tilemap tilemap) *tilemapWidget {
	newWidget := tilemapWidget{x: x, y: y, tilemap: tilemap}
	newWidget.width = 790
	newWidget.height = 700
	newWidget.panelScroll = r.Vector2{X: 0, Y: 0}

	return &newWidget
}

func (tmw tilemapWidget) Draw() {
	content := r.Rectangle{X: 0, Y: 0, Width: float32(tmw.tilemap.PixelWidth()), Height: float32(tmw.tilemap.PixelHeight())}
	view, newScrollState := r.GuiScrollPanel(tmw.Bounds(), content, tmw.panelScroll)
	tmw.panelScroll = newScrollState

	r.BeginScissorMode(int(view.X), int(view.Y), int(view.Width), int(view.Height))
	tmw.tilemap.Draw()
	r.EndScissorMode()
}

func (tmw tilemapWidget) Bounds() r.Rectangle {
	return r.Rectangle{X: float32(tmw.x), Y: float32(tmw.y), Width: float32(tmw.width), Height: float32(tmw.height)}
}
