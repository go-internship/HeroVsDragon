package main

import "testing"

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

func TestFetchHeroName(t *testing.T) {
	want := len(FetchHeroName())

	if want < 1 {
		t.Errorf("Couldn't fetch heroName from API")
	}
}

func TestSelectMainMenuItem(t *testing.T) {
	got := SelectMainMenuItem("1")
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestAttackHeroAndDragon(t *testing.T) {
	got := AttackHeroAndDragon("1")
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestInputHeroName(t *testing.T) {
	got := InputHeroName("   AA\t")
	want := "AA"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
