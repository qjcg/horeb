//go:build integration
// +build integration

package main

// Integration tests; must be run after successful "go install".
// go test -v -tags integration

import (
	"os"
	"os/exec"
	"path"
	"testing"
)

const cmdName = "horeb"

func TestBuildAndRunBinary(t *testing.T) {
	dir, err := os.MkdirTemp("", cmdName)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	binPath := path.Join(dir, cmdName)
	err = exec.Command("go", "build", "-o", binPath).Run()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("NoArgs", func(t *testing.T) {
		err := exec.Command(binPath).Run()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("BadArg", func(t *testing.T) {
		err := exec.Command(binPath, "NOT_A_BLOCK").Run()
		if _, ok := err.(*exec.ExitError); !ok {
			t.Fatal(err)
		}
	})

	t.Run("GoodArg", func(t *testing.T) {
		tests := []string{
			"geometric",
			"mahjong",
			"emoji",
		}
		for _, block := range tests {
			out, err := exec.Command(binPath, block).CombinedOutput()
			if err != nil {
				t.Fatalf("%s\n%s\n", err, out)
			}
		}
	})

	t.Run("MultiGoodArgs", func(t *testing.T) {
		tests := []struct {
			blocks []string
		}{
			{[]string{"geometric", "dominos"}},
			{[]string{"emoji", "mahjong"}},
		}
		for _, test := range tests {
			out, err := exec.Command(binPath, test.blocks...).CombinedOutput()
			if err != nil {
				t.Fatalf("%s\n%s\n", err, out)
			}
		}
	})
}
