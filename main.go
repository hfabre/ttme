package main

import (
	"github.com/sqweek/dialog"
	"os"
	ttme "ttme/src"
)

func main() {
	// TODO: bug with gtk3
	dialog.File().Load()

	fileToLoad := ""
	if len(os.Args) > 1 {
		fileToLoad = os.Args[1]
	}

	app := ttme.NewApp(1400, 900, "ttme")
	app.Start(fileToLoad)
}
