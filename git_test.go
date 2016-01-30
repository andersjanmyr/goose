package main

import (
	"os"
	"testing"
)

func assertDirExists(t *testing.T, dir string) {
	if fi, err := os.Stat(dir); err != nil {
		t.Error(err)
	} else if !fi.IsDir() {
		t.Errorf("%s is not a directory")
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func IgnoreTestGitHubGitProtocol(t *testing.T) {
	dir, err := gitClone("git@github.com:andersjanmyr/goose.git")
	assertNoError(t, err)
	assertDirExists(t, dir)
}

func IgnoreTestBitbucketHTTPS(t *testing.T) {
	dir, err := gitClone("https://bitbucket.org/andersjanmyr/dummy")
	assertNoError(t, err)
	assertDirExists(t, dir)
}
