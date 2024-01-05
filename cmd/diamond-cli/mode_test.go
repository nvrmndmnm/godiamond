package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewModeFactory(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	factory := NewModeFactory(box)

	if factory == nil {
		t.Fatal("Expected ModeFactory, got nil")
	}

	if len(factory.modes) != 3 {
		t.Errorf("Expected 3 modes, got %d", len(factory.modes))
	}

	_, deployExists := factory.modes["deploy"]
	_, cutExists := factory.modes["cut"]
	_, loupeExists := factory.modes["loupe"]

	if !deployExists || !cutExists || !loupeExists {
		t.Error("Not all expected modes were found in the ModeFactory")
	}

	if reflect.TypeOf(factory.modes["deploy"]).Kind() != reflect.Func {
		t.Error("Expected function, got different type")
	}

	if reflect.TypeOf(factory.modes["cut"]).Kind() != reflect.Func {
		t.Error("Expected function, got different type")
	}

	if reflect.TypeOf(factory.modes["loupe"]).Kind() != reflect.Func {
		t.Error("Expected function, got different type")
	}
}

func TestCreateMode(t *testing.T) {
	box, err := setupBox()
	if err != nil {
		t.Fatalf("Failed to create DiamondBox: %v", err)
	}

	factory := NewModeFactory(box)

	deployMode := factory.CreateMode("deploy")
	if _, ok := deployMode.(*DeployMode); !ok {
		t.Error("Expected DeployMode, got different type")
	}

	cutMode := factory.CreateMode("cut")
	if _, ok := cutMode.(*CutMode); !ok {
		t.Error("Expected CutMode, got different type")
	}

	loupeMode := factory.CreateMode("loupe")
	if _, ok := loupeMode.(*LoupeMode); !ok {
		t.Error("Expected LoupeMode, got different type")
	}

	unknownMode := factory.CreateMode("unknown")
	if unknownMode != nil {
		t.Error("Expected nil, got a mode")
	}
}

func TestPrintUsage(t *testing.T) {
	command := &Command{
		SubCommands: []*Command{
			{
				Name:        "test",
				Description: "Test command",
			},
		},
	}

	buf := new(bytes.Buffer)
	PrintUsage(buf, command)

	expected := "\nCommands:\n    test                 Test command\n\nArguments:\n"
	if buf.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, buf.String())
	}
}
