// +build integration

package main

// Integration tests; must be run after successful "go install".
// go test -v -tags integration

import (
	"os/exec"
	"testing"
)

const cmdName = "horeb"

func TestNoArgs(t *testing.T) {
	out, err := exec.Command(cmdName).CombinedOutput()
	if err != nil {
		t.Fatalf("%s\n%s\n", err, out)
	}
}

func TestBadArg(t *testing.T) {
	err := exec.Command(cmdName, "NOT_A_BLOCK").Run()
	if _, ok := err.(*exec.ExitError); !ok {
		t.Fatal(err)
	}
}

func TestGoodArg(t *testing.T) {
	tests := []string{
		"geometric",
		"mahjong",
		"emoji",
	}
	for _, block := range tests {
		out, err := exec.Command(cmdName, block).CombinedOutput()
		if err != nil {
			t.Fatalf("%s\n%s\n", err, out)
		}
	}
}

func TestMultiGoodArgs(t *testing.T) {
	tests := []struct {
		blocks []string
	}{
		{[]string{"geometric", "dominos"}},
		{[]string{"emoji", "mahjong"}},
	}
	for _, test := range tests {
		out, err := exec.Command(cmdName, test.blocks...).CombinedOutput()
		if err != nil {
			t.Fatalf("%s\n%s\n", err, out)
		}
	}
}
