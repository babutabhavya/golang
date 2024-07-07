package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to a person", func(t *testing.T) {
		got := Hello("Bhavya")
		wanted := "Hello Bhavya"
		if got != wanted {
			t.Errorf("got %q, wanted %q", got, wanted)
		}
	})

	t.Run("Say hello to a empty string", func(t *testing.T) {
		got := Hello("")
		wanted := "Hello "
		if got != wanted {
			t.Errorf("got %q, wanted %q", got, wanted)
		}
	})
}
