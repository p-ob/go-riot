package main

import (
	"testing"
)

func TestGetSummoner(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
