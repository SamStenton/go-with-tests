package integers

import "fmt"
import "testing"

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected %d but got %d", sum, expected)
		}
	}

	t.Run("adding two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}