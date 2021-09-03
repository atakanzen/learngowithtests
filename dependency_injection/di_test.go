package dependencyinjection_test

import (
	"bytes"
	di "learngowithtests/dependency_injection"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	di.Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
