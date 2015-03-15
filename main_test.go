package main

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	os.Mkdir("tmp", 0777)
	os.Chdir("tmp")
	generate("../test", "tapir")
	os.Chdir("..")
	os.RemoveAll("tmp")
}
