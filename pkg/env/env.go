package env

import (
	"fmt"
	"os"
	"strconv"
)

type ErrMissingEnvVar string

func (name ErrMissingEnvVar) Error() string {
	return "missing environment variable: " + string(name)
}

func MustString(name string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	panic(ErrMissingEnvVar(name))
}

// LookupInt ...
func MustInt(name string) int {
	val, ok := os.LookupEnv(name)
	if !ok {
		panic(ErrMissingEnvVar(name))
	}
	if val, err := strconv.Atoi(val); err != nil {
		return val
	}
	panic(fmt.Sprintf("failed parse environment variable: %s. Value: %s", name, val))
}
