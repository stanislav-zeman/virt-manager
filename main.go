package main

import (
	"embed"

	"github.com/stanislav-zeman/virt-manager/internal/app"
	"github.com/stanislav-zeman/virt-manager/internal/config"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

const configPath = "config/app.yaml"

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	conf := config.Parse(configPath)
	app := app.New()
	err := wails.Run(&options.App{
		Title:  conf.App.Title,
		Width:  conf.App.Width,
		Height: conf.App.Height,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{
			R: conf.App.Background.Colour.Red,
			G: conf.App.Background.Colour.Green,
			B: conf.App.Background.Colour.Blue,
			A: conf.App.Background.Colour.Alpha,
		},
		OnStartup:  app.Startup,
		OnShutdown: app.Close,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
