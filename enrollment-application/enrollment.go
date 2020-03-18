package main

import (
	"fmt"
	"log"
)

func register(address string, account string) (code string, err error) {

	if address == "" {
		err = fmt.Errorf("Address not informed")
		return
	}

	if account == "" {
		err = fmt.Errorf("Account not informed")
		return
	}

	log.Printf("Initing registration with address:%s account:%s\n", address, account)
	code = "JSD2342"
	return
}
