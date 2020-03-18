package main

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestRegister(t *testing.T) {
	client := resty.New()

	code, _ := register(client, "fakeaddress", "fakeaccount")

	if code == "" {
		t.Errorf("The registration not returns the code")
	}
}

func TestRegisterWithoutAddress(t *testing.T) {
	client := resty.New()

	_, err := register(client, "", "fakeaccount")

	if err == nil || err.Error() != "Address not informed" {
		t.Errorf("The not informed address handler is not working")
	}
}

func TestRegisterWithoutAccount(t *testing.T) {
	client := resty.New()

	_, err := register(client, "fakeaddress", "")

	if err == nil || err.Error() != "Account not informed" {
		t.Errorf("The not informed account handler is not working")
	}
}
