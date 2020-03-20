package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	log.Printf("Initializing registration with address:%s machineID:%s\n", address, machineID)

	schematicAddress := address
	if !strings.HasPrefix(address, "http://") {
		schematicAddress = "http://" + address
	}
	generateMachineIDURL := fmt.Sprintf("%s/machine/%s", schematicAddress, machineID)
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
	json.NewDecoder(response.Body).Decode(&result)
	if response.StatusCode != 200 {
		err = fmt.Errorf("Error requesting: %d", response.StatusCode)
		return
	}

	logrus.Infof("SSH Information received. Host: %s | Port: %s | User: %s | Password: %s | TunnelPort: %s", result["sshHost"], result["sshPort"], result["sshUsername"], result["sshPassword"], result["tunnelPort"])

	// agentInstance.SaveConfig(machineID, result["sshHost"], result["sshPort"], result["sshUsername"], result["sshPassword"], result["tunnelPort"])

	err = estabelishSSHTunnel(result["sshHost"], result["sshPort"], result["sshUsername"], result["sshPassword"], result["tunnelPort"])
	logrus.Infof("Connection estabilished to SSH server tunneling to port %s", result["tunnelPort"])
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
	go connect(sshHost, sshPortInt, sshUser, sshPassword, "localhost", 3389, "localhost", tunnelPortInt)
	return nil
}
