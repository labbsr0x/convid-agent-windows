package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func doRegister(address string, accountID string) map[string]string {
	ret, err := register(address, accountID)
	if err != nil {
		return map[string]string{
			"error": err.Error(),
		}
	}
	return ret
}

func main() {

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  500,
		Height: 500,
		Title:  "Convid Remote Desktop Provider",
		JS:     js,
		CSS:    css,
	})
	app.Bind(doRegister)
	app.Run()
}
