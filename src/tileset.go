package main

import rl "github.com/gen2brain/raylib-go/raylib"

type tileset struct {
	path string
	texture rl.Texture2D
}

func makeTileset(path string) tileset {
	return tileset{path, rl.LoadTexture(path)}
}

func (t tileset) unload() {
	rl.UnloadTexture(t.texture)
}

func (t tileset) tilesByLine() int32 {
	return t.texture.Width / TileWidth
}
