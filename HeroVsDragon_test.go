package main

import "testing"

func TestRand(t *testing.T) {
	t.Run("Testing randomize func", func(t *testing.T) {
		got := randomize(10, 100)
		want := randomize(10, 100)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
