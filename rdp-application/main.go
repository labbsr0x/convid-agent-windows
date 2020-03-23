package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/runtime"
)

var agentInstance *Agent

func doLoadConfig() map[string]string {
	err := agentInstance.LoadConfig()
	if err != nil {
		return map[string]string{
			"error": err.Error(),
		}
	}
	return map[string]string{
		"address":   agentInstance.config.Address,
		"machineID": agentInstance.config.MachineID,
	}
}

func doRegister(address string, machineID string) map[string]string {
	ret, err := register(address, machineID)
	if err != nil {
		return map[string]string{
			"error": err.Error(),
		}
	}
	return ret
}

func main() {

	var err error
	agentInstance, err = NewAgent()

	if err != nil {
		runtime.NewLog().New("Agent").Warn("Could not initialize agent")
		return
	}

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  500,
		Height: 500,
		Title:  "Convid Remote Desktop Provider",
		JS:     js,
		CSS:    css,
	})
	app.Bind(agentInstance)
	app.Bind(doLoadConfig)
	app.Bind(doRegister)
	app.Run()
}
