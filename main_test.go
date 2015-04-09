package main

import (
	"os"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	os.Mkdir("tmp", 0777)
	force = true
	err := generate("fixtures/test", "tmp/tapir", map[string]string{"NAME": "tapir"})
	if err != nil {
		t.Errorf("Failed while parsing", err)
	}
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
	expected := map[string]string{"ACCOUNT": "1234", "ANIMAL": "tapir"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MapValue(): %v, expected %v", actual, expected)
	}
}

func TestReplaceCc(t *testing.T) {
	actual := replace("NAME.cc", "NAME", "groovy_dingo")
	expected := "GroovyDingo"
	if actual != expected {
		t.Errorf("Replace(): %v, expected %v", actual, expected)
	}
}

func TestReplaceDa(t *testing.T) {
	actual := replace("NAME.da", "NAME", "groovy_dingo")
	expected := "groovy-dingo"
	if actual != expected {
		t.Errorf("Replace(): %v, expected %v", actual, expected)
	}
}

func TestReplaceDc(t *testing.T) {
	actual := replace("NAME.dc", "NAME", "groovy_dingo")
	expected := "groovyDingo"
	if actual != expected {
		t.Errorf("Replace(): %v, expected %v", actual, expected)
	}
}

func TestReplaceSc(t *testing.T) {
	actual := replace("NAME.sc", "NAME", "groovyDingo")
	expected := "groovy_dingo"
	if actual != expected {
		t.Errorf("Replace(): %v, expected %v", actual, expected)
	}
}

func TestReplaceSs(t *testing.T) {
	actual := replace("NAME.ss", "NAME", "groovyDingo")
	expected := "groovy dingo"
	if actual != expected {
		t.Errorf("Replace(): %v, expected %v", actual, expected)
	}
}
