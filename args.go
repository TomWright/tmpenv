package tmpenv

import (
	"os"
)

// AppendOSArgs will append os.Args with the given values.
// You can reset the args back to their previous value by executing the returned function.
func AppendOSArgs(args ...string) func() {
	originalArgs := os.Args
	newArgs := append(os.Args, args...)
	resetFunc := func() {
		os.Args = originalArgs
	}
	os.Args = newArgs

	return resetFunc
}
