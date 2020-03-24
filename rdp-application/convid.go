package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"

	"encoding/json"
)

// Agent is an running app on the machine
type Agent struct {
	filename string
	runtime  *wails.Runtime
	logger   *wails.CustomLogger
	watcher  *fsnotify.Watcher
	config   *AgentConfig
}

// AgentConfig is the configuration of running app
type AgentConfig struct {
	Address   string `json:"address"`
	MachineID string `json:"machine-id"`
}

// NewAgent attempts to create a new Agent instance
func NewAgent() (*Agent, error) {
	result := &Agent{}
	return result, nil
}

// LoadConfig loads the saved configurations of the instance, like the machineId its is set
func (t *Agent) LoadConfig() (err error) {
	t.logger.Infof("Loading config from: %s", t.filename)

	t.config = &AgentConfig{}

	_, err = os.Stat(t.filename)
	if os.IsNotExist(err) {
		t.logger.Warn("Not founded config file for load!")
		return nil
	}

	bytes, err := ioutil.ReadFile(t.filename)
	if err != nil {
		return fmt.Errorf("Unable to open config: %s", t.filename)
	}

	err = json.Unmarshal(bytes, t.config)
	if err != nil {
		return fmt.Errorf("Unable to parse config file: %v", err)
	}

	t.logger.Infof("Loaded config: %v", t.config)

	// t.runtime.Window.SetTitle("Convid - " + machineID)
	return err
}

//SaveConfig saves the current configuration
func (t *Agent) SaveConfig(address string, machineID string) error {
	t.logger.Infof("Saving config in: %s", t.filename)

	t.config.Address = address
	t.config.MachineID = machineID

	t.logger.Infof("Storing config: %v", t.config)

	d, err := json.Marshal(t.config)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return ioutil.WriteFile(t.filename, d, 0600)
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

	runID := time.Now().Format("run-2006-01-02-15-04-05")
	logLocation := filepath.Join(homedir, runID+".log")
	logFile, err := os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		logger.GlobalLogger.Fatal(err)
	}
	logger.GlobalLogger.SetOutput((io.MultiWriter(os.Stdout, logFile)))
	logrus.SetOutput((io.MultiWriter(os.Stdout, logFile)))
	logrus.RegisterExitHandler(func() {
		if logFile != nil {
			logFile.Close()
		}
	})

	t.filename = path.Join(homedir, "convid-machine-client.json")
	t.logger.Infof("filename resolved: %s", t.filename)
	t.runtime.Window.SetTitle("Convid Remote Desktop Provider")
	return nil
}
