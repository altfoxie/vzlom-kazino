package main

import (
	"embed"

	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/src
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:         "Взлом КАЗИНО",
		Width:         400,
		Height:        210,
		DisableResize: true,
		Frameless:     true,
		Assets:        assets,
		LogLevel:      logger.DEBUG,
		OnStartup:     app.startup,
		OnDomReady:    app.domReady,
		StartHidden:   true,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "Взлом КАЗИНО",
				Message: "взлом казино",
				Icon:    icon,
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
