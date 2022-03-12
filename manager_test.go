package config

import (
	"testing"
)

func TestGetSuccess(t *testing.T) {
	err := Parse("test/success", success)
	if err != nil {
		t.Errorf(err.Error())
	}

	if Get("app.name") != "Vitego" {
		t.Errorf("wrong value, excepted Vitego, get %s", Get("app.name"))
	}
}

func TestGetEmpty(t *testing.T) {
	err := Parse("test/success", success)
	if err != nil {
		t.Errorf(err.Error())
	}

	if Get("empty") != "" {
		t.Errorf("expected empty string, get %s", Get("empty"))
	}
}
