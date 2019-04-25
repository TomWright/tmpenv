package tmpenv_test

import (
	"github.com/tomwright/tmpenv"
	"os"
	"reflect"
	"testing"
)

func TestAppendOSArgs(t *testing.T) {
	startingArgs := os.Args
	newArgs := []string{"a", "b"}
	customArgs := append(startingArgs, newArgs...)

	if exp, got := startingArgs, os.Args; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected args [%v], got [%v]", exp, got)
	}

	resetFunc := tmpenv.AppendOSArgs(newArgs...)
	if exp, got := customArgs, os.Args; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected args [%v], got [%v]", exp, got)
	}

	resetFunc()

	if exp, got := startingArgs, os.Args; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected args [%v], got [%v]", exp, got)
	}
}
