package main

import rl "github.com/gen2brain/raylib-go/raylib"

type tileset struct {
	mc *mapConfiguration
	path string
	texture rl.Texture2D
}

func makeTileset(mc *mapConfiguration, path string) tileset {
	return tileset{mc, path, rl.LoadTexture(path)}
}

func (t tileset) unload() {
	rl.UnloadTexture(t.texture)
}

func (t tileset) tilesByLine() int32 {
	return t.texture.Width / int32(t.mc.tileWidth)
}
