package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	cmd := exec.Command("code", "/Users/tiagostutz/workspace/bb/github/labbsr0x")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return "FATAL007" + fmt.Sprintf("%s", err)
	}
	return "OK007"
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
	app.Bind(basic)
	app.Run()
}
