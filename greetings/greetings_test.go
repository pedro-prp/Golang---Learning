package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello
func TestHelloName(t *testing.T) {
	name := "Jose"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Jose")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") =%q, %v, want match for %#q,nil`, msg, err, want)
	}
}

// TestHelloEmpty
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
