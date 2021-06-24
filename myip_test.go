package myip

import "testing"

func TestGetMyIP(t *testing.T) {
	_, err := GetMyIP()

	if err != nil {
		t.Error(err)
	}
}
