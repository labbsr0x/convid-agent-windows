package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"
)

// Agent is an running app on the machine
type Agent struct {
	filename string
	runtime  *wails.Runtime
	logger   *wails.CustomLogger
	watcher  *fsnotify.Watcher
}

// NewAgent attempts to create a new Agent instance
func NewAgent() (*Agent, error) {
	result := &Agent{}
	return result, nil
}

// LoadConfig loads the saved configurations of the instance, like the machineId its is set
func (t *Agent) LoadConfig() (string, error) {
	t.logger.Infof("Loading config from: %s", t.filename)
	bytes, err := ioutil.ReadFile(t.filename)
	if err != nil {
		err = fmt.Errorf("Unable to open config: %s", t.filename)
	}
	machineID := string(bytes)
	t.runtime.Window.SetTitle("Convid - " + machineID)
	return machineID, err
}

//SaveConfig saves the current configuration
func (t *Agent) SaveConfig(machineID string, ssHost string, sshPort string, tunnelPort string) error {
	return ioutil.WriteFile(t.filename, []byte(machineID), 0600)
}

func (t *Agent) resolveFile() error {
	_, err := os.Stat(t.filename)
	if os.IsNotExist(err) {
		err = ioutil.WriteFile(t.filename, []byte(""), 0600)
		if err != nil {
			return err
		}
		t.logger.Infof("File created and initialized: %s", t.filename)
		return nil
	}
	return err
}

// WailsInit initiates a new instance of the App resources
func (t *Agent) WailsInit(runtime *wails.Runtime) error {
	logrus.Infof("Initializing Convid...")
	t.runtime = runtime
	t.logger = t.runtime.Log.New("Agent")
	t.logger.Info("Starting Convid Agent")

	homedir, err := runtime.FileSystem.HomeDir()
	t.logger.Infof("Homedir resolved: %s", homedir)
	if err != nil {
		return err
	}
	t.filename = path.Join(homedir, "convid-machine")
	t.logger.Infof("filename resolved: %s", t.filename)
	t.runtime.Window.SetTitle("Convid Remote Desktop Provider")
	return t.resolveFile()
}
