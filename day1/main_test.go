package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurn(t *testing.T) {
	next, numClicks := Turn(50, "L68")
	assert.Equal(t, next, 82)
	assert.Equal(t, numClicks, 1)

	next, numClicks = Turn(60, "R48")
	assert.Equal(t, next, 8)
	assert.Equal(t, numClicks, 1)

	next, _ = Turn(36, "L963")
	assert.Equal(t, next, 73)
}

func SplitTest(t *testing.T) {
	test := strings.Split("test1 test2", " ")

	assert.Equal(t, len(test), 2)
	assert.Equal(t, test[0], "test1")
	assert.Equal(t, test[1], "test2")

	test = strings.Split("test1\ntest2", "\n")

	assert.Equal(t, len(test), 2)
	assert.Equal(t, test[0], "test1")
	assert.Equal(t, test[1], "test2")
}
