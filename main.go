package main

import (
	"github.com/sqweek/dialog"
	ttme "ttme/src"
)

func main() {
	dialog.File().Load()
	app := ttme.NewApp(1400, 900, "ttme")
	app.Start()
}
