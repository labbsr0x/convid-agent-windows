package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func register(address string, account string) (result map[string]string, err error) {

	if address == "" {
		err = fmt.Errorf("Address not informed")
		return
	}

	if account == "" {
		err = fmt.Errorf("Account not informed")
		return
	}

	log.Printf("Initing registration with address:%s account:%s\n", address, account)

	schematicAddress := address
	if !strings.HasPrefix(address, "http://") {
		schematicAddress = "http://" + address
	}
	generateMachineIDURL := fmt.Sprintf("%s/account/%s/machine", schematicAddress, account)
	req, err := http.NewRequest("POST", generateMachineIDURL, nil)
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
	}

	logrus.Infof("SSH Information received. Host: %s | Port: %s | TunnelPort: %s", result["sshHost"], result["sshPort"], result["tunnelPort"])

	return
}
