package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	r := strings.NewReader(input)
	assert.Equal(t, 3, part1(r))
}

func TestPart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	r := strings.NewReader(input)
	assert.Equal(t, 6, part2(r))
}
