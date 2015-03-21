package main

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	name := New("brave_tapir")
	actual := name.CamelCase()
	expected := "BraveTapir"
	if actual != expected {
		t.Errorf("CamelCase(): %v, expected %v", actual, expected)
	}
}

func TestDromedarCase(t *testing.T) {
	name := New("brave_tapir")
	actual := name.DromedarCase()
	expected := "braveTapir"
	if actual != expected {
		t.Errorf("DromedarCase(): %v, expected %v", actual, expected)
	}
}

func TestDasherize(t *testing.T) {
	name := New("brave_tapir")
	actual := name.Dasherize()
	expected := "brave-tapir"
	if actual != expected {
		t.Errorf("DromedarCase(): %v, expected %v", actual, expected)
	}
}
