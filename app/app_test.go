package app

import (
	"testing"
)

//Test StartService function
func TestStartServiceForSuccess(t *testing.T) {
	tResult, err := StartService(true)

	if tResult == "" || err != nil {
		t.Fatalf(`StartService(true) = %q, `, err)
	}
}

//Test StartService function for failure
func TestStartServiceForFailure(t *testing.T) {
	_, err := StartService(false)

	if err == nil {
		t.Fatalf(`StartService(false) = %q, `, err)
	}
}
