package main

import (
	"bytes"
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
		"sshHost":     fmt.Sprintf("%s", resp["sshHost"]),
		"sshPort":     fmt.Sprintf("%s", resp["sshPort"]),
		"sshUsername": fmt.Sprintf("%s", resp["sshUsername"]),
		"sshPassword": fmt.Sprintf("%s", resp["token"]),
		"tunnelPort":  fmt.Sprintf("%s", resp["tunnelPort"]),
		"withTotp":    fmt.Sprintf("%s", resp["withTotp"]),
	}

	logrus.Infof("SSH Information received. Host: %s | Port: %s | User: %s | Password: %s | TunnelPort: %s", result["sshHost"], result["sshPort"], result["sshUsername"], result["sshPassword"], result["tunnelPort"])

	if resp["withTotp"] == false {
		handleConnection(result)
	} else {
		return result, nil
	}

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
	go connect(sshHost, sshPortInt, sshUser, sshPassword, "127.0.0.1", 43389, "127.0.0.1", tunnelPortInt)
	return nil
}

func connectTotp(address string, machineID string, totpCode string) (result map[string]string, err error) {
	if address == "" {
		err = fmt.Errorf("Address not informed")
		return
	}

	if machineID == "" {
		err = fmt.Errorf("Machine ID not informed")
		return
	}

	if totpCode == "" {
		err = fmt.Errorf("TOTP Code not informed")
		return
	}

	log.Printf("Initializing registration with address:%s machineID:%s\n", address, machineID)

	schematicAddress := address
	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		schematicAddress = "http://" + address
	}
	generateMachineIDURL := fmt.Sprintf("%s/machine/%s/token", schematicAddress, strings.TrimSpace(machineID))
	logrus.Debugf("Requesting account info to %s", generateMachineIDURL)

	var jsonReq = []byte(fmt.Sprintf("{ \"code\": \"%s\" }", totpCode))
	buf := bytes.NewReader(jsonReq)

	req, err := http.NewRequest("POST", generateMachineIDURL, buf)
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
		"sshHost":     fmt.Sprintf("%s", resp["sshHost"]),
		"sshPort":     fmt.Sprintf("%s", resp["sshPort"]),
		"sshUsername": machineID,
		"sshPassword": fmt.Sprintf("%s", resp["token"]),
		"tunnelPort":  fmt.Sprintf("%s", resp["machinePort"]),
	}

	logrus.Debugf("Result: %v", result)

	handleConnection(result)
	return
}

func handleConnection(result map[string]string) {
	err := estabelishSSHTunnel(result["sshHost"], result["sshPort"], result["sshUsername"], result["sshPassword"], result["tunnelPort"])
	if err != nil {
		logrus.Errorf("Error estabilishing tunnel. Details: %s", err)
	}

	agentInstance.runtime.Events.On("ConnectionSucceed", func(optionalData ...interface{}) {

		logrus.Infof("Connection estabilished to SSH server tunneling to port %s", result["tunnelPort"])
		if runtime.GOOS == "windows" { // invoke mstsc
			c := exec.Command("mstsc", "/v:127.0.0.1:43389")
			if err := c.Run(); err != nil {
				logrus.Infof("Error callinsg MSTSC: %s", err)
			}
		}

	})
}
