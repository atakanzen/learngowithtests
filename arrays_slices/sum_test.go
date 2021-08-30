package arraysslices_test

import (
	"fmt"
	arraysslices "learngowithtests/arrays_slices"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{2, 2, 1, 0, 1, 7}

	got := arraysslices.Sum(numbers)
	want := 13

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3}
	sum := arraysslices.Sum(numbers)
	fmt.Println(sum)
	//Output: 6
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		arraysslices.Sum(numbers)
	}
}

func TestSumAllTails(t *testing.T) {
	assertSlices := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("sums of not empty slices", func(t *testing.T) {
		got := arraysslices.SumAllTails([]int{1, 2, 3}, []int{0, 1})
		want := []int{5, 1}
		assertSlices(t, got, want)
	})

	t.Run("safe sums of empty slices", func(t *testing.T) {
		got := arraysslices.SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertSlices(t, got, want)
	})
}

func ExampleSumAllTails() {
	sums := arraysslices.SumAllTails([]int{1, 4, 18}, []int{2, 4, 6}, []int{7, 8, 9})
	fmt.Println(sums)
	//Output: [22 10 17]
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arraysslices.SumAllTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	}
}
