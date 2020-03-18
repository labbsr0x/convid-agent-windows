package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// RegistrationResponse the reponse with code
type RegistrationResponse struct {
	Code    string
	SSHHost string
	SSHPort string
}

func register(client *resty.Client, address string, account string) (code RegistrationResponse, err error) {

	if client == nil {
		err = fmt.Errorf("Resty http client not informed")
		return
	}

	if address == "" {
		err = fmt.Errorf("Address not informed")
		return
	}

	if account == "" {
		err = fmt.Errorf("Account not informed")
		return
	}

	log.Printf("Initing registration with address:%s account:%s\n", address, account)

	client.SetDebug(true)

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&RegistrationResponse{}).
		Post(fmt.Sprintf("%s/accounts/%s/machine", address, account))

	if err != nil {
		return
	}

	code = resp.Result().(RegistrationResponse)

	return
}
