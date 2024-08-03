package main

import (
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// ProgressLogger implements the io.Writer interface to log progress
type ProgressLogger struct {
	ch chan string
}

func (pl *ProgressLogger) Write(p []byte) (n int, err error) {
	pl.ch <- string(p)
	return len(p), nil
}

func Clone(url string, token string, path string, branch string, ch chan string) (*git.Repository, error) {
	progressLogger := &ProgressLogger{ch}

	r, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "risersh",
			Password: token,
		},
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          progressLogger,
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}
