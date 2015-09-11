package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"path"
)

func gitClone(repoUrl string) (string, error) {
	name := fmt.Sprintf("%x", md5.Sum([]byte(repoUrl)))
	target := path.Join(os.TempDir(), "goose", name)

	if fi, err := os.Stat(target); fi != nil {
		if err := os.RemoveAll(target); err != nil {
			return "", err
		}
	} else if !os.IsNotExist(err) {
		return "", err
	}

	cmd := exec.Command("git", "clone", repoUrl, target)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Fprint(os.Stderr, string(out))
		return "", err
	}
	return target, nil
}
