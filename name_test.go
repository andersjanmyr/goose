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
		t.Errorf("Dasherize(): %v, expected %v", actual, expected)
	}
}

func TestSnakeCaseDasherized(t *testing.T) {
	name := New("brave-tapir")
	actual := name.SnakeCase()
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("SnakeCase(): %v, expected %v", actual, expected)
	}
}

func TestSnakeCaseDromedar(t *testing.T) {
	name := New("braveTapir")
	actual := name.SnakeCase()
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("SnakeCase(): %v, expected %v", actual, expected)
	}
}

func TestSnakeCaseCamel(t *testing.T) {
	name := New("BraveTapir")
	actual := name.SnakeCase()
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("SnakeCase(): %v, expected %v", actual, expected)
	}
}
