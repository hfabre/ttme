package ttme

import r "github.com/lachee/raylib-goplus/raylib"

type showWidget struct {
	showGrid, showProperties bool
	x, y int
}

func (sw *showWidget) Draw() {
	sw.showGrid = r.GuiCheckBox(r.Rectangle{X: float32(sw.x), Y: float32(sw.y), Width: 16, Height: 16}, "Show grid", sw.showGrid)
	sw.showProperties = r.GuiCheckBox(r.Rectangle{X: float32(sw.x) + 100, Y: float32(sw.y), Width: 16, Height: 16}, "Show properties", sw.showProperties)
}
