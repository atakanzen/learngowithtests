package dependencyinjection_test

import (
	"bytes"
	di "learngowithtests/dependency_injection"
	"os"
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

func ExampleGreet() {

	// These can be used as a writer as well.
	// buffer := bytes.Buffer{}
	// writer := http.ResponseWriter

	stdout := os.Stdout
	di.Greet(stdout, "strawberries and cherries")
	//Output: Hello, strawberries and cherries

}
