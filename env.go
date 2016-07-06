package env

import (
	"os"
	"strconv"
	"time"
)

var (
	errVariableNotSet = "Variable %s not set"
	errParseVariable  = "Error trying to parse %s: %s"
)

// String will return the environment variable when it is not empty. If the
// value is empty (or not set), the notifier will be notified and the default
// value will be returned.
func String(env, def string) string {
	ev := os.Getenv(env)
	if ev == "" {
		notifier.Notify(errVariableNotSet, env)
		return def
	}

	return ev
}

// Bool will return true when the environment applies to the strconv true rules,
// it returns false if it applies to the strconv false rules. It will return the
// default value if none of these apply and the notifier will be notified.
func Bool(env string, def bool) bool {
	ev := os.Getenv(env)
	if ev == "" {
		notifier.Notify(errVariableNotSet, env)
		return def
	}

	b, err := strconv.ParseBool(ev)
	if err != nil {
		notifier.Notify(errParseVariable, env, err)
		return def
	}

	return b
}

// Duration will try and parse the duration given in the environment variable.
// If the variable is not set or not a valid duration, the notifier will be
// notified and the default value will be returned.
func Duration(env string, def time.Duration) time.Duration {
	ev := os.Getenv(env)
	if ev == "" {
		notifier.Notify(errVariableNotSet, env)
		return def
	}

	dur, err := time.ParseDuration(ev)
	if err != nil {
		notifier.Notify(errParseVariable, env, err.Error())
		return def
	}

	return dur
}

func parseInt(env string, base, bitSize int) (int64, bool) {
	ev := os.Getenv(env)
	if ev == "" {
		notifier.Notify(errVariableNotSet, env)
		return 0, false
	}

	i, err := strconv.ParseInt(ev, base, bitSize)
	if err != nil {
		notifier.Notify(errParseVariable, env, err.Error())
		return 0, false
	}

	return i, true
}

// Int64 will try and parse the value given in the environment variable as an
// int64. If the variable is not set or the value is not a valid int64, the
// notifier will be notified and the default value will be returned.
func Int64(env string, def int64) int64 {
	if v, ok := parseInt(env, 10, 64); ok {
		return v
	}

	return def
}

// Int32 will try and parse the value given in the environment variable as an
// int32. If the variable is not set or the value is not a valid int32, the
// notifier will be notified and the default value will be returned.
func Int32(env string, def int32) int32 {
	if v, ok := parseInt(env, 10, 32); ok {
		return int32(v)
	}

	return def
}

// Int will try and parse the value given in the environment variable as an int.
// If the variable is not set or the value is not a valid int, the notifier will
// be notified and the default value will be returned.
func Int(env string, def int) int {
	if v, ok := parseInt(env, 10, 0); ok {
		return int(v)
	}

	return def
}

func parseFloat(env string, bitSize int) (float64, bool) {
	ev := os.Getenv(env)
	if ev == "" {
		notifier.Notify(errVariableNotSet, env)
		return 0, false
	}

	i, err := strconv.ParseFloat(ev, bitSize)
	if err != nil {
		notifier.Notify(errParseVariable, env, err.Error())
		return 0, false
	}

	return i, true
}

// Float64 will try and parse the value given in the environment variable as an
// float64. If the variable is not set or the value is not a valid float64, the
// notifier will be notified and the default value will be returned.
func Float64(env string, def float64) float64 {
	if v, ok := parseFloat(env, 64); ok {
		return v
	}

	return def
}

// Float32 will try and parse the value given in the environment variable as an
// float32. If the variable is not set or the value is not a valid float32, the
// notifier will be notified and the default value will be returned.
func Float32(env string, def float32) float32 {
	if v, ok := parseFloat(env, 32); ok {
		return float32(v)
	}

	return def
}
