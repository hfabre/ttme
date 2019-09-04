package main

type screen interface {
	tick()
	load()
	unload()
}
