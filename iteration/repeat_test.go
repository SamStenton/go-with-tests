package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %q but got %q", got, expected)
		}
	}

	t.Run("repeat the letter a 5 times", func(t *testing.T) {
		got := Repeat("a")
		expected := "aaaaa"

		assertCorrectMessage(t, got, expected)
	})
}
