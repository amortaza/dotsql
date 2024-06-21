package test_utils

import "testing"

func Checki(t *testing.T, expected int, got int) {
	if got != expected {
		t.Errorf("got \"%d\", expected \"%d\"", got, expected)
	}
}

func Check(t *testing.T, expected string, got string) {
	if got != expected {
		t.Errorf("got \"%s\", expected \"%s\"", got, expected)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("test expected error")
	}
}

func YesError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("test did NOT expect error")
	}
}

func Get() string {
	return "I am test utils"
}
