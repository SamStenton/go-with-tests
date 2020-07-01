package iteration

import "fmt"
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

// Run benchmarks using 'go test -bench="."'
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}

func ExampleRepeat() {
	got := Repeat("b")
	fmt.Println(got)
    // Output: bbbbb
}