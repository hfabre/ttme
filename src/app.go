package ttme

import (
	"encoding/json"
	"fmt"
	r "github.com/lachee/raylib-goplus/raylib"
	"io/ioutil"
	"os"
)

type app struct {
	width, height int
	name string
	mousePressed bool
	mousePosition r.Vector2
	rightMousePressed bool
}

func NewApp(width, height int, name string) *app {
	r.InitWindow(width, height, name)

	return &app{
		width:         width,
		height:        height,
		name:          name,
		mousePressed:  false,
		mousePosition: r.NewVector2Zero(),
		rightMousePressed: false,
	}
}

func (a app) ShowInfo() {
	mousePosInfo := fmt.Sprintf("Mouse position: %f - %f", a.mousePosition.X, a.mousePosition.Y)
	mouseStateInfo := fmt.Sprintf("Mouse pressed: %v", a.mousePressed)

	r.DrawText(mousePosInfo, 10, 10, 10, r.Black)
	r.DrawText(mouseStateInfo, 10, 30, 10, r.Black)
}

func (a app) Init(ts *tileset, tsw *tilesetWidget, tm *tilemap, tmw *tilemapWidget, tscw *tilesetConfigurationWidget, tpw *tilePropertiesWidget, tmcw *tilemapConfigurationWidget, sw *showWidget) {
	*ts = *NewTileset(16, 16, "./assets/tilesetpkm.png")
	*tsw = *NewTilesetWidget(30, 425, 370, 370, ts)
	*tm = *NewTilemap(50, 50, ts)
	*tmw = *NewTilemapWidget(420, 95, 800, 700, tm)
	*tscw = *NewTilesetConfigurationWidget(30, 750, ts)
	*tpw = *NewTilePropertiesWidget(25, 200)
	*tmcw = *NewTilemapConfigurationWidget(80, 100, tm)
	*sw = showWidget{x: 550, y: 800, showGrid: false, showProperties: true}
}

func (a app) InitFromFile(path string, ts *tileset, tsw *tilesetWidget, tm *tilemap, tmw *tilemapWidget, tscw *tilesetConfigurationWidget, tpw *tilePropertiesWidget, tmcw *tilemapConfigurationWidget, sw *showWidget) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()

	var loadedTm tilemap
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &loadedTm)

	*tm = loadedTm
	tm.Tileset.LoadTexture()
	*ts = *tm.Tileset
	*tsw = *NewTilesetWidget(30, 425, 370, 370, ts)
	*tmw = *NewTilemapWidget(420, 95, 800, 700, tm)
	*tscw = *NewTilesetConfigurationWidget(30, 750, ts)
	*tpw = *NewTilePropertiesWidget(25, 200)
	*tmcw = *NewTilemapConfigurationWidget(80, 100, tm)
	*sw = showWidget{x: 550, y: 800, showGrid: false, showProperties: true}
}

func (a *app) Start(fileToLoad string) {
	tileset := tileset{}
	tilesetWidget := tilesetWidget{}
	tilemap := tilemap{}
	tilemapWidget := tilemapWidget{}
	tilsetConfigurationWidget := tilesetConfigurationWidget{}
	tilePropertiesWidget := tilePropertiesWidget{}
	tilemapConfigurationWidget := tilemapConfigurationWidget{}
	showWidget := showWidget{}
	mouseInTileMap := false

	// TODO: Not really proud of this way to init app state
	if len(fileToLoad) == 0 {
		a.Init(&tileset, &tilesetWidget, &tilemap, &tilemapWidget, &tilsetConfigurationWidget, &tilePropertiesWidget, &tilemapConfigurationWidget, &showWidget)
	} else {
		a.InitFromFile(fileToLoad, &tileset, &tilesetWidget, &tilemap, &tilemapWidget, &tilsetConfigurationWidget, &tilePropertiesWidget, &tilemapConfigurationWidget, &showWidget)
	}

	for !r.WindowShouldClose() {

		// Handle Mouse inputs

		a.mousePosition = r.GetMousePosition()

		if r.IsMouseButtonPressed(r.MouseLeftButton) {
			a.mousePressed = true

			if !tilePropertiesWidget.editMode && tilesetWidget.Contains(a.mousePosition.X, a.mousePosition.Y) {
				tilePropertiesWidget.Unset()
				tilesetWidget.SelectTile(a.mousePosition.X, a.mousePosition.Y)
			}
		}

		if r.IsMouseButtonReleased(r.MouseLeftButton) {
			a.mousePressed = false
		}

		if r.IsMouseButtonPressed(r.MouseRightButton) {
			a.rightMousePressed = true
		}

		if r.IsMouseButtonReleased(r.MouseRightButton) {
			a.rightMousePressed = false
		}

		// Handle tile pasting

		mouseInTileMap = tilemapWidget.Contains(a.mousePosition.X, a.mousePosition.Y)

		if a.mousePressed && mouseInTileMap {
			if tilePropertiesWidget.Selected() {
				tilemapWidget.SetTileFromPos(a.mousePosition.X, a.mousePosition.Y, tilesetWidget.selectedTile)
			} else {
				tilemapWidget.SetPropertyFromPos(a.mousePosition.X, a.mousePosition.Y, tilePropertiesWidget.SelectedProperty())
			}
		}

		// Handle tile removing

		if a.rightMousePressed && tilemapWidget.Contains(a.mousePosition.X, a.mousePosition.Y) {
			tilemapWidget.SetTileFromPos(a.mousePosition.X, a.mousePosition.Y, tile{Index: -1})
		}

		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		tilemapWidget.Draw(showWidget.showProperties, showWidget.showGrid)
		showWidget.Draw()
		tilesetWidget.Draw()
		tilsetConfigurationWidget.Draw(&tilemapWidget)
		tilemapConfigurationWidget.Draw(&tilemapWidget)
		tilePropertiesWidget.Draw()

		//a.ShowInfo()

		// Debug
		//tilesetinfo := fmt.Sprintf("Tileset: %v - %v / %v", tileset.TileWidth, tileset.TileHeight, tileset.ImagePath)
		//tilemapinfo := fmt.Sprintf("Tilemap: %v - %v", tilemap.Width, tilemap.Height)

		if mouseInTileMap {
			mouseTilePosInfo := fmt.Sprintf("Tile position: %v - %v", tilemapWidget.GetTileXFromPos(a.mousePosition.X), tilemapWidget.GetTileYFromPos(a.mousePosition.Y))
			r.DrawText(mouseTilePosInfo, 420, 803, 11, r.Black)
		}

		// Debug
		//r.DrawText(tilesetinfo, 10, 70, 10, r.Black)
		//r.DrawText(tilemapinfo, 10, 90, 10, r.Black)
		r.EndDrawing()
	}
	r.CloseWindow()
}