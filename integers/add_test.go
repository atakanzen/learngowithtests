package integers_test

import (
	"fmt"
	"learngowithtests/integers"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := integers.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("got %d want %d", sum, expected)
	}
}

// This example will appear in godoc with its output.
// Note that this function will be executed only if "Output" is provided.
// This comment will also be included in the godoc.
func ExampleAdd() {
	sum := integers.Add(1, 7)
	fmt.Println(sum)
	//Output: 8
}
