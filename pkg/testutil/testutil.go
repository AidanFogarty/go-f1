// Package testutil provides helper methods for newbie tests.
// Thanks, https://github.com/benbjohnson/testing
package testutil

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, expected, actual)
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
	}
}

// NotNil fails if an err is nil.
func NotNil(tb testing.TB, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d: %v\033[39m\n\n", filepath.Base(file), line, err.Error())
	}
}
