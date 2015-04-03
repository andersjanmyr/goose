package main

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	os.Mkdir("tmp", 0777)
	generate("fixtures/test", "tmp/tapir", map[string]string{"NAME": "tapir"})
	os.RemoveAll("tmp")
}

func TestGenerateMissingFunction(t *testing.T) {
	os.Mkdir("tmp", 0777)
	err := generate("fixtures/error", "tmp/fail", map[string]string{"NAME": "tapir"})
	if err == nil {
		t.Errorf("Expected error when parsing template")
	}
	os.RemoveAll("tmp")
}

func TestNewFilenameDc(t *testing.T) {
	actual := newFilename("~/.gooserc", "my-name", "prefixNAME.ccsuffix", map[string]string{"NAME": "my-name"})
	expected := "prefixMyNamesuffix"
	if actual != expected {
		t.Errorf("newFilename(): %v, expected %v", actual, expected)
	}
}
