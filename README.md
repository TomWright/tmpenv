# tmpenv

[![Build Status](https://travis-ci.org/TomWright/tmpenv.svg?branch=master)](https://travis-ci.org/TomWright/tmpenv)
[![codecov](https://codecov.io/gh/TomWright/tmpenv/branch/master/graph/badge.svg)](https://codecov.io/gh/TomWright/tmpenv)
[![Documentation](https://godoc.org/github.com/TomWright/tmpenv?status.svg)](https://godoc.org/github.com/TomWright/tmpenv)

Easily set temporary environment variables and commandline flags, usually within a testing environment.

## Quick start

`SetEnvVar`, `SetEnvVars` and `AppendOSArgs` all return a single `func`, which when executed will reset the given variables to their previous state.

*Important note - Things may not work as expected if you run your tests in parallel!*

### Setting environment variables

```
func TestSomething(t *testing.T) {
    log.Println(os.Getenv("A")) // ""
    
    t.Run("test one", func(t *testing.T) {
        resetEnvVars := tmpenv.SetEnvVars(map[string]string{
            "A":     "123",
        })
        defer resetEnvVars()

        log.Println(os.Getenv("A")) // "123"
    })
    
    log.Println(os.Getenv("A")) // ""
    
    t.Run("test two", func(t *testing.T) {
        resetEnvVars := tmpenv.SetEnvVars(map[string]string{
            "A":     "456",
        })
        defer resetEnvVars()

        log.Println(os.Getenv("A")) // "456"
    })

    log.Println(os.Getenv("A")) // ""
}
```