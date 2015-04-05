package main

import (
	"fmt"
	"os"
	"reflect"
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

func TestMapValue(t *testing.T) {
	var mapValue MapValue
	mapValue.Set("account=1234,animal=tapir")
	actual := mapValue.Data
	fmt.Println(mapValue)
	expected := map[string]string{"account": "1234", "animal": "tapir"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MapValue(): %v, expected %v", actual, expected)
	}
}
