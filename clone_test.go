package main

import (
	"os"
	"testing"
)

func TestClone(t *testing.T) {
	_, err := Clone("https://github.com/risersh/risersh", os.Getenv("GITHUB_TOKEN"), "/tmp/risersh/docker-build-example", "main")
	if err != nil {
		t.Fatal(err)
	}
}
