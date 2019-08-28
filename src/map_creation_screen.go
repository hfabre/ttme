package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sqweek/dialog"
)

// Gui layout
const BoardOffsetX = 300
const BoardOffsetY = 10

type mapCreationScreen struct {
	tileset tileset
	mcw mapConfigurationWidget
	bw boardWidget
	tsw tileSelectorWidget
	ew exportWidget
	mousePressed bool
}

func makeMapCreationScreen() mapCreationScreen {
	tileset := makeTileset("./assets/tilesetpkm.png")
	mcw := makeMapConfigurationWidget(10, 10)
	bw := emptyBoard(mcw.width, mcw.height, BoardOffsetX, BoardOffsetY)
	tsw := tileSelectorWidget{tileset, 10, 70, tile{0}}
	ew := exportWidget{200, 80}

	return mapCreationScreen{tileset, mcw, *bw, tsw, ew, false}
}

func (mcs mapCreationScreen) unload() {
	mcs.tileset.unload()
}

func (mcs *mapCreationScreen) tick() {
	// Updating

	// Find a tileset
	// TODO: Write a widget and make sure user can specify tiles attributes (height and width)

	if raygui.Button(rl.Rectangle{float32(mcs.tsw.offsetX), float32(mcs.tsw.offsetY) - 20, 80, 15}, "Select tileset") {
		tilesetPath, _ := dialog.File().Load()
		mcs.tileset = makeTileset(tilesetPath)
		mcs.tsw = tileSelectorWidget{mcs.tileset, 10, 70, tile{0}}
	}

	// Check map size changes

	if mcs.mcw.hasChanges {
		newBoard := emptyBoard(mcs.mcw.width, mcs.mcw.height, BoardOffsetX, BoardOffsetY)
		mcs.bw.copy(newBoard)
		mcs.bw = *newBoard
		mcs.mcw.hasChanges = false
	}

	// Handle Mouse inputs

	mousePosition := rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mcs.mousePressed = true

		if mcs.tsw.contains(mousePosition.X, mousePosition.Y) {
			mcs.tsw.selectTile(mousePosition.X, mousePosition.Y)
		}
	}

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		mcs.mousePressed = false
	}

	// Handle tile pasting

	if mcs.mousePressed && mcs.bw.contains(mousePosition.X, mousePosition.Y) {
		mcs.bw.setTileFromPos(mousePosition.X, mousePosition.Y, mcs.tsw.selectedTile)
	}

	// Drawing

	mcs.mcw.draw()
	mcs.bw.draw(mcs.tileset)
	mcs.tsw.draw()
	mcs.ew.draw(mcs.mcw, mcs.bw)
}
