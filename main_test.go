package main

import (
	"reflect"
	"testing"
)

func TestPathToFileSystem(t *testing.T) {
	expect := []struct {
		Input string
		Dirs  []string
		Path  string
	}{
		{
			"/about",
			[]string{},
			"about",
		},
		{
			"/about/time",
			[]string{"about"},
			"time",
		},
		{
			"about/damn/time",
			[]string{"about", "damn"},
			"time",
		},
		{
			"aboot",
			[]string{},
			"aboot",
		},
		{
			"",
			[]string{},
			"",
		},
		{
			"/",
			[]string{},
			"",
		},
	}

	for _, e := range expect {
		dirs, path := pathToFileSystem(e.Input)
		if !reflect.DeepEqual(e.Dirs, dirs) {
			t.Errorf("directories did not match on %q\n  expected: %v\n     actual: %v", e.Input, e.Dirs, dirs)
		}
		if path != e.Path {
			t.Errorf("path did not match on %q\n  expected: %v\n     actual: %v", e.Input, e.Path, path)
		}
	}
}
