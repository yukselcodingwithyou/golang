package greetings

import (
	"regexp"
	"testing"
)

// Ending a file's name with _test.go tells the go test command that this file contains test functions.

func TestHelloName(t *testing.T) {
	name := "Yuksel"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Yuksel")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}

}
