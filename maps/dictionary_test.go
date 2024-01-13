package mapsdict

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dict.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dict, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{key: value}
		err := dict.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDontExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "test definition"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, key, value string) {
	t.Helper()

	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, value)
}
