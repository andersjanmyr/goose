package main

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	os.Mkdir("tmp", 0777)
	os.Chdir("tmp")
	generate("../test", "tapir", map[string]string{"NAME": "tapir"})
	os.Chdir("..")
	os.RemoveAll("tmp")
}

func TestNewFilenameDc(t *testing.T) {
	actual := newFilename("~/.gooserc", "my-name", "prefixNAME.ccsuffix", map[string]string{"NAME": "my-name"})
	expected := "prefixMyNamesuffix"
	if actual != expected {
		t.Errorf("newFilename(): %v, expected %v", actual, expected)
	}
}
