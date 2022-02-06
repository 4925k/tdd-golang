package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just a test"}

		got, _ := dictionary.Search("test")
		want := "just a test"
		checkString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just a test"}
		_, err := dictionary.Search("test123")

		checkError(t, err, ErrorInvalid)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key, value := "test", "just a test"
		dictionary.Add(key, value)
		checkMap(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key, value := "test", "just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new value")
		checkError(t, err, ErrorKeyExists)
		checkMap(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key, value := "test", "just a test"
		dictionary := Dictionary{key: value}
		newValue := "another test"

		dictionary.Update(key, newValue)

		checkMap(t, dictionary, key, newValue)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{}
		newValue := "another test"

		err := dictionary.Update(key, newValue)

		checkError(t, err, ErrorKeyDoesntExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{key: "just a test"}

		dictionary.Delete(key)

		_, err := dictionary.Search(key)
		if err != ErrorInvalid {
			t.Errorf("expected %q to be deleted", key)
		}
	})

	t.Run("delete non existent word", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(key)
		checkError(t, err, ErrorKeyDoesntExist)

	})
}

func checkString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func checkError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func checkMap(t testing.TB, d Dictionary, key, value string) {
	t.Helper()
	got, err := d.Search("test")
	if err != nil {
		t.Errorf("add failed: %v", err)
	}

	if value != got {
		t.Errorf("got %q want %q", got, value)
	}
}
