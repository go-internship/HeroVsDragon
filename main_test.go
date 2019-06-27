package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasesAttackToDragon(t *testing.T) {
	assert.Equal(t, CasesAttackToDragon(70), 30, "They should be equal")
}

func TestCasesAttackToHero(t *testing.T) {
	assert.Equal(t, CasesAttackToHero(20), 80, "They should be equal")
}

func TestCheckWinner(t *testing.T) {
	assert.Equal(t, CheckWinner(80, 50), 1, "It should be 1")
}

func TestCheckCurrentHp(t *testing.T) {
	assert.Equal(t, CheckCurrentHp(1, 5), false, "It should be false")
}

func TestFetchHeroName(t *testing.T) {
	assert.NotEqual(t, len(FetchHeroName()), 0, "It should not be 0")
}

func TestSelectMainMenuItem(t *testing.T) {
	assert.Equal(t, SelectMainMenuItem("1"), true, "It should be true")
}

func TestAttackHeroAndDragon(t *testing.T) {
	assert.Equal(t, AttackHeroAndDragon("1"), true, "It should be true")
}

func TestInputHeroName(t *testing.T) {
	assert.Equal(t, InputHeroName("    AA\t"), "AA", "It should be without spaces at the start and the end")
	assert.NotEqual(t, len(InputHeroName(``)), 0, "It should contain a name from API")
}
