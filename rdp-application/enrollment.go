package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func register(address string, machineID string) (result map[string]string, err error) {

	if address == "" {
		err = fmt.Errorf("Address not informed")
		return
	}

	if machineID == "" {
		err = fmt.Errorf("Machine ID not informed")
		return
	}

	agentInstance.SaveConfig(address, machineID)

	log.Printf("Initializing registration with address:%s machineID:%s\n", address, machineID)

	schematicAddress := address
	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		schematicAddress = "http://" + address
	}
	generateMachineIDURL := fmt.Sprintf("%s/machine/%s", schematicAddress, strings.TrimSpace(machineID))
	logrus.Debugf("Requesting account info to %s", generateMachineIDURL)
	req, err := http.NewRequest("GET", generateMachineIDURL, nil)
	if err != nil {
		logrus.Errorf("Error preparing request: %s", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Error requesting: %s", err)
		return
	}
	var resp map[string]interface{}
	json.NewDecoder(response.Body).Decode(&resp)
	if response.StatusCode != 200 {
		err = fmt.Errorf("Error requesting: %d", response.StatusCode)
		return
	}

	result = map[string]string{
		"sshHost": fmt.Sprintf("%s", resp["sshHost"]),
		// "sshHost":     fmt.Sprintf("%s", "tunnel-convid.sandman.erro"),
		"sshPort":     fmt.Sprintf("%s", resp["sshPort"]),
		"sshUsername": fmt.Sprintf("%s", resp["sshUsername"]),
		"sshPassword": fmt.Sprintf("%s", resp["sshPassword"]),
		"tunnelPort":  fmt.Sprintf("%s", resp["tunnelPort"]),
	}

	logrus.Infof("SSH Information received. Host: %s | Port: %s | User: %s | Password: %s | TunnelPort: %s", result["sshHost"], result["sshPort"], result["sshUsername"], result["sshPassword"], result["tunnelPort"])

	err = estabelishSSHTunnel(result["sshHost"], result["sshPort"], result["sshUsername"], result["sshPassword"], result["tunnelPort"])
	if err != nil {
		logrus.Errorf("Error estabilishing tunnel. Details: %s", err)
	}

	agentInstance.runtime.Events.On("ConnectionSucceed", func(optionalData ...interface{}) {

		logrus.Infof("Connection estabilished to SSH server tunneling to port %s", result["tunnelPort"])
		if runtime.GOOS == "windows" { // invoke mstsc
			c := exec.Command("mstsc", "/v:127.0.0.1:43389")
			if err := c.Run(); err != nil {
				logrus.Infof("Error callinsg MSTSC: ", err)
			}
		}
	})
	return
}

//estabelishSSHTunnel has the SSH logic with remote tunnel
func estabelishSSHTunnel(sshHost string, sshPort string, sshUser string, sshPassword string, tunnelPort string) error {
	sshPortInt, err := strconv.Atoi(sshPort)
	if err != nil {
		return err
	}
	tunnelPortInt, err := strconv.Atoi(tunnelPort)
	if err != nil {
		return err
	}
	go connect(sshHost, sshPortInt, sshUser, sshPassword, "localhost", 43389, "localhost", tunnelPortInt)
	return nil
}
