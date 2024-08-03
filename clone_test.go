package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func TestClone(t *testing.T) {
	repoPath := "/tmp/risersh/docker-build-example"

	// Ensure cleanup after the test
	defer func() {
		err := os.RemoveAll(repoPath)
		if err != nil {
			t.Logf("Failed to clean up test directory: %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := make(chan string)

	go func() {
		for {
			select {
			case msg := <-ch:
				log.Printf("%s", msg)
			case <-ctx.Done():
				return
			}
		}
	}()

	_, err := Clone("https://github.com/risersh/docker-build-example", os.Getenv("GITHUB_TOKEN"), repoPath, "main", ch)
	if err != nil {
		t.Fatal(err)
	}

}
