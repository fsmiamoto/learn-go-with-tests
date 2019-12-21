package main

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	t.Run("Search known word", func(t *testing.T) {
		got, err := d.Search("test")
		want := "this is just a test"

		assertValue(t, got, want)
		assertError(t, err, nil)

	})

	t.Run("Search unknown word", func(t *testing.T) {
		_, err := d.Search("jar jar")
		assertError(t, err, ErrWordNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("Add new word", func(t *testing.T) {
		d := Dictionary{}
		want := "this is a test boys"

		err := d.Add("test", want)
		assertError(t, err, nil)

		got, err := d.Search("test")

		assertValue(t, got, want)
		assertError(t, err, nil)
	})

	t.Run("Add existing word", func(t *testing.T) {
		d := Dictionary{"test": "this is a test"}

		want := "this is a different test boys"
		err := d.Add("test", want)

		assertError(t, err, ErrWordExists)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"

		d := Dictionary{word: def}
		assertDefinition(t, d, word, def)

		newDef := "this is a updated test"
		err := d.Update(word, newDef)

		assertError(t, err, nil)
		assertDefinition(t, d, word, newDef)
	})

	t.Run("Update non-existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"

		d := Dictionary{word: def}
		assertDefinition(t, d, word, def)

		newDef := "this is a updated test"
		err := d.Update("not existing", newDef)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	d := Dictionary{word: "お前はもう死んでいる"}

	assertDefinition(t, d, word, "お前はもう死んでいる")

	d.Delete(word)

	_, err := d.Search(word)
	assertError(t, err, ErrWordNotFound)

}

func assertDefinition(t *testing.T, d Dictionary, word string, def string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != def {
		t.Errorf("got %q want %q", got, def)
	}
}

func assertValue(t *testing.T, got interface{}, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
