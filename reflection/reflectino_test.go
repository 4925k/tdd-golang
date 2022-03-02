package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"melina"},
			[]string{"melina"},
		},
		{
			"struct with multiple string field",
			struct {
				Name string
				City string
			}{"dibek", "kathmandu"},
			[]string{"dibek", "kathmandu"},
		},
		{
			"struct with no string field",
			struct {
				Name string
				Age  int
			}{"nishan", 26},
			[]string{"nishan"},
		},
		{
			"nested struct",
			Person{"sasim", Profile{24, "kathmandu"}},
			[]string{"sasim", "kathmandu"},
		},
		{
			"pointer to things",
			&Person{"samrat", Profile{24, "ktm"}},
			[]string{"samrat", "ktm"},
		},
		{
			"slices",
			[]Profile{
				{23, "lalitpur"},
				{25, "bhaktapur"},
			},
			[]string{"lalitpur", "bhaktapur"},
		},
		{
			"arrays",
			[2]Profile{
				{23, "lalitpur"},
				{25, "bhaktapur"},
			},
			[]string{"lalitpur", "bhaktapur"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v\n", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		var got []string
		sample := map[string]string{
			"hello": "world",
			"foo":   "bar",
		}
		walk(sample, func(input string) {
			got = append(got, input)
		})
		for _, v := range sample {
			assertContains(t, v, got)
		}
	})

	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{28, "delhi"}
			aChannel <- Profile{32, "delhi"}
			close(aChannel)
		}()

		var got []string
		want := []string{"delhi", "delhi"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v\n", got, want)
		}

	})

	t.Run("functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{32, "india"}, Profile{28, "india"}
		}

		var got []string
		want := []string{"india", "india"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, value string, values []string) {
	t.Helper()
	contains := false
	for _, v := range values {
		if v == value {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q", values, value)
	}
}
