package main

import "testing"

func TestRunMain(t *testing.T) {
	_, err := RunMain()
	if err != nil {
		t.Error("Failed runMain() function.")
	}
}
