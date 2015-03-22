package main

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	actual := CamelCase("brave_tapir")
	expected := "BraveTapir"
	if actual != expected {
		t.Errorf("CamelCase(): %v, expected %v", actual, expected)
	}
}

func TestDromedarCase(t *testing.T) {
	actual := DromedarCase("brave_tapir")
	expected := "braveTapir"
	if actual != expected {
		t.Errorf("DromedarCase(): %v, expected %v", actual, expected)
	}
}

func TestDasherize(t *testing.T) {
	actual := Dasherized("brave_tapir")
	expected := "brave-tapir"
	if actual != expected {
		t.Errorf("Dasherize(): %v, expected %v", actual, expected)
	}
}

func TestSnakeCaseDasherized(t *testing.T) {
	actual := SnakeCase("brave-tapir")
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("SnakeCase(): %v, expected %v", actual, expected)
	}
}

func TestSnakeCaseDromedar(t *testing.T) {
	actual := SnakeCase("braveTapir")
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("SnakeCase(): %v, expected %v", actual, expected)
	}
}

func TestSnakeCaseCamel(t *testing.T) {
	actual := SnakeCase("BraveTapir")
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("SnakeCase(): %v, expected %v", actual, expected)
	}
}
