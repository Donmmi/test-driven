package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test":"test value"}

	t.Run("search key exists", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "test value"

		assertNoError(t, err)
		assertValue(t, got, want)
	})

	t.Run("search key do not exist", func(t *testing.T) {
		_, err := dict.Search("test2")

		assertError(t, err, ErrKeyNotExists)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{"test":"test value"}

	t.Run("add non exists key", func(t *testing.T) {
		err := dict.Add("test2", "test2 value")
		assertNoError(t, err)

		got, err := dict.Search("test2")
		want := "test2 value"

		assertNoError(t, err)
		assertValue(t, got, want)
	})

	t.Run("add exists key", func(t *testing.T) {
		err := dict.Add("test", "test value")
		assertError(t, err, ErrKeyExists)
	})
}

func TestUpdate(t *testing.T) {
	dict := Dictionary{"test":"test value"}

	t.Run("update exists key", func(t *testing.T) {
		err := dict.Update("test", "test updated")
		assertNoError(t, err)

		v, err := dict.Search("test")
		assertNoError(t, err)
		assertValue(t, v, "test updated")
	})

	t.Run("update non exists key", func(t *testing.T) {
		err := dict.Update("test new", "test new value")
		assertError(t, err, ErrKeyNotExists)

		_, err = dict.Search("test new")
		assertError(t, err, ErrKeyNotExists)
	})
}

func TestDelete(t *testing.T) {
	dict := Dictionary{"test": "test value"}

	t.Run("delete exists key", func(t *testing.T) {
		dict.Delete("test")

		_, err := dict.Search("test")
		assertError(t, err, ErrKeyNotExists)
	})

	t.Run("delete non exists key", func(t *testing.T) {
		dict.Delete("test2")
	})
}

func assertValue(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}

func assertError(t *testing.T, err error, want error) {
	if err == nil {
		t.Fatal("should return key not exists err")
	}

	if err.Error() != want.Error() {
		t.Errorf("got:[%s], expected:[%s]", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error("should not return an error")
	}
}