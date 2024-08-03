package main

import (
	"io"
	"log"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// ProgressLogger implements the io.Writer interface to log progress
type ProgressLogger struct {
	logger *log.Logger
}

func (pl *ProgressLogger) Write(p []byte) (n int, err error) {
	pl.logger.Printf("%s", p)
	return len(p), nil
}

func Clone(url string, token string, path string, branch string) (*git.Repository, error) {
	// Create a log file
	logFile, err := os.Create(path + "/clone.log")
	if err != nil {
		return nil, err
	}
	defer logFile.Close()

	// Create a multi-writer to log to both file and stdout
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	logger := log.New(multiWriter, "", log.LstdFlags)

	progressLogger := &ProgressLogger{logger: logger}

	// ... existing code ...
	r, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "risersh",
			Password: token,
		},
		URL:      url,
		Progress: progressLogger,
	})
	if err != nil {
		return nil, err
	}

	return r, err
}
