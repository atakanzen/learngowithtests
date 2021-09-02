package maps_test

import (
	"fmt"
	"learngowithtests/maps"
	"strconv"
	"testing"
)

const (
	testKey          = "test"
	testValue        = "this is just a test"
	testUpdatedValue = "updated test value"
)

func TestSearch(t *testing.T) {
	dictionary := maps.Dictionary{testKey: testValue}

	t.Run("existing keyword", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := testValue
		assertStrings(t, got, want)
	})

	t.Run("non existing keyword", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, maps.ErrNotFound)
	})
}

func ExampleDictionary_Search() {
	dictionary := maps.Dictionary{"hi": "cześć"}
	val, err := dictionary.Search("hi")
	fmt.Println(val)
	fmt.Println(err)
	//Output: cześć
	// <nil>
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := make(maps.Dictionary)
		err := dictionary.Add(testKey, testValue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, testKey, testValue)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := maps.Dictionary{testKey: testValue}
		err := dictionary.Add(testKey, "ant")

		assertError(t, err, maps.ErrWordExists)
		assertDefinition(t, dictionary, testKey, testValue)
	})
}

func ExampleDictionary_Add() {
	dictionary := make(maps.Dictionary)
	err := dictionary.Add("apple", "jabłko")
	fmt.Println(err)
	//Output: <nil>

}

func TestUpdate(t *testing.T) {
	dictionary := maps.Dictionary{testKey: testValue}

	t.Run("existing keyword", func(t *testing.T) {

		err := dictionary.Update(testKey, testUpdatedValue)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, testKey, testUpdatedValue)
	})

	t.Run("non existing keyword", func(t *testing.T) {
		err := dictionary.Update("nonExistingKey", testUpdatedValue)
		assertError(t, err, maps.ErrWordDoesNotExist)
	})
}

func ExampleDictionary_Update() {
	dictionary := maps.Dictionary{"motto": "foo"}
	err := dictionary.Update("motto", "luctor et emergo")
	fmt.Println(err)
	//Output: <nil>

}

func TestDelete(t *testing.T) {
	dictionary := maps.Dictionary{testKey: testValue}

	dictionary.Delete(testKey)
	_, err := dictionary.Search(testKey)
	if err != maps.ErrNotFound {
		t.Errorf("expected %q to be deleted", err)
	}
}

func ExampleDictionary_Delete() {
	dictionary := maps.Dictionary{"delete": "me"}
	dictionary.Delete("delete")
	//Output:

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want error %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary maps.Dictionary, key, definition string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find the added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func generateBenchmarkDictionary(b *testing.B) maps.Dictionary {
	dictionary := make(maps.Dictionary, 0)
	for i := 0; i < b.N; i++ {
		dictionary.Add(strconv.Itoa(i), strconv.Itoa(i))
	}

	return dictionary
}
