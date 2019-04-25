package tmpenv

import "os"

// SetEnvVar will set the given environment variable value.
// You can reset the env var back to it's previous value by executing the returned function.
func SetEnvVar(name string, value string) func() {
	originalValue := os.Getenv(name)
	resetFunc := func() {
		_ = os.Setenv(name, originalValue)
	}
	_ = os.Setenv(name, value)
	return resetFunc
}

// SetEnvVars will set the given environment variables values.
// You can reset the env vars back to their previous value by executing the returned function.
func SetEnvVars(envVars map[string]string) func() {
	resetFuncs := make([]func(), 0)
	for k, v := range envVars {
		resetFunc := SetEnvVar(k, v)
		resetFuncs = append(resetFuncs, resetFunc)
	}
	resetFunc := func() {
		for _, r := range resetFuncs {
			r()
		}
	}
	return resetFunc
}
