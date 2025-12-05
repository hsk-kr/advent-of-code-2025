package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPart2(t *testing.T) {
	fmt.Println("test2")
	assert.Equal(t, IsValidIdPart2("565656"), false)
	assert.Equal(t, IsValidIdPart2("824824824"), false)
}

func TestGetRange(t *testing.T) {
	fmt.Println("test")
	minN, maxN := GetRange("1000000000-10000000000")
	assert.Equal(t, minN, int64(1000000000))
	assert.Equal(t, maxN, int64(10000000000))

	minN, maxN = GetRange("2121212118-2121212124")
	assert.Equal(t, minN, int64(2121212118))
	assert.Equal(t, maxN, int64(2121212124))
}
