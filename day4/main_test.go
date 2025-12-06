package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcJoltagePart2(t *testing.T) {
	assert.Equal(t, CalcJoltagePart2("987654321111111"), int64(987654321111))
	assert.Equal(t, CalcJoltagePart2("811111111111119"), int64(811111111119))
}
