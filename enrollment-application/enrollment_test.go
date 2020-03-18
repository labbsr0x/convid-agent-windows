package main

import (
	"testing"
)

func TestRegister(t *testing.T) {
	code, _ := register("fakeaddress", "fakeaccount")

	if code == "" {
		t.Errorf("The registration not returns the code")
	}
}

func TestRegisterWithoutAddress(t *testing.T) {
	_, err := register("", "fakeaccount")

	if err == nil || err.Error() != "Address not informed" {
		t.Errorf("The not informed address handler is not working")
	}
}

func TestRegisterWithoutAccount(t *testing.T) {
	_, err := register("fakeaddress", "")

	if err == nil || err.Error() != "Account not informed" {
		t.Errorf("The not informed account handler is not working")
	}
}
