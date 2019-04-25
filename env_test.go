package tmpenv_test

import (
	"github.com/tomwright/tmpenv"
	"os"
	"testing"
)

func TestSetEnvVar(t *testing.T) {
	_ = os.Setenv("X", "a")

	if exp, got := "a", os.Getenv("X"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}

	resetFunc := tmpenv.SetEnvVar("X", "b")

	if exp, got := "b", os.Getenv("X"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}

	resetFunc()

	if exp, got := "a", os.Getenv("X"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}

	_ = os.Setenv("X", "")
}

func TestSetEnvVars(t *testing.T) {
	_ = os.Setenv("X", "a")
	_ = os.Setenv("Y", "z")

	if exp, got := "a", os.Getenv("X"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}
	if exp, got := "z", os.Getenv("Y"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}

	resetFunc := tmpenv.SetEnvVars(map[string]string{
		"X": "b",
		"Y": "c",
	})

	if exp, got := "b", os.Getenv("X"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}
	if exp, got := "c", os.Getenv("Y"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}

	resetFunc()

	if exp, got := "a", os.Getenv("X"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}
	if exp, got := "z", os.Getenv("Y"); exp != got {
		t.Errorf("expected env var `%s`, got `%s`", exp, got)
	}

	_ = os.Setenv("X", "")
	_ = os.Setenv("Z", "")
}
