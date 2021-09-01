package iteration_test

import (
	"fmt"
	"learngowithtests/iteration"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat with default count", func(t *testing.T) {
		repeated := iteration.Repeat("a", 0)
		expected := "aaaaa"
		assertRepeated(t, repeated, expected)
	})

	t.Run("repeat with specified count", func(t *testing.T) {
		repeated := iteration.Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertRepeated(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("z", 0)
	}
}

func ExampleRepeat() {
	repeated := iteration.Repeat("ha", 3)
	fmt.Println(repeated)
	//Output: hahaha
}

func assertRepeated(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("got %q want %q", repeated, expected)
	}
}
