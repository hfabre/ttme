package ttme

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

type tilemapWidget struct {
	x, y, width, height int
	tilemap tilemap
	panelScroll r.Vector2
	camera r.Camera2D
	targetTexture r.RenderTexture2D
	view r.Rectangle
}

func NewTilemapWidget(x, y, width, height int, tilemap tilemap) *tilemapWidget {
	newWidget := tilemapWidget{x: x, y: y, tilemap: tilemap}
	newWidget.width = width
	newWidget.height = height
	newWidget.panelScroll = r.Vector2{X: 0, Y: 0}
	newWidget.camera = r.Camera2D{
		Offset:   r.NewVector2Zero(),
		Target:   r.NewVector2Zero(),
		Rotation: 0,
		Zoom:     1,
	}
	newWidget.targetTexture = r.LoadRenderTexture(width, height)
	return &newWidget
}

func (tmw *tilemapWidget) Draw() {
	content := r.Rectangle{X: 0, Y: 0, Width: float32(tmw.tilemap.PixelWidth()), Height: float32(tmw.tilemap.PixelHeight())}
	tmw.view, tmw.panelScroll = r.GuiScrollPanel(tmw.Bounds(), content, tmw.panelScroll)
	tmw.camera.Target = r.Vector2{X: -tmw.panelScroll.X, Y: -tmw.panelScroll.Y}

	r.BeginTextureMode(tmw.targetTexture)
	r.BeginMode2D(tmw.camera)
	tmw.tilemap.Draw()
	r.EndMode2D()
	r.EndTextureMode()

	textureSize := tmw.view
	textureSize.Height = -textureSize.Height
	r.DrawTextureRec(tmw.targetTexture.Texture, textureSize, r.Vector2{X: float32(tmw.x), Y: float32(tmw.y)}, r.White)
}

func (tmw tilemapWidget) Bounds() r.Rectangle {
	return r.Rectangle{X: float32(tmw.x), Y: float32(tmw.y), Width: float32(tmw.width), Height: float32(tmw.height)}
}
