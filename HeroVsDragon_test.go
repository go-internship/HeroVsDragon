package main

import "testing"

/*func TestRand(t *testing.T) {
	t.Run("Testing randomize func", func(t *testing.T) {
		got := randomize(10, 100)
		want := randomize(10, 100)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}*/

func TestCasesAttackToDragon(t *testing.T) {
	got := CasesAttackToDragon(70)
	want := 30

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCasesAttackToHero(t *testing.T) {
	got := CasesAttackToHero(20)
	want := 80

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestShowWinner(t *testing.T) {
	got := ShowWinner(-10, 10)
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCheckCurrentHp(t *testing.T) {
	got := CheckCurrentHp(1, 5)
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
