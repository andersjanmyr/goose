package main

import (
	"testing"
)

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

func TestSnakeCaseSpaceSeparated(t *testing.T) {
	actual := SnakeCase("Brave tapir")
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("SnakeCase(): %v, expected %v", actual, expected)
	}
}

func TestSnakeCaseTitleCase(t *testing.T) {
	actual := SnakeCase("Brave tapir")
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("SnakeCase(): %v, expected %v", actual, expected)
	}
}

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

func TestDasherized(t *testing.T) {
	actual := Dasherized("brave_tapir")
	expected := "brave-tapir"
	if actual != expected {
		t.Errorf("Dasherize(): %v, expected %v", actual, expected)
	}
}

func TestSpaceSeparated(t *testing.T) {
	actual := SpaceSeparated("brave_tapir")
	expected := "brave tapir"
	if actual != expected {
		t.Errorf("SpaceSeparated(): %v, expected %v", actual, expected)
	}
}

func TestTitleCase(t *testing.T) {
	actual := TitleCase("brave_tapir")
	expected := "Brave Tapir"
	if actual != expected {
		t.Errorf("SpaceSeparated(): %v, expected %v", actual, expected)
	}
}
