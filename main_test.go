package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckWinner(t *testing.T) {
	assert.Equal(t, CheckWinner(-10, 50), 2, "It should be 1")
}

func TestCheckGameEnd(t *testing.T) {
	assert.Equal(t, CheckGameEnd(1, 5), false, "It should be false")
}

func TestFetchHeroName(t *testing.T) {
	assert.NotEqual(t, len(FetchHeroName()), 0, FetchHeroName())
}

/*func TestSelectMainMenuItem(t *testing.T) {
	assert.Equal(t, SelectMainMenuItem(), false, "It should be false")
}*/
