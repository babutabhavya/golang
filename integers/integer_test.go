package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		got := Add(1, 2)
		wanted := 3
		if got != wanted {
			t.Errorf("got %v expected %v", got, wanted)
		}
	})
}
