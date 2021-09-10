package reflection_test

import (
	"learngowithtests/reflection"
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{
				"Chris",
			},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{
				"Arthur Morgan",
				"Cincinnati",
			},
			ExpectedCalls: []string{"Arthur Morgan", "Cincinnati"},
		},
		{
			Name: "Struct with a string and an int field",
			Input: struct {
				Name string
				Age  int
			}{
				"Joel",
				48,
			},
			ExpectedCalls: []string{"Joel"},
		},
		{
			Name: "Struct with nested fields",
			Input: Person{
				Name: "John Black",
				Profile: Profile{
					City: "Unknown",
					Age:  999,
				},
			},
			ExpectedCalls: []string{"John Black", "Unknown"},
		},
		{
			Name: "Pointers to things",
			Input: &Person{
				Name: "Pointed",
				Profile: Profile{
					City: "Point",
					Age:  2,
				},
			},
			ExpectedCalls: []string{"Pointed", "Point"},
		},
		{
			Name: "Slices",
			Input: []Profile{
				{
					City: "Stockholm",
					Age:  678,
				},
				{
					City: "Poznan",
					Age:  754,
				},
			},
			ExpectedCalls: []string{"Stockholm", "Poznan"},
		},
		{
			Name: "Arrays",
			Input: [2]Profile{
				{
					City: "London",
					Age:  34,
				},
				{
					City: "Brighton",
					Age:  35,
				},
			},
			ExpectedCalls: []string{"London", "Brighton"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := make([]string, 0)
			reflection.Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		reflection.Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("Channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{City: "Boston", Age: 24}
			aChannel <- Profile{City: "Los Angeles", Age: 18}
			close(aChannel)
		}()

		var got []string
		want := []string{"Boston", "Los Angeles"}

		reflection.Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Functions", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{City: "Berlin", Age: 25}, Profile{City: "Manchester", Age: 29}
		}

		var got []string
		want := []string{"Berlin", "Manchester"}

		reflection.Walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, v := range haystack {
		if v == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
