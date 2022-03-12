package config

import (
	"testing"
)
import "embed"

//go:embed test/success/*
var success embed.FS

func TestSuccess(t *testing.T) {
	err := Parse("test/success", success)
	if err != nil {
		t.Errorf(err.Error())
	}
}

//go:embed test/bad-ext/*
var badExt embed.FS

func TestBadExt(t *testing.T) {
	err := Parse("test/bad-ext", badExt)
	if err != nil {
		t.Errorf(err.Error())
	}
}

//go:embed test/bad-json/*
var badJson embed.FS

func TestBadJson(t *testing.T) {
	err := Parse("test/bad-json", badJson)
	if err == nil {
		t.Errorf("no error ?")
	}
}

func TestDirNotFound(t *testing.T) {
	err := Parse("test/bad-dir", success)
	if err == nil {
		t.Errorf("no error when directory didn't exist ?")
	}
}
