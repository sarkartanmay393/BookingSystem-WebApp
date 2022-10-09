package main

import "testing"

func TestRunMain(t *testing.T) {
	err := RunMain()
	if err != nil {
		t.Error("Failed runMain() function.")
	}
}
