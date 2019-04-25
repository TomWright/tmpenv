package tmpenv_test

import (
	"fmt"
	"github.com/tomwright/tmpenv"
	"os"
	"testing"
)

func ExampleSetEnvVar() {
	fmt.Println(os.Getenv("XYZ")) // ""

	reset := tmpenv.SetEnvVar("XYZ", "123")

	// Start tests that depend on XYZ
	fmt.Println(os.Getenv("XYZ")) // "123"
	// Stop tests that depend on XYZ

	reset()

	fmt.Println(os.Getenv("XYZ")) // ""

	// Output:
	//
	// 123
	//
}

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
