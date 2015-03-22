package main

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	os.Mkdir("tmp", 0777)
	os.Chdir("tmp")
	generate("../test", map[string]string{"NAME": "tapir"})
	os.Chdir("..")
	os.RemoveAll("tmp")
}
